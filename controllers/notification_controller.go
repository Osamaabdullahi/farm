package controllers

import (
	"farmconnect/config"
	"farmconnect/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateNotification - Create a new notification
func CreateNotification(c *gin.Context) {
	var notification models.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notification.ID = uuid.New()

	// Save to database
	if err := config.DB.Create(&notification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, notification)
}

// GetAllNotifications - Retrieve all notifications
func GetAllNotifications(c *gin.Context) {
	var notifications []models.Notification
	if err := config.DB.Find(&notifications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// GetNotificationsByUser - Retrieve notifications for a specific user
func GetNotificationsByUser(c *gin.Context) {
	userID := c.Param("user_id")
	var notifications []models.Notification

	if err := config.DB.Where("user_id = ?", userID).Find(&notifications).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No notifications found for this user"})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// MarkNotificationAsRead - Mark a notification as read
func MarkNotificationAsRead(c *gin.Context) {
	id := c.Param("id")
	var notification models.Notification

	// Check if the notification exists
	if err := config.DB.First(&notification, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	// Update the notification status
	notification.IsRead = true
	if err := config.DB.Save(&notification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notification)
}

// DeleteNotification - Delete a notification
func DeleteNotification(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Notification{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification deleted successfully"})
}


