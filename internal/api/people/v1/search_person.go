package v1

import (
	"context"
	"database/sql"
	"errors"

	"people/internal/converter"
	peopleV1 "people/shared/pkg/openapi/people/v1"
)

func (h *peopleHandler) SearchPerson(ctx context.Context, params peopleV1.SearchPersonParams) (peopleV1.SearchPersonRes, error) {
	idSet := params.ID.IsSet()
	surnameSet := params.Surname.IsSet()

	if idSet && surnameSet {
		return &peopleV1.BadRequest{
			Code:    400,
			Message: "only one parameter (id or surname) should be provided, not both",
		}, nil
	}

	if !idSet && !surnameSet {
		return &peopleV1.BadRequest{
			Code:    400,
			Message: "either id or surname must be provided",
		}, nil
	}

	var openAPIPerson peopleV1.Person
	if id, ok := params.ID.Get(); ok {
		if id <= 0 {
			return &peopleV1.BadRequest{
				Code:    400,
				Message: "Invalid person ID",
			}, nil
		}

		person, err := h.Service.SearchByID(id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return &peopleV1.NotFound{
					Code:    404,
					Message: "person not found",
				}, nil
			}
			return &peopleV1.InternalServer{
				Code:    500,
				Message: err.Error(),
			}, nil
		}

		openAPIPerson = converter.PersonToApiModel(person)
	}

	if surname, ok := params.Surname.Get(); ok {
		people, err := h.Service.SearchBySurname(surname)
		if err != nil {
			return &peopleV1.InternalServer{
				Code:    500,
				Message: err.Error(),
			}, nil
		}
		if len(people) == 0 {
			return &peopleV1.NotFound{
				Code:    404,
				Message: "no people found with this surname",
			}, nil
		}

		person := people[0]
		openAPIPerson = converter.PersonToApiModel(person)
	}

	return &openAPIPerson, nil
}
