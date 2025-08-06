package people

import "people/internal/model"

func (s *service) List() ([]model.Person, error) {
	return s.peopleRepo.List()
}
