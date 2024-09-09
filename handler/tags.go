package handler

import (
	"database/sql"
	"fmt"
	db "goapi/database"
	m "goapi/models"
	"strconv"

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
	db := db.Connect_db(false, false)
	row, err := db.QueryContext(c.Context(), query)
	if err != nil {
		fmt.Println("An error occured")
	}

	return row
}

func DeleteTagById(c *fiber.Ctx) interface{} {
	_, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return m.BadRequestError(m.BaseError{Message: "Id param should be a number", Method: c.Method()})
	}
	res, err := db.DeleteTagsById(c)
	if err != nil {
		return m.InternalRequestError(m.BaseError{Message: "Internal server error. Try again in a few minutes", Method: c.Method()})
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return m.InternalRequestError(m.BaseError{Message: "Internal server error. Try again in a few minutes", Method: c.Method()})
	}

	switch rowsAffected {
	case 0:
		errorMessage := fmt.Sprintf("Failed to delete Tag with id: '%s'", c.Params("id"))
		return m.NotFound(m.BaseError{Message: errorMessage, Method: c.Method()})
	default:
		return m.DeleteResponse(m.BaseError{Message: "Resource deleted", Method: c.Method()})
	}
}
func UpdateTagById(c *fiber.Ctx) interface{} {
	bodyRes := &m.PutTags{}
	errParser := c.BodyParser(bodyRes)

	if errParser != nil {
		fmt.Println(errParser)
	}
	_, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return m.BadRequestError(m.BaseError{Message: "Id param should be a string", Method: c.Method()})
	}
	res, errPut := db.PutTagsById(c)
	if errPut != nil {
		return m.InternalRequestError(m.BaseError{Message: errPut.Error(), Method: c.Method()})
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return m.InternalRequestError(m.BaseError{Message: "Internal server error. Try again in a few minutes", Method: c.Method()})
	}

	switch rowsAffected {
	case 0:
		errorMessage := fmt.Sprintf("Failed to update Tag with id: '%s'", c.Params("id"))
		return m.NotFound(m.BaseError{Message: errorMessage, Method: c.Method()})
	default:
		return m.DeleteResponse(m.BaseError{Message: "Resource deleted", Method: c.Method()})
	}
}
