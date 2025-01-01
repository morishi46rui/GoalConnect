package main

import (
	"database/sql"
	"fmt"
	"log"

	"goalconnect-backend/routes"
	_ "github.com/lib/pq"
)
func main() {
	dsn := "host=db user=admin password=secret dbname=goalconnect port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	fmt.Println("Connected to the database successfully!")

	r := routes.SetupRouter()
	r.Run(":8080")
}
