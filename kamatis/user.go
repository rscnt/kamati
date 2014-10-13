package kamatis

import (
	"fmt"
	"time"
)

type User struct {
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	CreatedAt      time.Time      `json:"created_at"`
	ToDoLists      []ToDoList     `json:"todo_lists"`
	InventorySheet InventorySheet `json:"inventory_sheet"`
}
