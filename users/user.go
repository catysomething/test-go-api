package users

import (
	"errors"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

// User type fields should be in camelCase, but if we change it to start lowercase,
// the field will be private and not exported. Instead, need to use tags (in backticks).
// You can use multiple tags for each field.

// User hold data for a single user
type User struct {
	ID   bson.ObjectId `json:"id" storm:"id"`
	Name string        `json:"name"`
	Role string        `json:"role"`
}

// specify path to access data
const (
	dbPath = "users.db"
)

//errors
var (
	ErrRecordInvalid = errors.New("Record is invalid")
)

// All retrieves all users from the database
func All() ([]User, error) {
	// will create db if it doesn't exist and return any error(s)
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	// don't forget to close connection!
	defer db.Close()

	// create list of users
	users := []User{}
	// use storm's All() function and point it to the newly created list
	// Storm requires use of point, even though slices are always passed by reference in Go anyway
	err = db.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// One returns a single user record from the database
// First (parenthetical) passes value(s) in, second returns
func One(id bson.ObjectId) (*User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	// don't forget to close connection!
	defer db.Close()

	u := new(User)
	// Storm's One method takes 3 arguments:
	//     1. Name of field containing unique ID
	//     2. Value of field we want to retrieve
	//     3. Pointer that will receive data (new keyword returns pointer by default so '&' is not needed.)
	err = db.One("ID", id, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Delete removes a single user record from the database
func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return err
	}

	// instead of returning the object, return the result of DeleteStruct(u)
	return db.DeleteStruct(u)
}

// Save updates or creates a given record in the database
// member of user struct
func (u *User) Save() error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(u)
}

func (u *User) validate() error {
	if u.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}
