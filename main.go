package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/fernandovmedina/abd/src/database"
	"github.com/fernandovmedina/abd/src/functions"
)

type User struct {
	Id       int
	Name     string
	Lastname string
	Email    string
	Password string
}

func main() {
	var err error

	if err = database.OpenDB(); err != nil {
		fmt.Println(err.Error())
	}

	var server = http.NewServeMux()

	server.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("./public/home.html")

		if err != nil {
			fmt.Println(err.Error())
		}

		temp.Execute(w, nil)
	})

	server.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("./public/login.html")

		if err != nil {
			fmt.Println(err.Error())
		}

		found, err := functions.Login(r.FormValue("email"), r.FormValue("password"))

		if err != nil {
			fmt.Println(err.Error())
		}

		if found {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
		}

		if err = temp.Execute(w, nil); err != nil {
			fmt.Println(err.Error())
		}
	})

	server.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("./public/register.html")

		if err != nil {
			fmt.Println(err.Error())
		}

		found, err := functions.Register(r.FormValue("name"), r.FormValue("lastname"), r.FormValue("email"), r.FormValue("password"))

		if err != nil {
			fmt.Println(err.Error())
		}

		if found {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
		}

		if err = temp.Execute(w, nil); err != nil {
			fmt.Println(err.Error())
		}
	})

	fmt.Println("Servidor corriendo en http://127.0.0.1:8080")

	if err := http.ListenAndServe(":8080", server); err != nil {
		fmt.Println(err.Error())
	}
}
