package database

import (
	"fmt"
	"sangamCyber/errs"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetDbClient() (*sqlx.DB, *errs.AppError) {
	dataSource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "postgres", "Admin@123", "udb")

	client, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		return nil, errs.NewUnexpectedErrorWithMsg("Error opening database connection: " + err.Error())
	}

	if err := client.Ping(); err != nil {
		return nil, errs.NewUnexpectedErrorWithMsg("Error pinging database: " + err.Error())
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client, nil
}

// func GetDbClient() (*sqlx.DB, *errs.AppError) {
// 	dataSource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require",
// 		"db-vps-instance.cbesywa6g0no.eu-north-1.rds.amazonaws.com", "postgres", "Senthil2001#", "vpsdb")

// 	client, err := sqlx.Open("postgres", dataSource)
// 	if err != nil {
// 		return nil, errs.NewUnexpectedErrorWithMsg("Error opening database connection: " + err.Error())
// 	}

// 	if err := client.Ping(); err != nil {
// 		return nil, errs.NewUnexpectedErrorWithMsg("Error pinging database: " + err.Error())
// 	}

// 	client.SetConnMaxLifetime(time.Minute * 3)
// 	client.SetMaxOpenConns(10)
// 	client.SetMaxIdleConns(10)

// 	return client, nil
// }
