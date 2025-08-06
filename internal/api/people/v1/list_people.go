package v1

import (
	"context"

	"people/internal/converter"
	peopleV1 "people/shared/pkg/openapi/people/v1"
)

func (h *peopleHandler) ListPeople(ctx context.Context) (peopleV1.ListPeopleRes, error) {
	people, err := h.Service.List()
	if err != nil {
		return &peopleV1.InternalServer{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	openAPIPeople := make(peopleV1.ListPeopleOKApplicationJSON, len(people))
	for i, person := range people {
		openAPIPeople[i] = converter.PersonToApiModel(person)
	}

	return &openAPIPeople, nil
}
