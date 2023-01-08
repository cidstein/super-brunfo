package model

import (
	"errors"
	"net/mail"
)

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type UserSession struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	Authenticated bool   `json:"authenticated"`
}

func NewUser(id, email, password, nickname string) User {
	return User{
		ID:       id,
		Email:    email,
		Password: password,
		Nickname: nickname,
	}
}

func (u *User) IsValid() error {
	if u.ID == "" {
		return errors.New("ID is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	if !valid(u.Email) {
		return errors.New("email is invalid")
	}

	if u.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (u *User) ComparePassword(password string) bool {
	return u.Password == password
}
