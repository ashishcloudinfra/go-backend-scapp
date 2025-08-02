package helpers

import (
	"context"
	"fmt"

	vision "cloud.google.com/go/vision/apiv1"
	"google.golang.org/api/option"
)

func DetectTextURI(file string) (string, error) {
	ctx := context.Background()

	// Create Vision API client
	client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile("serviceAccountKey.json"))
	if err != nil {
		return "", fmt.Errorf("failed to create Vision API client: %w", err)
	}
	defer client.Close()

	// Load image from URI
	image := vision.NewImageFromURI(file)

	// Perform text detection (max results: 10)
	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return "", fmt.Errorf("failed to detect text: %w", err)
	}

	// If no text is found
	if len(annotations) == 0 {
		return "", nil
	}

	// Extract detected text
	var detectedText string
	for _, annotation := range annotations {
		detectedText += annotation.Description + "\n"
	}

	return detectedText, nil
}
