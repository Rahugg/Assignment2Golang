package routes

import (
	"github.com/gofiber/fiber/v2"

	"assignment1GO/controllers"
	"assignment1GO/items"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("api/login", controllers.Login)
	app.Get("api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)

	app.Post("api/createItem", items.CreateItem)
	app.Get("api/filterPrice", items.FilterPrice)
	app.Get("api/filterRating", items.FilterRating)
	app.Post("api/search", items.Search)
	app.Post("api/giveRating", items.GiveRating)
}
