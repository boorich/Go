package controllers

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/empea-careercriminal/Go/models"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

// A controller to receive and process a network request aka Traffic Cop deciding which method to pass it on to
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) { //binding the ServeHTTP method to the UserController type. This actually implements the Handler interface only by naming the method exactly right.
	if r.URL.Path == "/users" { // this request object is the entire collection of users
		switch r.Method { // what's the method on the incoming request?
		case http.MethodGet: // in case of a GET request
			uc.getAll(w, r) // pass it on to this function
		case http.MethodPost: // in case of a POST request
			uc.Post(w, r) // pass it on to this function
		default:
			w.WriteHeader(http.StatusNotImplemented) // in case none of the above verbs is used for a request
		}
	} else {
		matches := uc.userIDPattern.FindAllStringSubmatch(r.URL.Path) // compiles a slice of all matching expressions as defined in the user controller like so userIDPattern: regexp.MustCompile(`^/users/(/d+)/?`)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1]) // make this an int
		if err == nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

// retrieve all users from models layer and return back out as JSON
func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

// return single user by ID
func (uc *userController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUsersByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

// Contructor Function
func newUserController() *userController { // return pointer to userController object.
	return &userController{ // the userController variable is created on the fly
		userIDPattern: regexp.MustCompile(`^/users/(/d+)/?`), //if that Regex works, we have a user
	}
}
