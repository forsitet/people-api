package people

func (r *Repository) AddFriend(personID, friendID int) error {
	_, err := r.db.Exec(`INSERT INTO friends (person_id, friend_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`, personID, friendID)
	return err
}
