package people

func (s *service) AddFriend(personID, friendID int) error {
	return s.friendRepo.AddFriend(personID, friendID)
}
