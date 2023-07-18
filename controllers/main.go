package controllers

import (
	"fmt"
	"keep-api/database"
	"keep-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetNotes(c *fiber.Ctx) error {
	notes := models.GetNotes{}
	db := database.DB.Db
	result := db.Model(models.Note{}).Find(&notes)
	if result.Error != nil {
		fmt.Println("failed to get notes: ", result.Error)
		return c.Status(400).SendString("failed to get notes")
	}

	return c.Status(fiber.StatusOK).JSON(notes)
}

func CreateNote(c *fiber.Ctx) error {
	note := models.Note{}
	// Parse body into struct
	if err := c.BodyParser(&note); err != nil {
		fmt.Println("failed to parse json body: ", err)
		return c.Status(400).SendString(err.Error())
	}

	db := database.DB.Db
	result := db.Create(&note)
	if result.Error != nil {
		fmt.Println("failed to create note: ", result.Error)
		return c.Status(400).SendString("failed to create note")
	}

	return c.Status(fiber.StatusOK).JSON("note created successfully")
}

func DeleteNote(c *fiber.Ctx) error {
	id := c.Params("id")

	note := models.Note{}
	db := database.DB.Db
	result := db.Where("id", id).Delete(&note)
	if result.Error != nil {
		fmt.Println("failed to get notes: ", result.Error)
		return c.Status(400).SendString("failed to get notes")
	}

	return c.Status(fiber.StatusOK).JSON("note deleted")
}

func UpdateNote(c *fiber.Ctx) error {
	note := models.Note{}
	// Parse body into struct
	if err := c.BodyParser(&note); err != nil {
		fmt.Println("failed to parse json body: ", err)
		return c.Status(400).SendString(err.Error())
	}

	db := database.DB.Db
	result := db.Updates(&note)
	if result.Error != nil {
		fmt.Println("failed to create note: ", result.Error)
		return c.Status(400).SendString("failed to create note")
	}

	return c.Status(fiber.StatusOK).JSON("note created successfully")
}
