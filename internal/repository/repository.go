package repository

import "people/internal/model"

type PeopleRepository interface {
	List() ([]model.Person, error)
	Create(person model.Person) (model.CreatePerson, error)
	SearchByID(id int) (model.Person, error)
	SearchBySurname(surname string) ([]model.Person, error)
	Update(person model.Person) error
}

type EmailRepository interface {
	AddEmail(personID int, email string) error
}

type FriendRepository interface {
	AddFriend(personID, friendID int) error
	GetFriends(personID int) ([]model.Person, error)
}
