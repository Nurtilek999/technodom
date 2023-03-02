package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func InitDB() (*sql.DB, error) {
	username := viper.GetString("Database.Username")
	password := viper.GetString("Database.Password")
	dbName := viper.GetString("Database.DBName")
	sslmode := viper.GetString("Database.Sslmode")

	connectionString := fmt.Sprintf("user = %s password = %s dbname = %s  port = %s sslmode = %s", username, password, dbName, ":5431", sslmode)
	DB, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Connection error ", err.Error())
	}
	return DB, nil
}
