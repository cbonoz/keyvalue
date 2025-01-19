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

	brevoClient := email.InitBrevo()

	server := routes.NewServer(queries, brevoClient)
	r := gin.Default()
	server.RegisterRoutes(r)
	r.Run(":8080")
}
