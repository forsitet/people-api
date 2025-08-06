package people

import (
	"log"

	"people/internal/model"
)

func (r *Repository) SearchByID(id int) (model.Person, error) {
	var person model.Person
	err := r.db.QueryRow(
		`SELECT id, name, surname, patronymic, gender, nationality, age FROM people WHERE id = $1`,
		id,
	).Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Gender, &person.Nationality, &person.Age)
	if err != nil {
		return model.Person{}, err
	}

	emails, err := r.getEmails(person.ID)
	if err != nil {
		return model.Person{}, err
	}
	person.Emails = emails

	return person, nil
}

func (r *Repository) SearchBySurname(surname string) ([]model.Person, error) {
	rows, err := r.db.Query(
		`SELECT id, name, surname, patronymic, gender, nationality, age FROM people WHERE surname = $1`,
		surname,
	)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

	return r.scanPeopleRows(rows)
}
