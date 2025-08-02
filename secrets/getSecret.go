package secrets

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"server.simplifycontrol.com/types"
)

var SecretJSON types.SecretJSONConfig

func FetchSecretFile() error {
	// Check if running locally using environment variable
	isLocal := os.Getenv("ENV") == "local"

	if isLocal {
		// Read from local secrets file
		fileBytes, err := os.ReadFile("secrets.json")
		if err != nil {
			return fmt.Errorf("failed to read local secrets file: %v", err)
		}

		err = json.Unmarshal(fileBytes, &SecretJSON)
		if err != nil {
			return fmt.Errorf("error parsing local JSON: %v", err)
		}

		log.Println("Local secrets loaded successfully")
		return nil
	}

	secretName := "serviceAccountKey"
	region := "us-east-1"

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return fmt.Errorf("failed to fetch aws config: %v", err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to fetch secrets value: %v", err)
	}

	// Decrypts secret using the associated KMS key.
	secretString := *result.SecretString

	// Convert JSON string to a Go map
	err = json.Unmarshal([]byte(secretString), &SecretJSON)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	fmt.Println("Secrets fetched successfully")
	return nil
}
