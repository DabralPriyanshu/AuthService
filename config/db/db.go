package config

import (
	env "Auth/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = env.GetString("DB_USER", "root")
	cfg.Passwd = env.GetString("DB_PASS", "root")
	cfg.Net = env.GetString("DB_NET", "tcp")
	cfg.Addr = env.GetString("DB_ADDR", "localhost")
	cfg.DBName = env.GetString("DB_NAME", "airbnb_auth_db")
	fmt.Println("Connecting to database ", cfg.DBName, cfg.FormatDSN())
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Error connecting DB!!!", err)
		return nil, err
	}
	fmt.Println("Trying to connect to database")
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Error pinging database: ", pingErr)
		return nil, pingErr

	}
	fmt.Println("Connected to database : ", cfg.DBName)
	return db, nil

}
