package people

import (
	"people/internal/model"
)

func (r *Repository) Create(person model.Person) (model.CreatePerson, error) {
	var id int
	err := r.db.QueryRow(
		`INSERT INTO people (name, surname, patronymic, gender, nationality, age) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		person.Name, person.Surname, person.Patronymic, person.Gender, person.Nationality, person.Age,
	).Scan(&id)
	if err != nil {
		return model.CreatePerson{}, err
	}
	for _, email := range person.Emails {
		if err := r.AddEmail(id, email); err != nil {
			return model.CreatePerson{}, err
		}
	}
	return model.CreatePerson{ID: id}, nil
}
