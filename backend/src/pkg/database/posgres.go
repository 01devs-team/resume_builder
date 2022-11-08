package database

import (
	"fmt"
	"os"
	"resume_builder_backend/src/models"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dsnPattern = "host=$host user=$user password=$password dbname=$dbname port=$port sslmode=$sslmode TimeZone=$TimeZone"

func SetupPostgres() {
	var err error

	retries := 3
	for retries > 0 {
		DB, err = gorm.Open(postgres.Open(connectionString()), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Println("failed to connect database in try " + strconv.Itoa(retries))
		time.Sleep(3 * time.Second)
		retries--
	}

	DB.AutoMigrate(&models.User{})
}

func connectionString() string {
	info := map[string]string{
		"host":     os.Getenv("POSTGRES_HOST"),
		"user":     os.Getenv("POSTGRES_USER"),
		"password": os.Getenv("POSTGRES_PASSWORD"),
		"dbname":   os.Getenv("POSTGRES_DB"),
		"port":     os.Getenv("POSTGRES_PORT"),
		"sslmode":  "disable",
		"TimeZone": "Asia/Tehran",
	}

	dsn := dsnPattern
	for key, value := range info {
		dsn = strings.Replace(dsn, "$"+key, value, -1)
	}
	return dsn
}
