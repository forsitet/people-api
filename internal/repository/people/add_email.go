package people

import (
	"log"

	"people/internal/model"
)

func (r *Repository) AddEmail(personID int, email string) error {
	result, err := r.db.Exec(`INSERT INTO emails (person_id, email) VALUES ($1, $2) ON CONFLICT (email) DO NOTHING`, personID, email)
	if err != nil {
		log.Print(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Print(err)
		return err
	}

	if rowsAffected == 0 {
		return model.ErrEmailAlreadyExists
	}

	return nil
}
