package firebase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
	"server.simplifycontrol.com/secrets"
)

// Initialize Firebase Storage Client
func getStorageClient() (*storage.Client, error) {
	ctx := context.Background()
	jsonData, err := json.MarshalIndent(secrets.SecretJSON.FirebaseConfig, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil, err
	}
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(jsonData))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func uploadFileToFirebase(client *storage.Client, bucketName, fileName string, file io.Reader) (string, error) {
	ctx := context.Background()
	bucket := client.Bucket(bucketName)
	obj := bucket.Object(fileName)

	// Open file writer
	wc := obj.NewWriter(ctx)
	wc.ContentType = "image/jpeg" // Change based on file type
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	// Make the file publicly accessible
	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", err
	}

	// Return the public URL of the uploaded file
	publicURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, fileName)
	return publicURL, nil
}

func DeleteFileFromFirebase(client *storage.Client, bucketName, fileName string) error {
	ctx := context.Background()
	bucket := client.Bucket(bucketName)
	obj := bucket.Object(fileName)

	if err := obj.Delete(ctx); err != nil {
		return err
	}
	return nil
}

func UploadToStorageAndGetPublicLinks(files []*multipart.FileHeader) ([]string, error) {
	// Initialize Firebase Storage client
	client, err := getStorageClient()
	if err != nil {
		fmt.Println("Firebase client error:", err)
		return nil, err
	}
	defer client.Close()

	bucketName := "sc-server-40e33.firebasestorage.app"

	var menuImageLinks []string

	// Upload each file
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// Generate a unique filename
		fileName := fmt.Sprintf("menus/%d-%s", time.Now().UnixNano(), fileHeader.Filename)

		// Upload to Firebase
		publicURL, err := uploadFileToFirebase(client, bucketName, fileName, file)
		if err != nil {
			fmt.Println("Upload error:", err)
			return nil, err
		}

		menuImageLinks = append(menuImageLinks, publicURL)
	}

	return menuImageLinks, nil
}
