package server

import (
	"context"
	"fmt"
	"log"
	// "os"
	"time"

	"cloud.google.com/go/storage"
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
