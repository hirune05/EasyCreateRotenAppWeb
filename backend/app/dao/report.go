package dao

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
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

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	if accessKey == "" {
		return "", errors.New("AWS_ACCESS_KEY_ID is not configured")
	}

	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if secretKey == "" {
		return "", errors.New("AWS_SECRET_ACCESS_KEY is not configured")
	}

	cfg := aws.Config{
		Region:      "us-east-1",
		Credentials: credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
	}

	var toEmailAddress []string
        if err :=  r.db.WithContext(ctx).Table("admin_users").Select("email").Scan(&toEmailAddress).Error; err != nil {
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
