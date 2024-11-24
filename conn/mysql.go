package conn

import (
	"TobiasFP/BotNana/config"
	"time"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetMysqlDB returns a pointer to the standard sql db.
func GetMysqlDB() (*gorm.DB, error) {
	config := config.GetConfig()

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	server := config.GetString("mysql.DB_SERVER")
	userName := config.GetString("mysql.DB_USERNAME")
	password := config.GetString("mysql.DB_PASSWORD")
	database := config.GetString("mysql.DB_DATABASE")
	cloudsql := config.GetBool("mysql.CLOUDSQL")
	protocol := "tcp"
	if cloudsql {
		protocol = "cloudsql"
	}
	dsn := userName + ":" + password + "@" + protocol + "(" + server + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	sqldb, err := db.DB()
	if err != nil {
		return db, err
	}

	sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(0)
	sqldb.SetConnMaxLifetime(time.Minute)
	return db, nil
}
