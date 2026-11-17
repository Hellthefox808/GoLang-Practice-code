// Package main initializes the MongoDB client, provisions timeouts,
// configures environment variables, and launches the HTTP service.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Hellthefox808/GoLang-Practice-code/25_mongodb_crud_api/controller"
	"github.com/Hellthefox808/GoLang-Practice-code/25_mongodb_crud_api/router"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	fmt.Println("=== 25: Production MongoDB CRUD Microservice ===")

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
		fmt.Println("Info: MONGO_URI unset. Defaulting to local instance: mongodb://localhost:27017")
	}

	dbName := os.Getenv("MONGO_DB")
	if dbName == "" {
		dbName = "netflix_watchlist"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// 1. Establish database connection with bounded context deadline
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)

	var collection *mongo.Collection
	if err != nil {
		fmt.Printf("Warning: Could not connect to MongoDB (%v). Running in degraded/mock mode.\n", err)
	} else {
		// Ping database to verify active liveness
		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			fmt.Printf("Warning: MongoDB ping failed (%v). Ensure database instance is running.\n", err)
		} else {
			fmt.Println("MongoDB cluster connected and verified successfully.")
			collection = client.Database(dbName).Collection("movies")
		}
	}

	// 2. Wire up Controller and Router
	movieController := controller.NewMovieController(collection)
	r := router.SetupRouter(movieController)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(server.ListenAndServe())
}
