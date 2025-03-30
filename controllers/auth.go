package controllers

import (
	"farmconnect/config"
	"farmconnect/models"
	"farmconnect/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Register a new user
func Register(c *gin.Context) {
	var req struct {
		
		ID uuid.UUID `json:"id"`
		FirstName  string    `json:"first_name"`
		LastName   string    `json:"last_name"`
		Email      string    `json:"email"`
		Password   string    `json:"password"` 
		Onboarding bool      `json:"onboarding"`
		Role       string    `json:"role"`  
		Status     string    `json:"status"` 
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hash the password before saving
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		ID :req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
		Role:      req.Role,
		Onboarding:req.Onboarding,
		CreatedAt : req.CreatedAt,
		UpdatedAt  :req.UpdatedAt,
		Status:req.Status,
	}

	// Save user to database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	
	c.JSON(http.StatusCreated, gin.H{"message": user})
}

// Login user and return token
func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	if err := config.DB.Preload("Farmer").Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Check password
	if !utils.CheckPassword(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID.String(), user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Store user session in local storage (simulated)
	c.SetCookie("auth_token", token, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}


// Get all users
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Preload("Farmer").Preload("Business").Find(&users)
	c.JSON(http.StatusOK, users)
}

// Get a single user by ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.Preload("Farmer").Preload("Logistics").Preload("Business").Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Update a user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var updateData models.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// If updating password, hash it
	if updateData.Password != "" {
		hashedPassword, err := utils.HashPassword(updateData.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		user.Password = hashedPassword
	}

	// Update fields
	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName
	user.Email = updateData.Email
	user.Onboarding = updateData.Onboarding
	user.Role = updateData.Role
	user.Status = updateData.Status

	config.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

// Delete a user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
