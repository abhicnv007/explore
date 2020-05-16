package cache

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

var client *firestore.Client

// Init is for setting up the client connection
func Init(ctx context.Context) {
	// Sets your Google Cloud Platform project ID.
	projectID := "explore-277406"
	var err error
	client, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
}

// Check if present in cache
func Check(ctx context.Context, q string) (bool, interface{}) {

	doc, err := client.Collection("search_results").Doc(q).Get(ctx)
	if err != nil {
		log.Println(err)
		return false, nil
	}

	return true, doc.Data()

}

// Add item to cache
func Add(ctx context.Context, q string, item interface{}) {
	_, err := client.Collection("search_results").Doc(q).Set(ctx, item)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
}

// Close the client connection
func Close() {
	client.Close()
}
