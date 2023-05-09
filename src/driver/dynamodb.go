package driver

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/devopscorner/golang-adot/src/config"
)

var (
	DB_Dynamo *dynamodb.DynamoDB
)

func ConnectDynamo() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region:     aws.String(config.AWSRegion()),
		DisableSSL: aws.Bool(true),
	}))

	database := dynamodb.New(sess, &aws.Config{
		Region: aws.String(config.AWSRegion()),
		Credentials: credentials.NewStaticCredentials(
			config.AWSAccessKey(),
			config.AWSSecretKey(),
			"",
		),
	})

	xray.AWS(database.Client)

	DB_Dynamo = database
}
