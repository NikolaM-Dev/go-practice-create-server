package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (u *User) ToJson() ([]byte, error) {
	user, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	return user, nil
}
