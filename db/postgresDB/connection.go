package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func DbConnection() (*sql.DB, error) {

	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Infof("error converting port to integer: %v", err)
		return nil, err
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Errorf("failed to connect to database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		log.Errorf("database ping failed: %v", err)
		return nil, err
	}

	log.Info("Connected to PostgreSQL with database/sql!")
	return db, nil
}
