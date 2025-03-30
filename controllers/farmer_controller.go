package controllers

import (
	"farmconnect/config"
	"farmconnect/models"
	"net/http"

	"github.com/gin-gonic/gin"
)



func CreateFarmer(c *gin.Context) {
	var farmer models.Farmer

	if err := c.ShouldBindJSON(&farmer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user exists
	var user models.User
	if err := config.DB.Where("id = ?", farmer.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Save logistics profile
	if err := config.DB.Create(&farmer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create farmer profile"})
		return
	}

	// Update user's onboarding status
	config.DB.Model(&user).Update("onboarding", true)

	c.JSON(http.StatusCreated, gin.H{"message": "farmer profile created successfully", "logistics": farmer})
}


// Get all farmers
func GetFarmers(c *gin.Context) {
	var farmers []models.Farmer
	config.DB.Find(&farmers)
	c.JSON(http.StatusOK, farmers)
}

// Get a single farmer by ID
func GetFarmerByID(c *gin.Context) {
	id := c.Param("id")
	var farmer models.Farmer

	if err := config.DB.Where("id = ?", id).First(&farmer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Farmer not found"})
		return
	}

	c.JSON(http.StatusOK, farmer)
}

// Update a farmer profile
func UpdateFarmer(c *gin.Context) {
	id := c.Param("id")
	var farmer models.Farmer

	if err := config.DB.Where("id = ?", id).First(&farmer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Farmer not found"})
		return
	}

	if err := c.ShouldBindJSON(&farmer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&farmer)
	c.JSON(http.StatusOK, gin.H{"message": "Farmer profile updated successfully", "farmer": farmer})
}

// Delete a farmer profile
func DeleteFarmer(c *gin.Context) {
	id := c.Param("id")
	var farmer models.Farmer

	if err := config.DB.Where("id = ?", id).First(&farmer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Farmer not found"})
		return
	}

	config.DB.Delete(&farmer)
	c.JSON(http.StatusOK, gin.H{"message": "Farmer profile deleted successfully"})
}
