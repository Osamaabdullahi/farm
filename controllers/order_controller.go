package controllers

import (
	"net/http"
	"time"

	"farmconnect/config"
	"farmconnect/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Create Order (Checkout)
func Checkout(c *gin.Context) {
	var orderReq struct {
		UserID       uuid.UUID `json:"user_id"`
		Items        []struct {
			ProduceID uuid.UUID `json:"produce_id"`
			Quantity  int       `json:"quantity"`
		} `json:"items"`
		PaymentMethod string `json:"payment_method"`
		FarmerID       uuid.UUID  `json:"farmer_id"`
	}

	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate user exists
	var user models.User
	if err := config.DB.First(&user, "id = ?", orderReq.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var totalAmount float64
	var orderItems []models.OrderItem

	// Process each item
	for _, item := range orderReq.Items {
		var produce models.Produce
		if err := config.DB.First(&produce, "id = ?", item.ProduceID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produce not found"})
			return
		}

		// Check stock availability
		if produce.AvailableStock < item.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock available"})
			return
		}

		// Deduct stock
		produce.AvailableStock -= item.Quantity
		config.DB.Save(&produce)

		// Calculate cost
		totalCost := float64(item.Quantity) * produce.PricePerUnit
		totalAmount += totalCost

		// Create OrderItem
		orderItems = append(orderItems, models.OrderItem{
			ID:        uuid.New(),
			ProduceID: produce.ID,
			Quantity:  item.Quantity,
			Price:     produce.PricePerUnit,
			TotalCost: totalCost,
		})
	}

	// Create Order
	order := models.Order{
		ID:            uuid.New(),
		UserID:        orderReq.UserID,
		TotalAmount:   totalAmount,
		Status:        "pending",
		PaymentMethod: orderReq.PaymentMethod,
		Items:         orderItems,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		FarmerID:       orderReq.FarmerID,
	}

	// Save Order & OrderItems
	config.DB.Create(&order)
	for i := range orderItems {
		orderItems[i].OrderID = order.ID
		config.DB.Create(&orderItems[i])
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully", "order_id": order.ID})
}

// Get All Orders
func GetAllOrders(c *gin.Context) {
	var orders []models.Order
	if err := config.DB.
		Preload("User").              // Load the User associated with the order
		Preload("Items").             // Load the OrderItems in each order
		Preload("Items.Produce").     // Load the Produce details inside each OrderItem
		Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}


// Get Order by ID
func GetOrderByID(c *gin.Context) {
	orderID := c.Param("id")
	var order models.Order

	if err := config.DB.Preload("Items").First(&order, "id = ?", orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// Update Order Status
func UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")
	var order models.Order

	if err := config.DB.First(&order, "id = ?", orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var statusUpdate struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	order.Status = statusUpdate.Status
	order.UpdatedAt = time.Now()
	config.DB.Save(&order)

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated", "status": order.Status})
}

// Delete Order
func DeleteOrder(c *gin.Context) {
	orderID := c.Param("id")
	var order models.Order

	if err := config.DB.First(&order, "id = ?", orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	config.DB.Delete(&order)
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}


// GetUserOrders fetches all orders for a specific user
func GetUserOrders(c *gin.Context) {
	userID := c.Param("userID")
	var orders []models.Order
	if err := config.DB.Where("user_id = ?", userID).Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orders not found for this user"})
		return
	}
	c.JSON(http.StatusOK, orders)
}


func GetOrdersByUser(c *gin.Context) {
	userID := c.Param("user_id") // Get user ID from the request URL

	var orders []models.Order
	if err := config.DB.
		Where("farmer_id = ?", userID).
		Preload("User").              // Load the User associated with the order
		Preload("Items").             // Load the OrderItems in each order
		Preload("Items.Produce").     // Load the Produce details inside each OrderItem
		Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	// If no orders found, return an empty array instead of null
	if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{"orders": []models.Order{}})
		return
	}

	c.JSON(http.StatusOK, orders)
}
