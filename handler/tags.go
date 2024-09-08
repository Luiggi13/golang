package handler

import (
	"database/sql"
	"fmt"
	db "goapi/database"
	m "goapi/models"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func GetByIdTag(c *fiber.Ctx) interface{} {
	var qr m.SelectTags

	rows := db.GetTagByIdHandler(c)
	for rows.Next() {
		err := rows.Scan(&qr.TagId, &qr.TagName, &qr.Public)

		if err == nil {
			return qr
		}
	}

	return m.NotFound(m.BaseError{Message: "Tag not found", Method: c.Method()})
}

func GetTags(c *fiber.Ctx) interface{} {
	var tagList []m.SelectTags = []m.SelectTags{}
	rows := db.GetAllTagsHandler(c)

	for rows.Next() {
		var tag m.SelectTags
		rows.Scan(&tag.TagId, &tag.TagName, &tag.Public)
		tagList = append(tagList, tag)
	}

	return tagList
}

func getTagForPost(c *fiber.Ctx, id string) *sql.Rows {
	query := fmt.Sprintf("select * from tags WHERE id = '%s'", id)
	db := db.Connect_db(false, false, false)
	row, err := db.QueryContext(c.Context(), query)
	if err != nil {
		fmt.Println("An error occured")
	}

	return row
}
