package kamatis

import (
	"labix.org/v2/mgo"
	"time"
)

type User struct {
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	CreatedAt      time.Time      `json:"created_at"`
	ToDoLists      []ToDoList     `json:"todo_lists"`
	InventorySheet InventorySheet `json:"inventory_sheet"`
}

/**
 * Fetch all the items of the users collection on the current mongodb
 * instance.
 */
func fetchAllUsers(db *mgo.Database) []User {
	users := []User{}
	err := db.C("users").Find(nil).All(&users)
	if err != nil {
		panic(err)
	}
	return users
}

/* TODO: write validation  */
func (user *User) valid() bool {
	return true
}
