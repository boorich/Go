package controllers

import (
	"encoding/json"
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
			uc.post(w, r) // pass it on to this function
		default:
			w.WriteHeader(http.StatusNotImplemented) // in case none of the above verbs is used for a request
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path) // compiles a slice of all matching expressions as defined in the user controller like so userIDPattern: regexp.MustCompile(`^/users/(/d+)/?`)
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
	u, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

// add a new user
func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

// update a user
func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of user must match ID in URL"))
		return
	}
	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

// delete a user
func (uc *userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

// helper function to convert incoming resource (e.g.update request) into an object
func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

// Contructor Function
func newUserController() *userController { // return pointer to userController object.
	return &userController{ // the userController variable is created on the fly
		userIDPattern: regexp.MustCompile(`^/users/(/d+)/?`), //if that Regex works, we have a user
	}
}
