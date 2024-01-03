package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// embedding
type Admin struct {
	email    string
	password string
	User     // embedding
}

// struct using method -- accessing struct thr methods
// receiver argument
func (u *User) OutputUserDetails() {
	//(*u).firstName -- this is correct version.
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// mutation methods
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "User",
			birthdate: "01/01/2000",
			createdAt: time.Now(),
		},
	}
}

// utility functions or constructor functions  ... validation logic in resuseable ways
func New(firstName string, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("fistname, lastname and birthdate are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthDate,
		createdAt: time.Now(),
	}, nil
}
