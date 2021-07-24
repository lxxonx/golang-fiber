package controllers

import (
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/golang-fiber/database"
	"github.com/lxxonx/golang-fiber/models"
)
type Input struct {
	Title		string	`json:"title"`
	Text		string	`json:"text"`
	Music 		[]uint8	`json:"music"`
}

func CreatePost(c *fiber.Ctx) error {
	var data map[string]string // string for key, string for value
	
	// parse data
	if err := c.BodyParser(&data); err != nil{
		// if err returns err
		return err
	}

	// get token from Cookie
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error){
		return []byte(SecretKey), nil
	})

	// verify token
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON((fiber.Map{
			"message": "unauthenticated",
		}))
	}
	// get user who logged in
	claims := token.Claims.(*jwt.StandardClaims)
	issuer, _ := strconv.ParseUint(claims.Issuer, 10, 64)
	
	post := models.Post{
		Title: data["title"],
		Text: data["text"],
		Music: data["music"],
		UserId: issuer,
	}
	
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
func UploadMusic(c *fiber.Ctx) error {
	// Get first file from form field "document":
	file, _ := c.FormFile("music")
	// Save file to root directory:
	return c.SaveFile(file, fmt.Sprintf("./musics/%s", file.Filename))
}