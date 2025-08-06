package people

import (
	"log"

	"people/internal/model"
)

func (r *Repository) List() ([]model.Person, error) {
	rows, err := r.db.Query(`SELECT id, name, surname, patronymic, gender, nationality, age FROM people`)
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
