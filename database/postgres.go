package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	port   = 5432
	driver = "postgres"
)

func Connect_db() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open(driver, psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("failed connection!")
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

// func createTable(db *sql.DB) error {
// 	_, err := db.Query("INSERT INTO public.todos (name, lastname) VALUES('christian', 'llansola')")
// 	if err != nil {
// 		log.Fatalln(err)
// 		fmt.Println("An error occured")
// 	}
// 	return nil
// }
