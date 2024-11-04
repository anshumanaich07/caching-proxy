package routes

import (
	"log"
	"net/http"
	"server/internal/api/controllers"
	"server/internal/repository"
	"server/internal/usecase"

	"go.mongodb.org/mongo-driver/mongo"
)

func InitRouter(db *mongo.Database) *http.ServeMux {
	router := http.NewServeMux()

	empRepo, err := repository.NewEmployeeRepo(db)
	if err != nil {
		log.Fatal(err)
	}
	empUC, err := usecase.NewEmployeeUsecase(empRepo)
	if err != nil {
		log.Fatal(err)
	}

	empController := controllers.NewEmployeeController(empUC)
	router.HandleFunc("GET /employee/{id}", empController.GetEmployeeByID)

	return router
}
