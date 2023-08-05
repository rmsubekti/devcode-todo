package model

import (
	"devcode-todo/helper"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	}

	MYSQL_USER := helper.GetEnv("MYSQL_USER", "root")
	MYSQL_PASS := helper.GetEnv("MYSQL_PASSWORD", "123123")
	MYSQL_HOST := helper.GetEnv("MYSQL_HOST", "127.0.0.1")
	MYSQL_PORT := helper.GetEnv("MYSQL_PORT", "3306")
	MYSQL_DB := helper.GetEnv("MYSQL_DBNAME", "contact-manager")

	var err error
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", MYSQL_USER, MYSQL_PASS, MYSQL_HOST, MYSQL_PORT, MYSQL_DB)
	if db, err = gorm.Open(mysql.Open(conn), &gorm.Config{}); err != nil {
		panic(err)
	}

}
