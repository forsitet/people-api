package people

import "people/internal/model"

func (s *service) Update(person model.Person) error {
	return s.peopleRepo.Update(person)
}
