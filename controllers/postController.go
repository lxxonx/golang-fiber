package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/golang-fiber/database"
	"github.com/lxxonx/golang-fiber/models"
)

func CreatePost(c *fiber.Ctx) error {
	var data map[string]string // string for key, string for value
	
	if err := c.BodyParser(&data); err != nil{
		return err
	}
	number, _ := strconv.ParseUint(data["userId"],10,64)
	post := models.Post{
		Title: data["title"],
		Text: data["text"],
		UserId: number,
	}
	
	// if user already exists 
	// it returns id = 0
	database.DB.Create(&post)
	
	if post.ID == 0{return c.JSON(fiber.Map{
		"message": "post already exists",
	})}

    return c.JSON(post)
}
func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post
	database.DB.Find(&posts)

	return c.JSON(posts)
}