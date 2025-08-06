package people

import "log"

func (r *Repository) getEmails(personID int) ([]string, error) {
	rows, err := r.db.Query(`SELECT email FROM emails WHERE person_id = $1`, personID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

	var emails []string
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		emails = append(emails, email)
	}
	return emails, nil
}
