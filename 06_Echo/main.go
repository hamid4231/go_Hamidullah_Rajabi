package main

import (
	"net/http"
	"slices"

	"github.com/labstack/echo/v4"
)

// Food struct for a food item
type Food struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

// Create a slice to store food data
var foods []Food = []Food{}

func main() {
	e := echo.New()

	// Make routes
	e.GET("/api/v1/foods", getAllFoods)
	e.GET("/api/v1/foods/:id", getFoodByID)
	e.POST("/api/v1/foods", addFood)
	e.PUT("/api/v1/foods/:id", updateFood)
	e.DELETE("/api/v1/foods/:id", deleteFood)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// 1. Get all food data
func getAllFoods(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"data": foods,
	})
}

// 2. Get food data by ID
func getFoodByID(c echo.Context) error {
	id := c.Param("id")
	var food Food
	for _, f := range foods {
		if f.ID == id {
			food = f
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": food})
}

// 3. Add new food data
func addFood(c echo.Context) error {
	var food Food
	if err := c.Bind(&food); err != nil {
		return err
	}
	foods = append(foods, food)
	return c.JSON(http.StatusCreated, echo.Map{
		"data": food,
	})
}

// 4. Update food data
func updateFood(c echo.Context) error {
	var food Food
	c.Bind(&food)
	foodid := c.Param("id")
	for i := 0; i < len(foods); i++ {
		if foods[i].ID == foodid {
			foods[i].Name = food.Name
			foods[i].Description = food.Description
			foods[i].Price = food.Price
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": food,
	})
}

// 5. Delete food data
func deleteFood(c echo.Context) error {
	id := c.Param("id")
	foods = slices.DeleteFunc(foods, func(f Food) bool {
		return f.ID == id
	})
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
