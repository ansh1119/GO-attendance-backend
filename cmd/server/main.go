package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ansh1119/GO-attendance-backend.git/db"
	"github.com/ansh1119/GO-attendance-backend.git/handlers"
	"github.com/ansh1119/GO-attendance-backend.git/repository"
	"github.com/ansh1119/GO-attendance-backend.git/router"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "attendance_db"
	}
	collName := os.Getenv("COLLECTION_NAME")
	if collName == "" {
		collName = "events"
	}
	// connect to Mongo
	client := db.ConnectMongo(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("error disconnecting mongo: %v", err)
		}
	}()

	database := client.Database(dbName)
	// setup repo + handler
	eventRepo := repository.NewEventRepository(database, collName)
	eventHandler := handlers.NewEventHandler(eventRepo)
	r := router.SetupRouter(eventHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
