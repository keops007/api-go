package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"go-api/config"
	"go-api/models"
)

// Struct used for login/register requests
type AuthInput struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

// =======================
// REGISTER HANDLER
// =======================
func Register(c *gin.Context) {
    var input AuthInput

    // Validate incoming JSON
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "invalid input",
        })
        return
    }

    // Check if user already exists
    var existingUser models.User
    config.DB.Where("email = ?", input.Email).First(&existingUser)

    if existingUser.ID != 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "user already exists",
        })
        return
    }

    // Hash the password using bcrypt
    hashedPassword, err := bcrypt.GenerateFromPassword(
        []byte(input.Password),
        bcrypt.DefaultCost,
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "failed to hash password",
        })
        return
    }

    // Create user object
    user := models.User{
        Email:    input.Email,
        Password: string(hashedPassword),
    }

    // Save user to database
    result := config.DB.Create(&user)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "failed to create user",
        })
        return
    }

    // Return created user (without password ideally in real apps)
    c.JSON(http.StatusCreated, gin.H{
        "message": "user registered successfully",
    })
}

// =======================
// LOGIN HANDLER
// =======================
func Login(c *gin.Context) {
    var input AuthInput

    // Validate request body
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "invalid input",
        })
        return
    }

    // Find user by email
    var user models.User
    result := config.DB.Where("email = ?", input.Email).First(&user)

    if result.Error != nil {
        // Do not reveal whether email exists or not (security best practice)
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "invalid credentials",
        })
        return
    }

    // Compare hashed password with provided password
    err := bcrypt.CompareHashAndPassword(
        []byte(user.Password),
        []byte(input.Password),
    )

    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "invalid credentials",
        })
        return
    }

    // Get JWT secret from environment
    jwtSecret := os.Getenv("JWT_SECRET")

    // Create JWT token with claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "email":   user.Email,
        "exp":     time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24h
    })

    // Sign the token
    tokenString, err := token.SignedString([]byte(jwtSecret))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "could not generate token",
        })
        return
    }

    // Return token to client
    c.JSON(http.StatusOK, gin.H{
        "token": tokenString,
    })
}
