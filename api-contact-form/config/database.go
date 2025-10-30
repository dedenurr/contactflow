package config

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DB is the global database connection instance
var DB *gorm.DB

// InitDB initializes the database connection and performs auto-migration
func InitDB(){
	//retrieve database configuration from environment variables
	dbUser := GetEnv("DB_USER", "user")
	dbPassword := GetEnv("DB_PASSWORD", "password")
	dbHost := GetEnv("DB_HOST", "localhost")
	dbPort := GetEnv("DB_PORT", "3306")
	dbName := GetEnv("DB_NAME", "contactsdb")

	// construct the DSN (Data Source Name) for MySQL connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error

	// Open the Database connection using GORM with singular table naming strategy
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// configure connection pool settings
	sqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to get database instance: " + err.Error())
	}

	sqlDB.SetMaxOpenConns(10) // set maximum number of open connections	
	sqlDB.SetMaxIdleConns(5)  // set maximum number of idle connections
	sqlDB.SetConnMaxLifetime(time.Hour) // maximum amount of time a connection may be reused

	// auto-migrate the Contact model to create/update the database table
	if err := DB.AutoMigrate(&model.Contact{}); err != nil {
		panic(fmt.Sprintf("Failed to auto-migrate database: %v", err))
	}
}
