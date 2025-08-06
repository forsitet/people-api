package v1

import (
	"people/internal/service"
	peopleV1 "people/shared/pkg/openapi/people/v1"
)

type peopleHandler struct {
	Service service.PeopleService
}

func NewPeopleHandler(service service.PeopleService) peopleV1.Handler {
	return &peopleHandler{Service: service}
}
