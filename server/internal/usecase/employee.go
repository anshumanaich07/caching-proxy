package usecase

import (
	"context"
	"server/internal/domain"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type employeeUsecase struct {
	EmployeeRepo domain.EmployeeRepository
}

func NewEmployeeUsecase(repo domain.EmployeeRepository) (domain.EmployeeUseCase, error) {
	return employeeUsecase{
		EmployeeRepo: repo,
	}, nil
}

func (uc employeeUsecase) GetEmployeeByID(ctx context.Context, id string) (*domain.Employee, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to convert string to ObjectID")
	}
	return uc.EmployeeRepo.GetEmployeeByID(ctx, objID)

}
