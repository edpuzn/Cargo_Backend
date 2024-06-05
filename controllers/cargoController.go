package controllers

import (
	"Cargo_Dash/database"
	"Cargo_Dash/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// Tüm cargo kayıtlarını getirir
func AllCargo(c *fiber.Ctx) error {
	var cargos []models.Cargo

	database.DB.Find(&cargos)

	return c.JSON(cargos)
}

// Yeni bir cargo kaydı oluşturur
func CreateCargo(c *fiber.Ctx) error {
	cargo := new(models.Cargo)

	if err := c.BodyParser(cargo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	database.DB.Create(&cargo)

	return c.JSON(cargo)
}

func GetByCargo(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	cargo := models.Cargo{}

	database.DB.First(&cargo, id)

	return c.JSON(cargo)
}

// Belirli bir cargo kaydını getirir
func GetCargo(c *fiber.Ctx) error {
	id := c.Params("tracking_number")
	cargo := models.Cargo{}
	result := database.DB.Where("tracking_number = ?", id).First(&cargo)
	if result.Error != nil {
		// Veritabanı hatası durumunda hata mesajı döndür
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	if result.RowsAffected == 0 {
		// Kayıt bulunamadı durumunda hata mesajı döndür
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Cargo not found",
		})
	}
	return c.JSON(cargo)
}

// Belirli bir cargo kaydını günceller
func UpdateCargo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cargo := new(models.Cargo)
	if err := c.BodyParser(cargo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result := database.DB.Model(&models.Cargo{}).Where("id = ?", id).Updates(cargo)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.JSON(cargo)
}

// Belirli bir cargo kaydını siler
func DeleteCargo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result := database.DB.Delete(&models.Cargo{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
