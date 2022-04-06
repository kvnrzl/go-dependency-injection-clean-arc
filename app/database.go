package app

import (
	"database/sql"
	"golang_restful_api/helper"
	"time"
)

func NewDB() *sql.DB {
	DB, errDB := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_restful_api?parseTime=true")
	helper.PanicIfError(errDB)

	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(20)
	DB.SetConnMaxIdleTime(10 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)

	return DB
}
