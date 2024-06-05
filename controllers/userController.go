package controllers

import (
	"Cargo_Dash/database"
	"Cargo_Dash/middleeware"
	"Cargo_Dash/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllUsers(c *fiber.Ctx) error {
	if err := middleeware.IsAuthorized(c, "users"); err != nil {
		return err
	}

	var users []models.User

	database.DB.Preload("Role").Find(&users)

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	if err := middleeware.IsAuthorized(c, "users"); err != nil {
		return err
	}
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	if err := middleeware.IsAuthorized(c, "users"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	if err := middleeware.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	if err := middleeware.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)

	return nil
}
