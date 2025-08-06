package v1

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"people/internal/model"
	peopleV1 "people/shared/pkg/openapi/people/v1"
)

func (h *peopleHandler) UpdatePerson(ctx context.Context, req *peopleV1.UpdatePerson, params peopleV1.UpdatePersonParams) (peopleV1.UpdatePersonRes, error) {
	if params.ID <= 0 {
		return &peopleV1.BadRequest{
			Code:    400,
			Message: "Invalid person ID",
		}, nil
	}

	if req.Name.IsSet() {
		if name, ok := req.Name.Get(); ok && strings.TrimSpace(name) == "" {
			return &peopleV1.BadRequest{
				Code:    400,
				Message: "Name cannot be empty",
			}, nil
		}
	}

	if req.Surname.IsSet() {
		if surname, ok := req.Surname.Get(); ok && strings.TrimSpace(surname) == "" {
			return &peopleV1.BadRequest{
				Code:    400,
				Message: "Surname cannot be empty",
			}, nil
		}
	}

	if req.Age.IsSet() {
		if age, ok := req.Age.Get(); ok && (age < 0 || age > 190) {
			return &peopleV1.BadRequest{
				Code:    400,
				Message: "Age must be between 0 and 190",
			}, nil
		}
	}

	existingPerson, err := h.Service.SearchByID(params.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &peopleV1.NotFound{
				Code:    404,
				Message: "Person not found",
			}, nil
		}
		return &peopleV1.InternalServer{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	updatedPerson := updatePersonToModel(req, existingPerson)

	if err := h.Service.Update(updatedPerson); err != nil {
		return &peopleV1.InternalServer{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &peopleV1.UpdatePersonOK{}, nil
}

func updatePersonToModel(updatePerson *peopleV1.UpdatePerson, existingPerson model.Person) model.Person {
	result := existingPerson

	if updatePerson.Name.IsSet() {
		if name, ok := updatePerson.Name.Get(); ok {
			result.Name = name
		}
	}

	if updatePerson.Surname.IsSet() {
		if surname, ok := updatePerson.Surname.Get(); ok {
			result.Surname = surname
		}
	}

	if updatePerson.Patronymic.IsSet() {
		if patronymic, ok := updatePerson.Patronymic.Get(); ok {
			result.Patronymic = patronymic
		}
	}

	if updatePerson.Gender.IsSet() {
		if gender, ok := updatePerson.Gender.Get(); ok {
			result.Gender = string(gender)
		}
	}

	if updatePerson.Nationality.IsSet() {
		if nationality, ok := updatePerson.Nationality.Get(); ok {
			result.Nationality = nationality
		}
	}

	if updatePerson.Age.IsSet() {
		if age, ok := updatePerson.Age.Get(); ok {
			result.Age = age
		}
	}

	result.Emails = updatePerson.Emails

	return result
}
