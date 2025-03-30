package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Claims struct for JWT payload
type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new JWT token
func GenerateToken(userID, role string) (string, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken validates the JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}


	return claims, nil
}



// func CreateFarmer(c *gin.Context) {
// 	// Parse form data first
// 	err := c.Request.ParseMultipartForm(10 << 20) // 10MB limit
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
// 		return
// 	}

// 	// Extract JSON fields from form data
// 	req := struct {
// 		UserID           uuid.UUID `form:"user_id" binding:"required"`
// 		FarmName         string    `form:"farm_name" binding:"required"`
// 		FarmLocation     string    `form:"farm_location" binding:"required"`
// 		TypesOfCrops     string    `form:"types_of_crops" binding:"required"`
// 		HarvestFrequency string    `form:"harvest_frequency" binding:"required"`
// 		VerificationID   string    `form:"verification_id" binding:"required"`
// 		PreferredPayment string    `form:"preferred_payment" binding:"required"`
// 		PhoneNumber      string    `form:"phone_number" binding:"required"`
// 	}{}

// 	// Bind form values
// 	if err := c.ShouldBind(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Check if user exists
// 	var user models.User
// 	if err := config.DB.Where("id = ?", req.UserID).First(&user).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	// Handle multiple file uploads
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
// 		return
// 	}

// 	files := form.File["farm_photos"]
// 	var photoURLs []string

// 	for _, file := range files {
// 		// Save file locally (or replace with cloud storage logic)
// 		filePath := "uploads/" + uuid.New().String() + "_" + file.Filename
// 		err := c.SaveUploadedFile(file, filePath)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
// 			return
// 		}

// 		// Store file path (update this for cloud storage)
// 		photoURLs = append(photoURLs, filePath)
// 	}

// 	// Create a new Farmer
// 	farmer := models.Farmer{
// 		ID:               uuid.New(),
// 		UserID:           req.UserID,
// 		FarmName:         req.FarmName,
// 		FarmLocation:     req.FarmLocation,
// 		TypesOfCrops:     req.TypesOfCrops,
// 		HarvestFrequency: req.HarvestFrequency,
// 		VerificationID:   req.VerificationID,
// 		PreferredPayment: req.PreferredPayment,
// 		PhoneNumber:      req.PhoneNumber,
// 		FarmPhotos:       photoURLs, // Store as CSV (or use a JSON array in DB)
// 		CreatedAt:        time.Now(),
// 		UpdatedAt:        time.Now(),
// 	}

// 	// Save farmer profile
// 	if err := config.DB.Create(&farmer).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create farmer profile"})
// 		return
// 	}

// 	// Update onboarding status for user
// 	config.DB.Model(&user).Update("onboarding", true)

// 	c.JSON(http.StatusCreated, gin.H{"message": "Farmer profile created successfully", "farmer": farmer})
// }


// Create a new farmer profile
// func CreateFarmer(c *gin.Context) {
// 	var req struct {
// 		UserID           uuid.UUID `json:"user_id"`
// 		FarmName         string    `json:"farm_name"`
// 		FarmLocation     string    `json:"farm_location"`
// 		TypesOfCrops     string    `json:"types_of_crops"`
// 		HarvestFrequency string    `json:"harvest_frequency"`
// 		VerificationID   string    `json:"verification_id"`
// 		PreferredPayment string    `json:"preferred_payment"`
// 		PhoneNumber      string    `json:"phone_number"`
// 		FarmPhotos       []string  `json:"farm_photos"` // Accept array of image URLs
// 	}

// 	// Parse JSON request
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Check if user exists
// 	var user models.User
// 	if err := config.DB.Where("id = ?", req.UserID).First(&user).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	// Create a new Farmer
// 	farmer := models.Farmer{
// 		ID:               uuid.New(),
// 		UserID:           req.UserID,
// 		FarmName:         req.FarmName,
// 		FarmLocation:     req.FarmLocation,
// 		TypesOfCrops:     req.TypesOfCrops,
// 		HarvestFrequency: req.HarvestFrequency,
// 		VerificationID:   req.VerificationID,
// 		PreferredPayment: req.PreferredPayment,
// 		PhoneNumber:      req.PhoneNumber,
// 		FarmPhotos:       req.FarmPhotos, // Save multiple photo URLs
// 		CreatedAt:        time.Now(),
// 		UpdatedAt:        time.Now(),
// 	}

// 	// Save farmer profile
// 	if err := config.DB.Create(&farmer).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create farmer profile"})
// 		return
// 	}

// 	// Update onboarding status for user
// 	config.DB.Model(&user).Update("onboarding", true)

// 	c.JSON(http.StatusCreated, gin.H{"message": "Farmer profile created successfully", "farmer": farmer})
// }
