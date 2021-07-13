package controllers

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
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
func DeletePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error){
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON((fiber.Map{
			"message": "unauthenticated",
		}))
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var post models.Post
	database.DB.Where("ID = ? ", c.Params("id")).First(&post)

	issuer, _ := strconv.ParseUint(claims.Issuer,10,64)
	if post.UserId == issuer{
	database.DB.Where("ID = ?", c.Params("id")).Delete(&post)
	}
	return nil
}