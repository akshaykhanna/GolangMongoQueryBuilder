package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func main() {
	// Rest of the code will go here
	fmt.Println("Hello World")
	client := connectDB()
	collection := client.Database("inventory").Collection("formattypes")
	fmt.Printf("collection: %+v \n", collection)
	disconnectDB(client)
}

func disconnectDB(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	handleError(err)
	fmt.Println("Connection to MongoDB closed.")
}

func connectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	handleError(err)
	err = client.Ping(context.TODO(), nil)
	handleError(err)
	fmt.Println("Connected to MongoDB!")
	return client
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error: ", err)
	}
}
