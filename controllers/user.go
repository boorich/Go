package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

// A controller to receive and process a network request
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) { //binding the ServeHTTP method to the UserController type. This actually implements the Handler interface only by naming the method exactly right.
	w.Write([]byte("Hello from the User Controller")) //writing back to the HTTP response object
}

// Contructor Function
func newUserController() *userController { // return pointer to userController object.
	return &userController{ // the userController variable is created on the fly
		userIDPattern: regexp.MustCompile(`^/users/(/d+)/?`), //if that Regex works, we have a user
	}
}
