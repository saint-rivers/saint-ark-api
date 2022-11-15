package utils

import (
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func SaveFile(ctx *fiber.Ctx) (string, string, string, error) {
	file, err := ctx.FormFile("myFile")

	if err != nil { /* handle error */
		return "error", "id", "filename", err
	}

	// get file extension
	newFileName := rename(file.Filename)
	ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", newFileName))

	fmt.Println("Successfully Uploaded File")
	createdUrl := fmt.Sprintf("http://localhost:8080/f/%s", newFileName)

	return createdUrl, newFileName, file.Filename, nil
}

func rename(name string) string {
	ext := filepath.Ext(name)
	return "upload-" + GenerateAlphanumericString(24) + ext
}
