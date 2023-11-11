package db

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}
func DatabaseConnection() {
	app := fiber.New()

	awsConfig := &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}

	sees, err := session.NewSession(awsConfig)
	if err != nil {
		panic(err)
	}
	db := dynamodb.New(sees)

	app.Listen(":3000")

}
