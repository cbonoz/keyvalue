package main

import (
	"context"
	"database/sql"
	"keyvalue-api/constants"
	"keyvalue-api/db"
	"keyvalue-api/email"
	"keyvalue-api/routes"
	"keyvalue-api/sqlc_generated"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/gin-gonic/gin"
)

func main() {

	ctx := context.Background()
	err := constants.Init()
	if err != nil {
		log.Fatal(err)
	}

	// init db pool
	pool, err := pgxpool.New(ctx, constants.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	// init db queries
	queries := sqlc_generated.New(pool)
	cp := pool.Config().ConnConfig.ConnString()
	dbConn, err := sql.Open("pgx", cp)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrate(dbConn)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	brevoClient := email.InitBrevo()

	server := routes.NewServer(queries, brevoClient)
	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Update with your Vite server URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	r.Use(cors.New(corsConfig))
	server.RegisterRoutes(r)
	log.Default().Println("Server started on port " + constants.SERVER_PORT)
	r.Run(":" + constants.SERVER_PORT)
}
