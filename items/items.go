package items

import (
	"assignment1GO/database"
	"assignment1GO/models"
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

	return c.JSON(item)
}

func FilterRating(c *fiber.Ctx) error {
	var item models.Item

	database.DB.Order("rating").Find(&item)
	return c.JSON(item)
}

func FilterPrice(c *fiber.Ctx) error {
	var item models.Item

	database.DB.Order("price").Find(&item)
	return c.JSON(models.Item{})
}

func GiveRating(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	// UPDATE users SET name='hello', age=18 WHERE role = 'admin';
	database.DB.Model(models.Item{}).Where("name = ?", data["name"]).Updates(models.Item{Rating: data["rating"]})
	return c.JSON(models.Item{})
}