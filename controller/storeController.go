package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/raihanbachdim/ProjectSIU/db"
	"github.com/raihanbachdim/ProjectSIU/model"
	"github.com/raihanbachdim/ProjectSIU/util"
	"gorm.io/gorm"
)

func CreateStore(c *fiber.Ctx) error {
	var stores model.Stores
	if err := c.BodyParser(&stores); err != nil {
		fmt.Println("Tidak bisa mengparse body")
	}
	if err := db.DB.Create(&stores).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "payload tidak valdi",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Okay!",
	})

}

func AllStores(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getstore []model.Stores
	db.DB.Preload("User").Offset(offset).Limit(limit).Find(&getstore)
	db.DB.Model(&model.Stores{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": getstore,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func DetailStore(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var storess model.Stores
	db.DB.Where("id=?", id).Preload("User").First(&storess)
	return c.JSON(fiber.Map{
		"data": storess,
	})
}
func UpdateStore(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	store := model.Stores{
		Id: uint(id),
	}
	if err := c.BodyParser(&store); err != nil {
		fmt.Println("Tidak bisa mengparse body")
	}
	db.DB.Model(&store).Updates(store)
	return c.JSON(fiber.Map{
		"message": "Data berhasil di update",
	})
}

func UniqueStore(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)
	var store []model.Stores
	db.DB.Model(&store).Where("user_id=?", id).Preload("User").Find(&store)

	return c.JSON(store)
}
func DeleteStore(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	store := model.Stores{
		Id: uint(id),
	}
	deletequery := db.DB.Delete(&store)
	if errors.Is(deletequery.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Data tidak ditemukan!",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data berhasil dihapus",
	})

}
