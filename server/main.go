package main

import (
	"context"
	"keyvalue-api/constants"
	"keyvalue-api/email"
	"keyvalue-api/routes"
	"keyvalue-api/sqlc_generated"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gin-gonic/gin"
)

func main() {

	ctx := context.Background()

	// init db pool
	pool, err := pgxpool.New(ctx, constants.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	// init db queries
	queries := sqlc_generated.New(pool)

	brevoClient := email.InitBrevo()

	server := routes.NewServer(queries, brevoClient)
	r := gin.Default()
	server.RegisterRoutes(r)
	r.Run(":8080")
}
