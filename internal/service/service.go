package service

import (
	"people/internal/model"
)

type PeopleService interface {
	List() ([]model.Person, error)
	Create(person model.Person) (model.CreatePerson, error)
	SearchByID(id int) (model.Person, error)
	SearchBySurname(surname string) ([]model.Person, error)
	Update(person model.Person) error

	AddEmail(personID int, email string) error

	AddFriend(personID, friendID int) error
	GetFriends(personID int) ([]model.Person, error)
}
