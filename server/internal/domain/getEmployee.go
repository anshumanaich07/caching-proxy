package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	Name       string `json:"name" bson:"name"`
	Age        int    `json:"age" bson:"age"`
	Department string `json:"department" bson:"department"`
	Position   string `json:"position" bson:"position"`
	Salary     int    `json:"salary" bson:"salary"`
}

type EmployeeUseCase interface {
	GetEmployeeByID(ctx context.Context, id string) (*Employee, error)
}

type EmployeeRepository interface {
	GetEmployeeByID(ctx context.Context, id primitive.ObjectID) (*Employee, error)
}
