package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/golang-fiber/database"
	"github.com/lxxonx/golang-fiber/models"
	"golang.org/x/crypto/bcrypt"
)


func Register(c *fiber.Ctx) error {
	var data map[string]string // string for key, string for value
	
	if err := c.BodyParser(&data); err != nil{
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte((data["password"])), 14)
	user := models.User{
		Name: data["name"],
		Email: data["email"],
		Password: password,
	}

	database.DB.Create(&user)
	
    return c.JSON(user)
}