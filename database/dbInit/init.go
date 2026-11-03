package init

import (
	"database/sql"
	"fmt"
	"os"

	u "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Config() (*sql.DB, error) {
	 loadErr := u.Load()
	if loadErr != nil {
		return nil,loadErr
		
	}
	
	dbCreds := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	
	db, sqlErr := sql.Open("postgres", dbCreds)	
	if sqlErr != nil {
		return nil, sqlErr
	}
	return db, nil

}


