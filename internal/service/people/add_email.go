package people

func (s *service) AddEmail(personID int, email string) error {
	return s.emailRepo.AddEmail(personID, email)
}
