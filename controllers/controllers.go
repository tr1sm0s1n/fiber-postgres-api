package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tr1sm0s1n/fiber-postgres-api/models"
	"gorm.io/gorm"
)

type Controllers struct {
	db *gorm.DB
}

func NewControllers(db *gorm.DB) *Controllers {
	return &Controllers{db: db}
}

func (ct *Controllers) CreateOne(c *fiber.Ctx) error {
	var newCertificate models.Certificate
	if err := c.BodyParser(&newCertificate); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	result := ct.db.Create(&newCertificate)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).SendString(result.Error.Error())
	}

	return c.Status(http.StatusCreated).JSON(newCertificate)
}

func (ct *Controllers) ReadAll(c *fiber.Ctx) error {
	var certificates []models.Certificate
	result := ct.db.Find(&certificates)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(certificates)
}

func (ct *Controllers) ReadOne(c *fiber.Ctx) error {
	var oldCertificate models.Certificate
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	result := ct.db.First(&oldCertificate, "id = ?", id)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	}

	return c.Status(http.StatusOK).JSON(oldCertificate)
}

func (ct *Controllers) UpdateOne(c *fiber.Ctx) error {
	var update models.Certificate
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	result := ct.db.First(&update, "id = ?", id)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	}

	if err := c.BodyParser(&update); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	result = ct.db.Save(&update)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(update)
}

func (ct *Controllers) DeleteOne(c *fiber.Ctx) error {
	var trash models.Certificate
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	result := ct.db.First(&trash, "id = ?", id)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	}

	result = ct.db.Delete(&models.Certificate{}, id)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(trash)
}
