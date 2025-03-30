package controllers

import (
	"farmconnect/config"
	"farmconnect/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create a new logistics provider
func CreateLogistics(c *gin.Context) {
	var logistics models.Logistics

	if err := c.ShouldBindJSON(&logistics); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user exists
	var user models.User
	if err := config.DB.Where("id = ?", logistics.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Save logistics profile
	if err := config.DB.Create(&logistics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create logistics profile"})
		return
	}

	// Update user's onboarding status
	config.DB.Model(&user).Update("onboarding", true)

	c.JSON(http.StatusCreated, gin.H{"message": "Logistics profile created successfully", "logistics": logistics})
}

// Get all logistics providers
func GetLogisticsProviders(c *gin.Context) {
	var logistics []models.Logistics
	config.DB.Preload("User").Find(&logistics) // Preload User data
	c.JSON(http.StatusOK, logistics)
}

// Get a single logistics provider by ID
func GetLogisticsByID(c *gin.Context) {
	id := c.Param("id")
	var logistics models.Logistics

	if err := config.DB.Preload("User").Where("id = ?", id).First(&logistics).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Logistics provider not found"})
		return
	}

	c.JSON(http.StatusOK, logistics)
}

// Update a logistics provider
func UpdateLogistics(c *gin.Context) {
	id := c.Param("id")
	var logistics models.Logistics

	if err := config.DB.Where("id = ?", id).First(&logistics).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Logistics provider not found"})
		return
	}

	if err := c.ShouldBindJSON(&logistics); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&logistics)
	c.JSON(http.StatusOK, gin.H{"message": "Logistics profile updated successfully", "logistics": logistics})
}

// Delete a logistics provider
func DeleteLogistics(c *gin.Context) {
	id := c.Param("id")
	var logistics models.Logistics

	if err := config.DB.Where("id = ?", id).First(&logistics).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Logistics provider not found"})
		return
	}

	config.DB.Delete(&logistics)
	c.JSON(http.StatusOK, gin.H{"message": "Logistics profile deleted successfully"})
}
