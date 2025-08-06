package people

import (
	"log"

	"people/internal/model"
)

func (r *Repository) GetFriends(personID int) ([]model.Person, error) {
	rows, err := r.db.Query(`
		SELECT p.id, p.name, p.surname, p.patronymic, p.gender, p.nationality, p.age 
		FROM people p 
		JOIN friends f ON p.id = f.friend_id 
		WHERE f.person_id = $1
	`, personID)
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
