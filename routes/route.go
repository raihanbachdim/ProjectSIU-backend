package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raihanbachdim/ProjectSIU/controller"
	"github.com/raihanbachdim/ProjectSIU/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Use(middleware.IsAuthenticate)
	app.Post("/api/store", controller.CreateStore)
	app.Get("/api/allstore", controller.AllStores)
	app.Get("/api/allstore/:id", controller.DetailStore)
	app.Put("/api/updatestore/:id", controller.UpdateStore)
	app.Get("/api/uniquestore", controller.UniqueStore)
	app.Delete("/api/deletestore/:id", controller.DeleteStore)
	app.Post("/api/upload-image", controller.UploadImg)
	app.Static("/api/uploads", "./uploads")
}
