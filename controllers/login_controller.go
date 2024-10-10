package controllers

import (
	"gopro/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// LoginRequest represents the expected login request body
type LoginRequest struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required"`
}

// Secret key for signing JWT (loaded from environment variable)
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Login handles user login and generates a JWT token
func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	var user models.User
	if err := models.DB.Where("email = ?", req.Email).Find(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check if the provided password is correct
	if err := user.CheckPassword(req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the JWT token
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

// Generate a JWT token for a user
func generateJWT(user models.User) (string, error) {
	// Define JWT claims
	claims := jwt.MapClaims{
		"sub": user.ID,                               // Subject (user ID)
		"exp": time.Now().Add(time.Hour * 72).Unix(), // Token expiry (72 hours)
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString(jwtSecret)
}
