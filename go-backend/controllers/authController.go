package controllers

import (
	"go-backend/database"
	"go-backend/models"
	"go-backend/util"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	//read the data
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	//check password and the confirmation equals
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password do not match",
		})
	}

	//write data to user struct
	user := models.User{
		Firstname: data["first_name"],
		Lastname:  data["last_name"],
		Email:     data["email"],
		RoleId:    1,
	}
	// set the password
	user.SetPassword(data["password"])
	//write to database
	result := database.DB.Create(&user)
	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate entry") {
		c.Status(http.StatusConflict)
		return c.JSON(fiber.Map{
			"message": "Email already exists",
		})
	} else if result.Error != nil && !strings.Contains(result.Error.Error(), "Duplicate entry") {
		c.Status(http.StatusConflict)
		return c.JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	//read the data
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	//get user from db where email
	database.DB.Where("email = ?", data["email"]).First(&user)

	//if user is empty
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "not found",
		})
	}

	//compare user in database with sended password
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password is incorrect",
		})
	}

	//create token
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	//make cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	//return
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	//get data from cookie
	id, _ := util.ParseJwt(cookie)

	var user models.User

	database.DB.Where("id = ? ", id).First(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add((-time.Hour)),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)
	userId, _ := strconv.Atoi(id)
	user := models.User{
		Id:        uint(userId),
		Firstname: data["first_name"],
		Lastname:  data["last_name"],
		Email:     data["email"],
	}

	database.DB.Model(&user).Where("id = ?", id).Updates(user)

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//check password and the confirmation equals
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password do not match",
		})
	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)
	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
	}

	user.SetPassword(data["password"])

	database.DB.Model(&user).Where("id = ?", id).Updates(user)

	return c.JSON(user)
}
