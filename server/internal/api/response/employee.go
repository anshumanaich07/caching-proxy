package response

import "server/internal/domain"

type EmployeeSuccessResponse struct {
	Data    *domain.Employee `json:"data"`
	Message string           `json:"message"`
}

type EmployeeFailureResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
