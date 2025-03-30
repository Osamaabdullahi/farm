package controllers

import (
	"farmconnect/config"
	"farmconnect/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProduce - Create a new produce listing
func CreateProduce(c *gin.Context) {
	var produce models.Produce
	if err := c.ShouldBindJSON(&produce); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save to database
	if err := config.DB.Create(&produce).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, produce)
}

// GetAllProduce - Retrieve all produce listings
func GetAllProduce(c *gin.Context) {
	var produceList []models.Produce
	if err := config.DB.Preload("Farmer").Find(&produceList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, produceList)
}

// GetProduceByID - Retrieve a single produce item by ID
func GetProduceByID(c *gin.Context) {
	id := c.Param("id")
	var produce models.Produce

	if err := config.DB.Preload("Farmer").First(&produce, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produce not found"})
		return
	}

	c.JSON(http.StatusOK, produce)
}

func GetProduceByUser(c *gin.Context) {
	userID := c.Param("user_id") // Get user ID from request parameters
	var produce []models.Produce

	// Query the database to find all produce items for the given user
	if err := config.DB.Preload("Farmer").
		Where("farmer_id = ?", userID).
		Find(&produce).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No produce found for this user"})
		return
	}

	c.JSON(http.StatusOK, produce)
}


// UpdateProduce - Allow partial updates to a produce listing
func UpdateProduce(c *gin.Context) {
	id := c.Param("id")
	var existingProduce models.Produce

	// Check if the produce listing exists
	if err := config.DB.First(&existingProduce, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produce not found"})
		return
	}

	// Create a map to hold the update fields
	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update only the provided fields
	if err := config.DB.Model(&existingProduce).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the updated record
	c.JSON(http.StatusOK, existingProduce)
}


// DeleteProduce - Delete a produce listing
func DeleteProduce(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Produce{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produce deleted successfully"})
}
