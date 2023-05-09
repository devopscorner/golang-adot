package repository

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/devopscorner/golang-adot/src/config"
	"github.com/devopscorner/golang-adot/src/driver"
	"github.com/devopscorner/golang-adot/src/model"
	"github.com/devopscorner/golang-adot/src/observability"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	// Connect to database
	driver.ConnectDatabase()
}

// GET /books
// Find all books
func GetAll(ctx *gin.Context) []model.Book {
	var books []model.Book

	seg := xray.GetSegment(ctx.Request.Context())
	seg.Name = config.OtelServiceName()

	_, subseg := xray.BeginSubsegment(ctx.Request.Context(), "DynamoDB-GetAll")

	if config.DbConnection() == "dynamo" {
		input := &dynamodb.ScanInput{
			TableName: aws.String(config.DbDatabase()),
		}
		result, _ := driver.DYN.ScanWithContext(ctx.Request.Context(), input)
		err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &books)

		// Add metadata to the subsegment
		subseg.AddMetadata("Table", config.DbDatabase())
		subseg.AddMetadata("Operation", "Scan")

		if err != nil {
			return nil
		}
	} else {
		driver.DB.WithContext(ctx.Request.Context()).Find(&books)

		// Add metadata to the subsegment
		subseg.AddMetadata("Database", config.DbDatabase())
		subseg.AddMetadata("Operation", "Find")
	}

	observability.WriteResponseXRaySegment(seg, ctx.Writer)

	return books
}

// GET /books/:id
// Find a book
func GetByID(ctx *gin.Context, id string) (*model.Book, error) {
	var book model.Book

	seg := xray.GetSegment(ctx.Request.Context())
	seg.Name = config.OtelServiceName()

	_, subseg := xray.BeginSubsegment(ctx.Request.Context(), "DynamoDB-GetByID")

	if config.DbConnection() == "dynamo" {
		input := &dynamodb.GetItemInput{
			TableName: aws.String(config.DbDatabase()),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(id),
				},
			},
		}
		result, err := driver.DYN.GetItemWithContext(aws.BackgroundContext(), input)
		if err != nil {
			return nil, err
		}

		if len(result.Item) == 0 {
			// Item not found in DynamoDB table
			return nil, nil
		}

		err = dynamodbattribute.UnmarshalMap(result.Item, &book)

		// Add metadata to the subsegment
		subseg.AddMetadata("Table", config.DbDatabase())
		subseg.AddMetadata("Operation", "Get")
		subseg.AddMetadata("id", id)

		if err != nil {
			return nil, err
		}
	} else {
		err := driver.DB.First(&book, id).Error

		// Add metadata to the subsegment
		subseg.AddMetadata("Database", config.DbDatabase())
		subseg.AddMetadata("Operation", "Get")
		subseg.AddMetadata("id", id)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}
	}

	observability.WriteResponseXRaySegment(seg, ctx.Writer)

	return &book, nil
}

// POST /books
// Create new book
func Create(ctx *gin.Context, book *model.Book) error {
	t := time.Now().UnixNano()
	id_time := strconv.FormatInt(t, 10)
	book.ID = id_time
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	seg := xray.GetSegment(ctx.Request.Context())
	seg.Name = config.OtelServiceName()

	_, subseg := xray.BeginSubsegment(ctx.Request.Context(), "DynamoDB-Create")
	defer subseg.Close(nil)

	if config.DbConnection() == "dynamo" {
		item, err := dynamodbattribute.MarshalMap(book)
		if err != nil {
			return err
		}

		input := &dynamodb.PutItemInput{
			TableName: aws.String(config.DbDatabase()),
			Item:      item,
		}

		_, err = driver.DYN.PutItemWithContext(ctx.Request.Context(), input)

		// Add metadata to the subsegment
		subseg.AddMetadata("Table", config.DbDatabase())
		subseg.AddMetadata("Operation", "Put")
		subseg.AddMetadata("Item", book)

		if err != nil {
			return err
		}
	} else {
		err := driver.DB.WithContext(ctx.Request.Context()).Create(&book).Error

		// Add metadata to the subsegment
		subseg.AddMetadata("Database", config.DbDatabase())
		subseg.AddMetadata("Operation", "Put")
		subseg.AddMetadata("Item", book)

		if err != nil {
			return err
		}
	}

	observability.WriteResponseXRaySegment(seg, ctx.Writer)

	return nil
}

// PUT /books/:id
// Update a book
func Update(ctx *gin.Context, book *model.Book) error {
	book.UpdatedAt = time.Now()

	seg := xray.GetSegment(ctx.Request.Context())
	seg.Name = config.OtelServiceName()

	_, subseg := xray.BeginSubsegment(ctx.Request.Context(), "DynamoDB-Update")
	defer subseg.Close(nil)

	if config.DbConnection() == "dynamo" {
		item, err := dynamodbattribute.MarshalMap(book)
		if err != nil {
			return err
		}
		input := &dynamodb.PutItemInput{
			Item:      item,
			TableName: aws.String(config.DbDatabase()),
		}
		_, err = driver.DYN.PutItemWithContext(ctx.Request.Context(), input)

		// Add metadata to the subsegment
		subseg.AddMetadata("Table", config.DbDatabase())
		subseg.AddMetadata("Operation", "Put")
		subseg.AddMetadata("Item", book)

		if err != nil {
			return err
		}
	} else {
		err := driver.DB.WithContext(ctx.Request.Context()).Save(&book).Error

		// Add metadata to the subsegment
		subseg.AddMetadata("Database", config.DbDatabase())
		subseg.AddMetadata("Operation", "Put")
		subseg.AddMetadata("Item", book)

		if err != nil {
			return err
		}
	}

	observability.WriteResponseXRaySegment(seg, ctx.Writer)

	return nil
}

// DELETE /books/:id
// Delete a book
func Delete(ctx *gin.Context, id string) error {
	seg := xray.GetSegment(ctx.Request.Context())
	seg.Name = config.OtelServiceName()

	_, subseg := xray.BeginSubsegment(ctx.Request.Context(), "DynamoDB-Delete")
	defer subseg.Close(nil)

	if config.DbConnection() == "dynamo" {
		input := &dynamodb.DeleteItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(id),
				},
			},
			TableName: aws.String(config.DbDatabase()),
		}
		_, err := driver.DYN.DeleteItemWithContext(ctx.Request.Context(), input)

		// Add metadata to the subsegment
		subseg.AddMetadata("Table", config.DbDatabase())
		subseg.AddMetadata("Operation", "Delete")
		subseg.AddMetadata("id", id)

		if err != nil {
			return err
		}
	} else {
		err := driver.DB.WithContext(ctx.Request.Context()).Delete(&model.Book{}, id).Error

		// Add metadata to the subsegment
		subseg.AddMetadata("Database", config.DbDatabase())
		subseg.AddMetadata("Operation", "Delete")
		subseg.AddMetadata("id", id)

		if err != nil {
			return err
		}
	}

	observability.WriteResponseXRaySegment(seg, ctx.Writer)

	return nil
}
