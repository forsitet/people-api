package v1

import (
	"context"

	"people/internal/converter"
	peopleV1 "people/shared/pkg/openapi/people/v1"
)

func (h *peopleHandler) CreatePerson(ctx context.Context, req *peopleV1.PostCreatePerson) (peopleV1.CreatePersonRes, error) {
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
