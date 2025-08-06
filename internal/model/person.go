package model

type Person struct {
	ID          int      `json:"id" db:"id"`
	Name        string   `json:"name" db:"name"`
	Surname     string   `json:"surname" db:"surname"`
	Patronymic  string   `json:"patronymic" db:"patronymic"`
	Gender      string   `json:"gender" db:"gender"`
	Nationality string   `json:"nationality" db:"nationality"`
	Age         int      `json:"age" db:"age"`
	Emails      []string `json:"emails" db:"-"`
}

type Email struct {
	Email string `json:"email" db:"email"`
}

type CreatePerson struct {
	ID int `json:"id" db:"id"`
}
