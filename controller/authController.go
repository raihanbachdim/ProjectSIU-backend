package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/raihanbachdim/ProjectSIU/db"
	"github.com/raihanbachdim/ProjectSIU/model"
	"github.com/raihanbachdim/ProjectSIU/util"
)

func emailValidation(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9. %+\-]+@[a-z0-9. %+\-]+\.[a-z0-9. %+\-]`)
	return Re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData model.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Tidak bisa mengparse body")
	}

	if len(data["password"].(string)) <= 7 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password minimal terdiri dari 8 karakter",
		})
	}

	if !emailValidation(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Mohon masukan alamat email yang valdi",
		})
	}

	db.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email yang anda masukan sudah terdaftar, mohon masukan email lain",
		})
	}
	user := model.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}

	user.SetPassword(data["password"].(string))
	err := db.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Akun telah berhasil dibuat",
	})

}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Tidak bisa mengparse body")
	}
	var user model.User
	db.DB.Where("email=?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Email yang anda masukan tidak terdaftar, silahkan daftarkan akun anda",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password yang anda masukan tidak sesuai",
		})
	}
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Selamat anda berhasil masuk",
		"user":    user,
	})
}

type Claims struct {
	jwt.StandardClaims
}
