package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {

	db, err := sql.Open(
		"mysql",
		"root:@tcp(127.0.0.1:3306)/belajar_golang",
	)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Koneksi database berhasil")

	return db, nil
}