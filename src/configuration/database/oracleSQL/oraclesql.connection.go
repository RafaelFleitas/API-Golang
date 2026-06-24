package oraclesql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	_ "github.com/sijms/go-ora/v2"
)

func InitConnection() (*sql.DB, error) {
	host := os.Getenv("ORACLE_HOST")
	port := os.Getenv("ORACLE_PORT")
	user := os.Getenv("ORACLE_USER")
	password := os.Getenv("ORACLE_PASSWORD")
	service := os.Getenv("ORACLE_SERVICE")

	connStr := fmt.Sprintf("oracle://%s:%s@%s:%s/%s?SSL=false", user, password, host, port, service)

	db, err := sql.Open("oracle", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logger.Info("Conseguiu se conectar")

	return db, nil

}
