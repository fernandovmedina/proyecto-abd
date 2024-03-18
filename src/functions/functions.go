package functions

import (
	"github.com/fernandovmedina/abd/src/database"
)

type User struct {
	Id       int
	Name     string
	Lastname string
	Email    string
	Password string
}

func Login(email, pass string) (bool, error) {
	rows, err := database.DB.Query("SELECT ID,NAME,LASTNAME,EMAIL,PASS FROM AUTH.USERS WHERE EMAIL=? AND PASS=?", &email, &pass)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	var u = new(User)

	for rows.Next() {
		if err = rows.Scan(&u.Id, &u.Name, &u.Lastname, &u.Email, &u.Password); err != nil {
			return false, err
		}
	}

	if u.Id == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func Register(name, lastname, email, password string) (bool, error) {
	if name == "" || lastname == "" || email == "" || password == "" {
		return false, nil
	}

	_, err := database.DB.Exec("INSERT INTO AUTH.USERS(NAME,LASTNAME,EMAIL,PASS)VALUES(?,?,?,?)", &name, &lastname, &email, &password)

	if err != nil {
		return false, err
	}

	return true, nil
}
