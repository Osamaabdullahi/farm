package config

import (
	"farmconnect/models"
	"fmt"
	"log"
	"os"

	"github.com/lpernett/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using system environment variables.")
	}

	// Retrieve DATABASE from environment (Render provides this)
	dsn := os.Getenv("DATABASE")
	if dsn == "" {
		log.Fatal("DATABASE environment variable is not set!")
	}
	
	// Initialize the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection established")

	// Enable the uuid-ossp extension
	err = DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	if err != nil {
		panic("failed to create uuid-ossp extension: " + err.Error())
	}

	// Ensure ENUM types exist before migration
	err = CreateEnumTypes(DB)
	if err != nil {
		log.Fatal("Failed to create ENUM types:", err)
	}


	// AutoMigrate models
	err = DB.AutoMigrate(&models.User{},&models.Farmer{},&models.Business{},&models.Logistics{},&models.Produce{},&models.Order{},&models.OrderItem{})
	if err != nil {
		log.Fatal("Failed to migrate models:", err)
	}

	log.Println("Database connected successfully!")

}




// CreateEnumTypes ensures ENUM types exist before migration
func CreateEnumTypes(db *gorm.DB) error {
	// Create user_role ENUM if not exists
	err := db.Exec(`
		DO $$ 
		BEGIN 
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN 
				CREATE TYPE user_role AS ENUM ('farmer', 'business', 'logistics', 'admin'); 
			END IF; 
		END $$;
	`).Error
	if err != nil {
		return err
	}

	// Create user_status ENUM if not exists
	err = db.Exec(`
		DO $$ 
		BEGIN 
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_status') THEN 
				CREATE TYPE user_status AS ENUM ('active', 'pending', 'banned'); 
			END IF; 
		END $$;
	`).Error
	if err != nil {
		return err
	}

	return nil
}
