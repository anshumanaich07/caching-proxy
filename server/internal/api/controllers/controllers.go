package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"server/internal/api/response"
	"server/internal/domain"
)

type EmployeeController struct {
	EmployeeUC domain.EmployeeUseCase
}

func NewEmployeeController(empUC domain.EmployeeUseCase) EmployeeController {
	return EmployeeController{
		EmployeeUC: empUC,
	}
}

func (ec EmployeeController) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	if r.Context() == nil {
		ctx = context.TODO()
	}

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	res, err := ec.EmployeeUC.GetEmployeeByID(ctx, id)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
	}
	if res == nil {
		w.WriteHeader(http.StatusBadRequest)
		errRes := response.EmployeeFailureResponse{
			Error:   "Record not found",
			Message: "The requested resource could not be located",
		}
		json.NewEncoder(w).Encode(errRes)
		return
	}

	w.WriteHeader(http.StatusOK)
	succRes := response.EmployeeSuccessResponse{
		Data:    res,
		Message: "Employee found successfully",
	}

	json.NewEncoder(w).Encode(succRes)
}
