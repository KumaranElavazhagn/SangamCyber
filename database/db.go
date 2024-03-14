package database

import (
	"fmt"
	"sangamCyber/errs"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetDbClient() (*sqlx.DB, *errs.AppError) {
	dataSource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require",
		"sangam-database.cxmgec4i6e3d.us-east-1.rds.amazonaws.com", "sangam", "2010604450", "sangam")

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
