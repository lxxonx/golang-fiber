package controllers

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/golang-fiber/database"
	"github.com/lxxonx/golang-fiber/models"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

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
	
	// if user already exists 
	// it returns id = 0
	database.DB.Create(&user)
	
	if user.ID == 0{return c.JSON(fiber.Map{
		"message": "user already exists",
	})}

    return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string // string for key, string for value
	
	if err := c.BodyParser(&data); err != nil{
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		// user doesn't exist
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			// fiber.map is like dictionary or object type of js
			"message": "user not found",
		})
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON((fiber.Map{
			"message": "incorrect password",
		}))
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour *24).Unix(), // 1day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON((fiber.Map{
			"message": "could not login",
		}))
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour *24), //1 day
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UserPost(c *fiber.Ctx) error {
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

	var user models.User

	database.DB.Where("ID = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	// remove cookie => set expire time past
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour), //1 hr ago
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "logout success",
	})
}
func Delete(c *fiber.Ctx) error {
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

	var user models.User

	database.DB.Where("ID = ?", claims.Issuer).First(&user)

	return nil
}
func UserGet(c *fiber.Ctx) error {
	
	var user models.User
	var posts []models.Post
	database.DB.Where("id = ?", c.Params("id")).First(&user)
	database.DB.Where("user_id = ?", c.Params("id")).Find(&posts)

	user.Posts = posts

	return c.JSON(user)
}
func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Set("gorm:auto_preload", true).Find(&users)
	return c.JSON(users)
}