package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/EdgarAllanPoo/test-go-api/src/infrastructure/db"

	"github.com/EdgarAllanPoo/test-go-api/src/usecases"

	"github.com/EdgarAllanPoo/test-go-api/src/interface/controllers"

	"github.com/EdgarAllanPoo/test-go-api/src/infrastructure/router"

	"github.com/EdgarAllanPoo/test-go-api/src/interface/repository"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
	dbHandler  db.DBHandler
)

func getProductController() controllers.ProductController {
	productRepo := repository.NewProductRepo(dbHandler)
	productInteractor := usecases.NewProductInteractor(productRepo)
	productController := controllers.NewProductController(productInteractor)
	return *productController
}

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "App is up and running..")
	})

	dbURL := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")

	dbHandler, err = db.NewDBHandler(dbURL, dbName)
	if err != nil {
		log.Println("Unable to connect to the DataBase")
		return
	}

	port := os.Getenv("PORT")

	productController := getProductController()
	httpRouter.POST("/product", productController.Add)
	httpRouter.GET("/products", productController.FindAll)
	httpRouter.GET("/product/{id}", productController.FindById)
	httpRouter.DELETE("/product/{id}", productController.Delete)
	httpRouter.PUT("/product/{id}", productController.Update)
	httpRouter.SERVE(port)
}
