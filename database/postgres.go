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

func getPostgresConfig(host, user, password, dbname string, port int) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

func setEnvironment() string {
	var host, user, password, dbname string

	if os.Getenv("ENVIRONMENT") == "development" {
		host = os.Getenv("DEV_POSTGRES_HOST")
		user = os.Getenv("DEV_POSTGRES_USER")
		password = os.Getenv("DEV_POSTGRES_PASSWORD")
		dbname = os.Getenv("DEV_POSTGRES_DB")
	} else if os.Getenv("ENVIRONMENT") == "live" {
		host = os.Getenv("POSTGRES_HOST")
		user = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname = os.Getenv("POSTGRES_DB")
	}
	fmt.Println(host, user, password, dbname, port)
	return getPostgresConfig(host, user, password, dbname, port)
}
func Connect_db(migration bool, seed bool, clean bool) *sql.DB {
	db, err := sql.Open(driver, setEnvironment())
	if err != nil {
		panic(err)
	}
	if migration {
		err = MakeMigrationStructure(db)
		if err != nil {
			fmt.Println("err migration")
		}
		defer db.Close()
	}

	if seed {
		err = SeedInitialData(db)
		if err != nil {
			fmt.Println("err seed")
		}
		defer db.Close()
	}

	if clean {
		err = CleanTables(db)
		if err != nil {
			fmt.Println("err clean")
		}
		defer db.Close()
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("failed connection!")
		panic(err)
	}

	return db
}

func InsertQR(d m.QRStruct) error {
	db := Connect_db(false, false, false)
	_, err := db.Exec("INSERT INTO public.qrs (url_text, qr_code, userid,premium,created_at) VALUES($4, $1, $2, $3, now());", d.QrCode, d.User, d.Premium, d.UrlText)
	if err != nil {
		log.Fatalln(err)
		fmt.Println("An error occured")
	}
	defer db.Close()
	return nil
}

func GetAll(c *fiber.Ctx) *sql.Rows {
	db := Connect_db(false, false, false)
	queryJoin := "select qrs.id, qrs.qr_code, qrs.userid, qrs.url_text, qrs.premium, tags.name AS tag_name from qrs left join tags on qrs.id_tag = tags.id order by qrs.id asc"
	row, err := db.Query(queryJoin)
	if err != nil {
		fmt.Println("An error occured")
	}

	return row
}

func GetById(c *fiber.Ctx) *sql.Rows {
	query := fmt.Sprintf("select qr_code, userid, premium from qrs WHERE id = %s", c.Params("id"))
	db := Connect_db(false, false, false)
	row, err := db.QueryContext(c.Context(), query)
	if err != nil {
		fmt.Println("An error occured")
	}

	return row
}

func DeleteById(c *fiber.Ctx) (sql.Result, error) {
	query := fmt.Sprintf("DELETE FROM qrs WHERE id = %s", c.Params("id"))
	db := Connect_db(false, false, false)
	deleted, err := db.ExecContext(c.Context(), query)
	if err != nil {
		fmt.Println("An error occured")
	}

	defer db.Close()
	return deleted, nil
}

func MakeMigrationStructure(db *sql.DB) error {
	b, err := os.ReadFile("./database/sqls/structure.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}
	b2, err := os.ReadFile("./database/sqls/tags.sql")
	if err != nil {
		return err
	}

	rows2, err := db.Query(string(b2))
	if err != nil {
		return err
	}
	rows.Close()
	return rows2.Close()
}

func SeedInitialData(db *sql.DB) error {
	b, err := os.ReadFile("./database/sqls/inserts.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}

func CleanTables(db *sql.DB) error {
	b, err := os.ReadFile("./database/sqls/clean.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}
