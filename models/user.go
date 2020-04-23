package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

// slice of pointers of User objects
var (
	users  []*User
	nextID = 1 // implicit type of int
)

// GetUsers is getting users
func GetUsers() []*User {
	return users
}

// AddUser is invoked to ad a user
func AddUser(u User) (User, error) { // consume type user, return either a user or an error
	if u.ID != 0 {
		return User{}, errors.New("New user must not include ID or it must be set to zero.")
	}
	u.ID = nextID             // apply next ID
	nextID++                  // increase next ID
	users = append(users, &u) // append addressOf the u that came in, since this is what the users slice stores
	return u, nil             // return the newly created user object and nil, because no error
}
func getUserByID(id int) (User, error) { // maybe asking for an User by a non-assigned id in collection should error properly
	for _, u := range users { // looping through the users mapping wildcarding the iterator
		if u.ID == id { // if the provided id in the call matches the currently tested user's ID
			return *u, nil // return den value this pointer points at (dereferencing) and nil indicating no error happened
		}
	}
	return User{}, fmt.Errorf("No valid user matching the provided ID '%v' found", id) // when no matching ID error like this

}

func updateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u // set addressOf of that candidate's matching user to new value
			return u, nil // and return user without an error
		}
	}
	return User{}, fmt.Errorf("No valid user matching the provided ID '%v' found", u.ID) //incoming user's id
}

func removeUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...) // append everything fron the slice before the user we found and after the user we found
			return nil
		}
	}
	return fmt.Errorf("No valid user matching the provided ID '%v' found", id) //incoming id param
}
