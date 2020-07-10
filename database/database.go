package database

import (
	"fmt"

	"github.com/JustSteveKing/go-ship/config"
	"github.com/jinzhu/gorm"
)

// DBCon is the database connection
var (
	DBCon *gorm.DB
)

// InitDB initialises our database connection
func InitDB() {
	// Get our config
	config := config.New()
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Connection,
		config.Database.Host,
		config.Database.Name,
	)

	var err error

	DBCon, err = gorm.Open("mysql", dsn)

	if err != nil {

		panic("failed to connect database")
	}
}
