package main

import (
	"log"
	"net/http"

	"github.com/bharatkalra244/car-app/models"
	"github.com/bharatkalra244/car-app/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Car struct {
	CAR_NAME    string `json:"carName"`
	CAR_HP      int32  `json:"carHp"`
	CAR_COMPANY string `json:"carCompany"`
	CAR_ENGINE  string `json:"carEngine"`
	//	CarID       int64  `json:"carID"`
}

type Repository struct {
	DB *gorm.DB
}

// Method created for Repository due to the presence of (r *Repository)
func (r *Repository) CreateCar(context *fiber.Ctx) error {
	car := Car{}
	//Work of body parser is to parse through the json body that we are getting in the request
	err := context.BodyParser(&car)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&car).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create car"})
		return err

	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Car has been added"})
	return nil
}

func (r *Repository) DeleteCar(context *fiber.Ctx) error {
}

func (r *Repository) GetCars(context *fiber.Ctx) error {
	carModels := &[]models.Cars{}

	err := r.DB.Find(carModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get cars"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "cars fetched successfully",
			"data":    carModels,
		})
	return nil
}

func (r *Repository) GetCarsByID(context *fiber.Ctx) error {
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_cars", r.CreateCar)
	api.Delete("/delete_car/:id", r.DeleteCar)
	api.Get("/get_cars/:id", r.GetCarsByID)
	api.Get("/cars", r.GetCars)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal(err, "Could Not Load The Database !")
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
