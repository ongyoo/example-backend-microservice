// service1.go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://root:example@mongo:27017")
	_, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	// GET endpoint for service1
	r.GET("/service1/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from Service 1"})
	})

	// Run server
	r.Run(":8081")
}
