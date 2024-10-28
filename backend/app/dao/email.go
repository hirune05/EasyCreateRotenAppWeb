package dao

import (
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

var _ repository.ReportRepository = (*ReportRepositoryImpl)(nil)

type ReportRepositoryImpl struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepositoryImpl {
	return &ReportRepositoryImpl{db: db}
}

const region = "us-east-1"

var (
	fromEmailAddress = "mail@roten-app.com"
	// toEmailAddress   = "fkys2932@gmail.com"
	toEmailAddress = "ii440092@icloud.com"
	subject        = "タイトル"
	body           = "本文"
)

func (r *ReportRepositoryImpl) SendEmail(ctx context.Context, tx *gorm.DB) (string, error) {

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

	client := sesv2.NewFromConfig(cfg)

	input := &sesv2.SendEmailInput{
		FromEmailAddress: &fromEmailAddress,
		Destination: &types.Destination{
			ToAddresses: []string{toEmailAddress},
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
