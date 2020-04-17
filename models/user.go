package models

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
	u.ID = nextID             //apply next ID
	nextID++                  // increas next ID
	users = append(users, &u) // append addressOf the u that came in, since this is whhat the users slice stores
	return u, nil             // return the newly created user object and nil, because no error
}
