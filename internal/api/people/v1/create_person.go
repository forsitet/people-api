package v1

import (
	"context"
	"strings"

	"people/internal/converter"
	peopleV1 "people/shared/pkg/openapi/people/v1"
)

func (h *peopleHandler) CreatePerson(ctx context.Context, req *peopleV1.PostCreatePerson) (peopleV1.CreatePersonRes, error) {
	if strings.TrimSpace(req.Name) == "" {
		return &peopleV1.BadRequest{
			Code:    400,
			Message: "Name is required",
		}, nil
	}

	if strings.TrimSpace(req.Surname) == "" {
		return &peopleV1.BadRequest{
			Code:    400,
			Message: "Surname is required",
		}, nil
	}

	if strings.TrimSpace(req.Patronymic) == "" {
		return &peopleV1.BadRequest{
			Code:    400,
			Message: "Patronymic is required",
		}, nil
	}

	created, err := h.Service.Create(converter.PersonToModel(req))
	if err != nil {
		return &peopleV1.InternalServer{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &peopleV1.CreatePerson{
		ID: created.ID,
	}, nil
}
