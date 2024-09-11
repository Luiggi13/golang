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

func Connect_db(migration bool, clean bool) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open(driver, psqlInfo)
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
	db := Connect_db(false, false)
	_, err := db.Exec("INSERT INTO public.qrs (qr_code, userid, url_text, premium, id_tag, created_at) VALUES($1, $2, $3, $4, $5, now());", d.QrCode, d.UserId, d.UrlText, d.Premium, d.IdTag)
	if err != nil {
		log.Fatalln(err)
		fmt.Println("An error occured")
	}
	defer db.Close()
	return nil
}

func GetAll(c *fiber.Ctx) *sql.Rows {
	db := Connect_db(false, false)
	queryJoin := "select qrs.id, qrs.qr_code, qrs.id_user, qrs.url_text, qrs.premium, tags.name AS tag_name from qrs inner join tags on qrs.id_tag = tags.id order by qrs.id asc"
	row, err := db.Query(queryJoin)
	if err != nil {
		fmt.Println("An error occured")
	}

	return row
}

func GetById(c *fiber.Ctx) *sql.Rows {
	query := fmt.Sprintf("select qrs.id, qrs.qr_code, qrs.userid, qrs.url_text, qrs.premium, tags.name AS tag_name from qrs inner join tags on qrs.id_tag = tags.id where qrs.id =  %s order by qrs.id desc", c.Params("id"))

	db := Connect_db(false, false)
	row, err := db.QueryContext(c.Context(), query)
	if err != nil {
		fmt.Println("An error occured")
	}

	return row
}

func GetTagByIdHandler(c *fiber.Ctx) *sql.Rows {
	query := fmt.Sprintf("select * from tags WHERE id = '%s'", c.Params("id"))
	db := Connect_db(false, false)
	row, err := db.QueryContext(c.Context(), query)
	if err != nil {
		fmt.Println("An error occured")
	}

	return row
}

func GetAllTagsHandler(c *fiber.Ctx) *sql.Rows {
	db := Connect_db(false, false)
	query := "select * from tags order by id asc"
	row, err := db.Query(query)
	if err != nil {
		fmt.Println("An error occured")
	}

	return row
}

func DeleteById(c *fiber.Ctx) (sql.Result, error) {
	query := fmt.Sprintf("DELETE FROM qrs WHERE id = %s", c.Params("id"))
	db := Connect_db(false, false)
	deleted, err := db.ExecContext(c.Context(), query)
	if err != nil {
		fmt.Println("An error occured")
	}

	defer db.Close()
	return deleted, nil
}

func DeleteTagsById(c *fiber.Ctx) (sql.Result, error) {
	query := fmt.Sprintf("DELETE FROM tags WHERE id = '%s'", c.Params("id"))
	db := Connect_db(false, false)
	deleted, err := db.ExecContext(c.Context(), query)
	if err != nil {
		fmt.Println("An error occured")
	}

	defer db.Close()
	return deleted, nil
}

func PutTagsById(c *fiber.Ctx) (sql.Result, error) {
	bodyRes := &m.PutTags{}
	err := c.BodyParser(bodyRes)
	if err != nil {
		return nil, fmt.Errorf("error parsing body: %v", err)
	}

	if bodyRes.TagName == "" {
		return nil, fmt.Errorf("the 'name' field is required and must not be empty")
	}
	if bodyRes.Public == nil {
		return nil, fmt.Errorf("the 'public' field is required and must not be empty")
	}
	query := fmt.Sprintf("UPDATE tags SET name='%s', public=%v WHERE id = '%s'", bodyRes.TagName, bodyRes.Public, c.Params("id"))
	db := Connect_db(false, false)
	updatedRow, errQuery := db.ExecContext(c.Context(), query)

	if errQuery != nil {
		return nil, fmt.Errorf("error updating database: %v", errQuery)
	}

	defer db.Close()
	return updatedRow, nil
}

func MakeMigrationStructure(db *sql.DB) error {
	structureSQL, structureSQLError := os.ReadFile("./database/sqls/structure.sql")
	if structureSQLError != nil {
		return structureSQLError
	}

	rowsStructure, rowsStructureError := db.Query(string(structureSQL))
	if rowsStructureError != nil {
		return rowsStructureError
	}

	tagsSql, tagsSqlError := os.ReadFile("./database/sqls/inserts.sql")
	if tagsSqlError != nil {
		return tagsSqlError
	}

	rowsTagsSQL, rowsTagsSQLError := db.Query(string(tagsSql))
	if rowsTagsSQLError != nil {
		return rowsTagsSQLError
	}

	rowsStructure.Close()
	return rowsTagsSQL.Close()
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
