package people

import (
	"database/sql"

	"people/internal/model"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) scanPeopleRows(rows *sql.Rows) ([]model.Person, error) {
	var people []model.Person
	for rows.Next() {
		var p model.Person
		if err := rows.Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Gender, &p.Nationality, &p.Age); err != nil {
			return nil, err
		}
		emails, err := r.getEmails(p.ID)
		if err != nil {
			return nil, err
		}
		p.Emails = emails
		people = append(people, p)
	}
	return people, nil
}
