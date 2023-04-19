package driver

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/devopscorner/golang-adot/src/config"
	"github.com/guregu/dynamo"
)

var (
	DB_Dynamo *dynamo.DB
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

	database := dynamo.New(sess, &aws.Config{
		Region: aws.String(config.AWSRegion()),
		Credentials: credentials.NewStaticCredentials(
			config.AWSAccessKey(),
			config.AWSSecretKey(),
			"",
		),
	})

	database.Table(config.DbDatabase())

	DB_Dynamo = database
}
