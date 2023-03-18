// import "go.mongodb.org/mongo-driver/mongo"

// serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
// clientOptions := options.Client().
//     ApplyURI("mongodb+srv://rerheza:<password>@cluster0.984mmpj.mongodb.net/?retryWrites=true&w=majority").
//     SetServerAPIOptions(serverAPIOptions)
// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// defer cancel()
// client, err := mongo.Connect(ctx, clientOptions)
// if err != nil {
//     log.Fatal(err)
// }

package main

import (
	"fmt"
	"log"
	"net/http"

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

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "App is up and running..")
	})

	var err error
	dbHandler, err = db.NewDBHandler("mongodb+srv://rerheza:astaga@cluster0.984mmpj.mongodb.net/?retryWrites=true&w=majority", "product")
	if err != nil {
		log.Println("Unable to connect to the DataBase")
		return
	}

	productController := getProductController()
	httpRouter.POST("/product/add", productController.Add)
	httpRouter.GET("/product", productController.FindAll)
	httpRouter.GET("/product/{id}", productController.FindById)
	httpRouter.SERVE(":8000")
}
