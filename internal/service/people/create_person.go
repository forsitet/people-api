package people

import (
	"context"

	"people/internal/model"
)

func (s *service) Create(person model.Person) (model.CreatePerson, error) {
	// Если возраст, пол или национальность не указаны, получаем их из внешних API
	if person.Age == 0 || person.Gender == "" || person.Nationality == "" {
		ctx := context.Background()
		externalInfo, err := s.externalClient.GetPersonInfo(ctx, person.Name)
		if err == nil {
			if person.Age == 0 && externalInfo.Age > 0 {
				person.Age = externalInfo.Age
			}
			if person.Gender == "" && externalInfo.Gender != "" {
				person.Gender = externalInfo.Gender
			}
			if person.Nationality == "" && externalInfo.Nationality != "" {
				person.Nationality = externalInfo.Nationality
			}
		}
	}

	return s.peopleRepo.Create(person)
}
