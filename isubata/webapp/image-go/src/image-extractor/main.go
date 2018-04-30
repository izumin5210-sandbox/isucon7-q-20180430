package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	os.Exit(run())
}

func run() int {
	fileRoot := getEnvOrDefault("ISUBATA_FILE_ROOT", "./data")

	db, err := connect()
	if err != nil {
		log.Println("failed to connect with DB: %v", err)
		return 1
	}

	rows, err := db.Query("select name, data from image")
	if err != nil {
		log.Println("failed to connect with DB: %v", err)
		return 1
	}
	defer rows.Close()

	for rows.Next() {
		var (
			name string
			data []byte
		)
		err := rows.Scan(&name, &data)
		if err != nil {
			log.Println("failed to scan row: %v", err)
			return 1
		}
		path := filepath.Join(fileRoot, "icons", name)
		fmt.Println(path)
		ioutil.WriteFile(path, data, 0644)
	}

	return 0
}

func connect() (*sql.DB, error) {
	host := getEnvOrDefault("ISUBATA_DB_HOST", "127.0.0.1")
	port := getEnvOrDefault("ISUBATA_DB_PORT", "3306")
	username := getEnvOrDefault("ISUBATA_DB_USER", "root")
	password := getEnvOrDefault("ISUBATA_DB_PASSWORD", "")
	if password != "" {
		password = ":" + password
	}
	database := "isubata"
	addr := fmt.Sprintf("%s%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4", username, password, host, port, database)
	return sql.Open("mysql", addr)
}

func getEnvOrDefault(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return defaultValue
}
