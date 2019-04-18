package main

import (
	_ "github.com/lib/pq"
	"net/http"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id    int
	Name  string
	Email string

	password string
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register", r.Context().Value("saas"))
	Respond(w, User{Name: "Hamed"}, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

	})

	tokenString, err := token.SignedString([]byte("dsdsdsds"))
	if err != nil {
		panic(err)
	}
	Respond(w, tokenString, nil)
}
