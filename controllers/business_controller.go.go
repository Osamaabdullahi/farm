package controllers

import (
	"farmconnect/config"
	"farmconnect/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create a new business profile
func CreateBusiness(c *gin.Context) {
	var business models.Business

	// Parse JSON request
	if err := c.ShouldBindJSON(&business); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user exists
	var user models.User
	if err := config.DB.Where("id = ?", business.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Validate required fields
	if business.BusinessName == "" || business.BusinessType == "" ||
		business.BusinessLocation == "" || business.VerificationID == "" ||
		business.PreferredPayment == "" || business.ContactPersonName == "" ||
		business.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}
	
	// Set creation time
	business.CreatedAt = time.Now()
	business.UpdatedAt = time.Now()

	// Save business profile
	if err := config.DB.Create(&business).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create business profile"})
		return
	}

	// Mark onboarding as complete
	config.DB.Model(&user).Update("onboarding", true)

	c.JSON(http.StatusCreated, gin.H{"message": "Business profile created successfully", "business": business})
}


// Get all businesses
func GetBusinesses(c *gin.Context) {
	var businesses []models.Business
	config.DB.Find(&businesses)
	c.JSON(http.StatusOK, businesses)
}

// Get a single business by ID
func GetBusinessByID(c *gin.Context) {
	id := c.Param("id")
	var business models.Business

	if err := config.DB.Where("id = ?", id).First(&business).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Business not found"})
		return
	}

	c.JSON(http.StatusOK, business)
}

// Update a business profile
func UpdateBusiness(c *gin.Context) {
	id := c.Param("id")
	var business models.Business

	if err := config.DB.Where("id = ?", id).First(&business).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Business not found"})
		return
	}

	if err := c.ShouldBindJSON(&business); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&business)
	c.JSON(http.StatusOK, gin.H{"message": "Business profile updated successfully", "business": business})
}

// Delete a business profile
func DeleteBusiness(c *gin.Context) {
	id := c.Param("id")
	var business models.Business

	if err := config.DB.Where("id = ?", id).First(&business).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Business not found"})
		return
	}

	config.DB.Delete(&business)
	c.JSON(http.StatusOK, gin.H{"message": "Business profile deleted successfully"})
}
