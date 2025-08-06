package v1

import (
	"context"
	"database/sql"
	"errors"

	"people/internal/model"
	peopleV1 "people/shared/pkg/openapi/people/v1"
)

func (h *peopleHandler) AddEmail(ctx context.Context, req *peopleV1.Email, params peopleV1.AddEmailParams) (peopleV1.AddEmailRes, error) {
	// Проверяем существование человека
	_, err := h.Service.SearchByID(params.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &peopleV1.NotFound{
				Code:    404,
				Message: "Person not found",
			}, nil
		}
		return &peopleV1.InternalServer{
			Code:    500,
			Message: "Failed to check person existence",
		}, nil
	}

	if err := h.Service.AddEmail(params.ID, req.Email); err != nil {
		// Проверяем, является ли ошибка связанной с дублированием email
		if errors.Is(err, model.ErrEmailAlreadyExists) {
			return &peopleV1.AlreadyExists{
				Code:    300,
				Message: "Email already exists",
			}, nil
		}
		return &peopleV1.InternalServer{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &peopleV1.AddEmailOK{}, nil
}
