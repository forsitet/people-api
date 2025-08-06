package people

func (r *Repository) AddEmail(personID int, email string) error {
	_, err := r.db.Exec(`INSERT INTO emails (person_id, email) VALUES ($1, $2) ON CONFLICT (email) DO NOTHING`, personID, email)
	return err
}
