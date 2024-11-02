package dao

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"gorm.io/gorm"
)

type ReportRepositoryImpl struct {
	db *gorm.DB
}

const region = "us-east-1"

var _ repository.ReportRepository = (*ReportRepositoryImpl)(nil)

var (
	fromEmailAddress = "mail@roten-app.com"
)

func NewReportRepository(db *gorm.DB) *ReportRepositoryImpl {
	return &ReportRepositoryImpl{db: db}
}

func (r *ReportRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, report *object.Report) error {
	if err := tx.WithContext(ctx).Create(report).Error; err != nil {
		return fmt.Errorf("failed to insert report: %w", err)
	}
	return nil
}

func (r *ReportRepositoryImpl) SendEmail(ctx context.Context, tx *gorm.DB, subject string, body string) (string, error) {

	var accessKey string
	var awsConfig string
	var secretKey string

	awsConfig = os.Getenv("AWS_CONFIG")
	if awsConfig == "" {
		accessKey = os.Getenv("AWS_ACCESS_KEY_ID")
		secretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	} else {
		accessKey, _ = getEnvVariable("AWS_CONFIG", "AWS_ACCESS_KEY_ID")
		secretKey, _ = getEnvVariable("AWS_CONFIG", "AWS_SECRET_ACCESS_KEY")
	}

	cfg := aws.Config{
		Region:      "us-east-1",
		Credentials: credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
	}

	var toEmailAddress []string
	if err := r.db.WithContext(ctx).Table("admin_users").Select("email").Scan(&toEmailAddress).Error; err != nil {
		return "", fmt.Errorf("failed to find email by admin_users: %w", err)
	}

	client := sesv2.NewFromConfig(cfg)

	input := &sesv2.SendEmailInput{
		FromEmailAddress: &fromEmailAddress,
		Destination: &types.Destination{
			ToAddresses: toEmailAddress,
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Data: &body,
					},
				},
				Subject: &types.Content{
					Data: &subject,
				},
			},
		},
	}

	res, err := client.SendEmail(ctx, input)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return "", err
	}
	fmt.Println("Email message ID:", *res.MessageId)

	return "Email message ID:" + *res.MessageId, nil
}

func getEnvVariable(jsonEnvKey string, key string) (string, error) {
	// Get the environment variable that contains JSON data
	jsonEnv := os.Getenv(jsonEnvKey)
	if jsonEnv == "" {
		return "", errors.New(fmt.Sprintf("Environment variable %s is not configured", jsonEnvKey))
	}

	// Parse the JSON data
	var jsonData map[string]string
	err := json.Unmarshal([]byte(jsonEnv), &jsonData)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Failed to parse JSON from %s: %s", jsonEnvKey, err))
	}

	// Extract the value for the given key
	if val, exists := jsonData[key]; exists {
		return val, nil
	}
	return "", errors.New(fmt.Sprintf("Key %s not found in JSON from %s", key, jsonEnvKey))
}
