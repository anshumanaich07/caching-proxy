package repository

import (
	"context"
	"server/internal/domain"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type employeeRepo struct {
	db *mongo.Database
}

func NewEmployeeRepo(db *mongo.Database) (domain.EmployeeRepository, error) {
	return employeeRepo{
		db: db,
	}, nil
}

func (repo employeeRepo) GetCollName() string {
	return "employee"
}

func (repo employeeRepo) GetEmployeeByID(ctx context.Context, id primitive.ObjectID) (*domain.Employee, error) {
	coll := repo.db.Collection(repo.GetCollName())

	var employee *domain.Employee
	err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(&employee)
	if err != nil {
		return nil, errors.Wrap(err, "unable to find employee")
	}

	return employee, nil
}
