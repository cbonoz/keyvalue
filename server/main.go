package main

import (
	"keyvalue-api/constants"
	"keyvalue-api/email"
	"keyvalue-api/models"
	"keyvalue-api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open(constants.DATABASE_URL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate the schema
	db.AutoMigrate(&models.KeyValue{}, &models.APIKey{})
	brevoClient := email.InitBrevo()

	server := routes.NewServer(db, brevoClient)
	r := gin.Default()
	server.RegisterRoutes(r)
	r.Run(":8080")
}
