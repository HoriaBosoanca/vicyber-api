package server

// import (
// 	"log"
// 	"fmt"
// 	"context"

// 	// google bucket
// 	"cloud.google.com/go/storage"
// 	"google.golang.org/api/iterator"
// )

func AddImageToBucket(base64image string) string {
	return "bucket url"
}

// func Bucket() {
// 	ctx := context.Background()

//     // Initialize client
//     client, err := storage.NewClient(ctx)
//     if err != nil {
//         log.Fatalf("Failed to create client: %v", err)
//     }
//     defer client.Close()

//     // Define bucket name
//     bucketName := "vicyberbucket"

//     // List objects in the bucket
//     it := client.Bucket(bucketName).Objects(ctx, nil)
//     for {
//         objAttrs, err := it.Next()
//         if err == iterator.Done {
//             break
//         }
//         if err != nil {
//             log.Fatalf("Error listing objects: %v", err)
//         }
//         fmt.Println("Object name:", objAttrs.Name)
//     }
// }
