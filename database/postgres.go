package database

import (
	"database/sql"
	"fmt"
	m "goapi/models"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	port   = 5432
	driver = "postgres"
)

func Connect_db() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open(driver, psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("failed connection!")
		panic(err)
	}

	return db
}

func InsertQR(d m.InsertQrDB) error {
	db := Connect_db()
	_, err := db.Exec("INSERT INTO public.qrs (qr_code, userid,premium,created_at) VALUES($1, $2, $3, now());", d.QrCode, d.User, d.Premium)
	if err != nil {
		log.Fatalln(err)
		fmt.Println("An error occured")
	}
	defer db.Close()
	return nil
}
