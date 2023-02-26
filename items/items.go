package items

import (
	"assignment1GO/database"
	"assignment1GO/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func CreateItem(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	item := models.Item{
		Name:   data["name"],
		Price:  data["price"],
		Rating: data["rating"],
	}

	database.DB.Create(&item)

	return c.JSON(item)
}

func Search(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	var item models.Item

	database.DB.Where("name=?", data["name"]).First(&item)

	if item.Name == "" {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "item not found",
		})
	}

	return c.JSON(item)
}

func FilterRating(c *fiber.Ctx) error {
	//var item models.Item
	//
	//database.DB.Order("rating").Find(&item)
	//return c.JSON(item)

	var items []models.Item
	database.DB.Order("rating").Find(&items)

	a, _ := json.Marshal(items)
	n := len(a) //Find the length of the byte array
	s := string(a[:n])
	return c.JSON(s)
}

func FilterPrice(c *fiber.Ctx) error {
	var items []models.Item
	database.DB.Order("price").Find(&items)

	a, _ := json.Marshal(items)
	n := len(a) //Find the length of the byte array
	s := string(a[:n])
	return c.JSON(s)
}

func GiveRating(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	// UPDATE users SET name='hello' WHERE role = 'admin';
	database.DB.Model(models.Item{}).Where("name = ?", data["name"]).Updates(models.Item{Rating: data["rating"]})
	return c.JSON("success")
}
