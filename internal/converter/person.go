package converter

import (
	"people/internal/model"
	peopleV1 "people/shared/pkg/openapi/people/v1"
)

func PersonToModel(person *peopleV1.PostCreatePerson) model.Person {
	var gender string
	if person.Gender.IsSet() {
		if g, ok := person.Gender.Get(); ok {
			gender = string(g)
		}
	}
	var nationality string
	if person.Nationality.IsSet() {
		if n, ok := person.Nationality.Get(); ok {
			nationality = n
		}
	}
	var age int
	if person.Age.IsSet() {
		if a, ok := person.Age.Get(); ok {
			age = a
		}
	}
	return model.Person{
		Name:        person.Name,
		Surname:     person.Surname,
		Patronymic:  person.Patronymic,
		Gender:      gender,
		Nationality: nationality,
		Age:         age,
		Emails:      person.Emails,
	}
}

func PersonToApiModel(person model.Person) peopleV1.Person {
	return peopleV1.Person{
		ID:          person.ID,
		Name:        person.Name,
		Surname:     person.Surname,
		Patronymic:  person.Patronymic,
		Gender:      person.Gender,
		Nationality: person.Nationality,
		Age:         person.Age,
		Emails:      person.Emails,
	}
}
