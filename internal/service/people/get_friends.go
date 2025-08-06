package people

import "people/internal/model"

func (s *service) GetFriends(personID int) ([]model.Person, error) {
	return s.friendRepo.GetFriends(personID)
}
