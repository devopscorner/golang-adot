package model

import (
	"time"
)

// --------------------------------------
// SQLite, MySQL, Postgres
// --------------------------------------
// import "gorm.io/gorm"

// ORM: gorm.io/gorm
//      gorm.io/driver/sqlite

// --------------------------------------
// DynamoDB
// --------------------------------------
// import "github.com/guregu/dynamo"

// ORM: github.com/guregu/dynamo

type Book struct {
	// gorm.Model
	ID        string    `json:"id" dynamo:"id" gorm:"primaryKey"`
	Title     string    `json:"title" dynamo:"title" validate:"required"`
	Author    string    `json:"author" dynamo:"author" validate:"required"`
	Year      string    `json:"year" dynamo:"year" validate:"required"`
	CreatedAt time.Time `json:"created_at" dynamo:"created_at"`
	UpdatedAt time.Time `json:"updated_at" dynamo:"updated_at"`
}
