package database

import (
	"database/sql"
	"fmt"
	m "goapi/models"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
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

	err = db.Ping()
	if err != nil {
		fmt.Println("failed connection!")
		panic(err)
	}

	return db
}

func InsertQR(d m.QRStruct) error {
	db := Connect_db()
	_, err := db.Exec("INSERT INTO public.qrs (url_text, qr_code, userid,premium,created_at) VALUES($4, $1, $2, $3, now());", d.QrCode, d.User, d.Premium, d.UrlText)
	if err != nil {
		log.Fatalln(err)
		fmt.Println("An error occured")
	}
	defer db.Close()
	return nil
}

func GetAll(c *fiber.Ctx) *sql.Rows {
	db := Connect_db()
	// rows.Scan(&qr.QrCode, &qr.User, &qr.Premium, &qr.UrlText)
	row, err := db.QueryContext(c.Context(), "select qr_code, userid, premium, url_text from qrs")
	if err != nil {
		fmt.Println("An error occured")
	}

	return row
}

func GetById(c *fiber.Ctx) *sql.Rows {
	query := fmt.Sprintf("select qr_code, userid, premium from qrs WHERE id = %s", c.Params("id"))
	db := Connect_db()
	row, err := db.QueryContext(c.Context(), query)
	if err != nil {
		fmt.Println("An error occured")
	}

	return row
}

func DeleteById(c *fiber.Ctx) (sql.Result, error) {
	query := fmt.Sprintf("DELETE FROM qrs WHERE id = %s", c.Params("id"))
	db := Connect_db()
	deleted, err := db.ExecContext(c.Context(), query)
	if err != nil {
		fmt.Println("An error occured")
		// return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to delete QR code")
	}

	defer db.Close()
	return deleted, nil
}
