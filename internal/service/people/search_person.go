package people

import "people/internal/model"

func (s *service) SearchByID(id int) (model.Person, error) {
	return s.peopleRepo.SearchByID(id)
}

func (s *service) SearchBySurname(surname string) ([]model.Person, error) {
	return s.peopleRepo.SearchBySurname(surname)
}
