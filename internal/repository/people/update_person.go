package people

import "people/internal/model"

func (r *Repository) Update(person model.Person) error {
	_, err := r.db.Exec(
		`UPDATE people SET name=$1, surname=$2, patronymic=$3, gender=$4, nationality=$5, age=$6 WHERE id=$7`,
		person.Name, person.Surname, person.Patronymic, person.Gender, person.Nationality, person.Age, person.ID,
	)
	return err
}
