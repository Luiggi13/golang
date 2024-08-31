package handler

import (
	"encoding/base64"
	"fmt"
	modelQR "goapi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

func CreateQrCode(c *fiber.Ctx) interface{} {
	defaultQr := "iVBORw0KGgoAAAANSUhEUgAABAAAAAQAAQMAAABF07nAAAAABlBMVEX///8AAABVwtN+AAAEFklEQVR42uzdTW7iSBzA0UIsvOQIOQpHM0fjKBwhy15EeDQy4/LfVQY7ymDovN+mJVJtvW2pvpIkSZIkSZIkSZIkSZIkSZIkSZIkSXqtdt1Mp2HIof/hMgz+Sikd+39SSk3/42dKae5bZwAAAAAAAAAAAAAAAIBXB1ymPzY9oOkedhoGpx7wOf3WHgAAAAAAAAAAAAAAAOBNALswk4zf/Lf8Hz6GWem+/9Of/LceEL8FAAAAAAAAAAAAAAAA8KOA2+B9Xj681Q7rhgAAAAAAAAAAAAAAAAB/JaC6W/T8xMkpAAAAAAAAAAAAAAAAwFMAs3tKi8HP3tQKAAAAAAAAAAAAAADwOwD3DvLdroLpN5Ne857S21UwbT83bP+fu2gAAAAAAAAAAAAAAAAAngK4VzM9aHidG1Wblf5AAAAAAAAAAAAAAAAAABsA9mFO+SecN8yAeN6weEJixDyGU4kAAAAAAAAAAAAAAAAA99oVJ/ja4SDfYboM+FVow57S0TbV4+J1QwAAAAAAAAAAAAAAAIDNAal4zD2+/N5PTrvaNaG3wYfpIuPCq2AAAAAAAAAAAAAAAAAANgfEu2hGo9vpwmX8dDyVWBt1XTY5BQAAAAAAAAAAAAAAAKgCmuIyl9rO06+wupgnktdiAyoAAAAAAAAAAAAAAADAywO6sMiX95R2tWtCo+MUfgl7SnfL1w0BAAAAAAAAAAAAAAAANgWMHHnYvbXI6o017fia0Nm1UAAAAAAAAAAAAAAAAICFHcPcML/NUAUUe0qLdcP1AQAAAAAAAAAAAAAAADwZsMvPPhynr0J0d173ezC4GAUAAAAAAAAAAAAAAADw8oDqnZ5N7ZspLEmeKs8LdqsWLgEAAAAAAAAAAAAAAABGjvMwnyvvgOkqjWZ9eU9pXjfsvnvcDwAAAAAAAAAAAAAAAGBTwIPR1XtjuumNMKOlyAQAAAAAAAAAAAAAAADwBoA4erRieRpmpcvPG4auqze1AgAAAAAAAAAAAAAA/FLAvZr8gkPYLbqvPfdQXBO67gUHAAAAAAAAAAAAAAAAgO0Au5mp5X/TzuJ1v4/pntJ43vBz7RMSAAAAAAAAAAAAAAAAAC8BuNRmpbVbZh4BirXQhSunAAAAAAAAAAAAAAAAAJfpQb5mZm4Y1w1TAHTT1cWUr5U5AwAAAAAAAAAAAAAAALwrIOXzhrUp7GhPae4wnErcFZeJAgAAAAAAAAAAAAAAALwtoKudN7x9ug1/y5/9CL+fAQAAAAAAAAAAAAAAAL65p3QEKD4dX3Co7Ty9rlo3BAAAAAAAAAAAAAAAANgWMHcVTFw3LM4bjka1Y+b684YAAAAAAAAAAAAAAAAAmwIkSZIkSZIkSZIkSZIkSZIkSZIkSfqN/RMAAP///PsrTyhvCHUAAAAASUVORK5CYII="
	var inputUrl modelQR.QrInput

	if err := c.BodyParser(&inputUrl); err != nil {
		c.Status(fiber.StatusBadRequest)
		modelQR.ApiBadRequest.Method.SetMethod(c)
		return modelQR.ApiBadRequest
	}
	png, err := qrcode.Encode(inputUrl.Url, qrcode.Highest, 1024)
	if err != nil {
		return defaultQr
	} else {
		encoded := base64.StdEncoding.EncodeToString(png)
		return encoded
	}
}

func CreateTestQrCode(c *fiber.Ctx) interface{} {
	var inputUrl modelQR.QrInput
	if err := c.BodyParser(&inputUrl); err != nil {
		c.Status(fiber.StatusBadRequest)
		modelQR.ApiBadRequest.Method.SetMethod(c)
		return modelQR.ApiBadRequest
	}
	fmt.Println(inputUrl)
	return 200
}
