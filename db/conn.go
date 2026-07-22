package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	// Database connection parameters
	host     = "dpg-d9g0rrnlk1mc73c4u82g-a.ohio-postgres.render.com"
	port     = 5432
	user     = "joel"
	password = "zCY8YOMEdXwaePaVpH88kPVsR4I78FOI"
	dbname   = "curriculum_ch6o"
)

func ConnectDB() (*sql.DB, error) {
	// Create a connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=verify-full",
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
