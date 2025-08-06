package people

import (
	"people/internal/external"
	"people/internal/repository"
	def "people/internal/service"
)

var _ def.PeopleService = (*service)(nil)

type service struct {
	peopleRepo     repository.PeopleRepository
	emailRepo      repository.EmailRepository
	friendRepo     repository.FriendRepository
	externalClient external.Client
}

func NewService(peopleRepo repository.PeopleRepository, emailRepo repository.EmailRepository, friendRepo repository.FriendRepository) *service {
	return &service{
		peopleRepo:     peopleRepo,
		emailRepo:      emailRepo,
		friendRepo:     friendRepo,
		externalClient: external.NewClient(),
	}
}
