package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	// "google.golang.org/api/iterator"
)

func AddImageToBucket(imageData []byte) (string, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return "", err
	}
	defer client.Close()

	bucketName := "vicyberbucket"
	bucket := client.Bucket(bucketName)
	
	// Create a unique filename (you can use timestamps or UUIDs to ensure uniqueness)
	objectName := fmt.Sprintf("images/%d.jpg", time.Now().Unix())

	// Create an object in the bucket and upload the image data
	object := bucket.Object(objectName)
	writer := object.NewWriter(ctx)
	writer.ContentType = "image/jpeg" // Set the appropriate content type

	// Write the image data to the storage object
	if _, err := writer.Write(imageData); err != nil {
		log.Fatalf("Failed to write data to bucket: %v", err)
		return "", err
	}

	// Close the writer to finalize the upload
	if err := writer.Close(); err != nil {
		log.Fatalf("Failed to close writer: %v", err)
		return "", err
	}

	// Generate the public URL of the uploaded image
	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName)
	log.Println("Image uploaded successfully. URL:", url)

	return url, nil
}

func GetImageFromBucket(imageUrl string) ([]byte, error) {
	// Send an HTTP GET request to fetch the image
	resp, err := http.Get(imageUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch image: %w", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch image: status code %d", resp.StatusCode)
	}

	// Read the image data
	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read image data: %w", err)
	}

	return imageData, nil
}

func DeleteImageFromBucket(imageURL string) error {
	// Extract the object path from the URL
	objectPath := "images/" + imageURL[strings.LastIndex(imageURL, "/")+1:]
	// log.Println("Object Path for deletion:", objectPath)

	// Initialize the Google Cloud Storage client
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Println("Error creating storage client:", err)
		return err
	}
	defer client.Close()

	// Create the object handle for the specific object path
	object := client.Bucket("vicyberbucket").Object(objectPath)

	// Attempt to delete the image from the bucket
	if err := object.Delete(context.Background()); err != nil {
		log.Println("Failed to delete image from bucket:", err)
		return err
	}

	// log.Println("Image deleted successfully from bucket:", objectPath)
	return nil
}

// func ListBucketObjects() {
//     client, err := storage.NewClient(context.Background())
//     if err != nil {
//         log.Fatalf("Failed to create client: %v", err)
//     }
//     defer client.Close()

//     bucket := client.Bucket("vicyberbucket")
//     query := &storage.Query{}
//     it := bucket.Objects(context.Background(), query)

//     for {
//         obj, err := it.Next()
//         if err == iterator.Done {
//             break
//         }
//         if err != nil {
//             log.Fatalf("Failed to list objects: %v", err)
//         }
//         fmt.Printf("Object: %s\n", obj.Name)
//     }
// }
