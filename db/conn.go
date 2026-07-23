package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	// Database connection parameters
	host     = "aws-0-sa-east-1.pooler.supabase.com"
	port     = 6543
	user     = "postgres.fmdbezyxzpozkupibykm"
	password = "r5i9BxqNxRKtP8Ym"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	// Create a connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	// Open a connection to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to" + dbname)
	return db, nil
}
