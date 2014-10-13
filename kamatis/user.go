package kamatis

import (
    "time"
    "fmt"
)

type User struct {
    FirstName       string
    LastName        string
    CreatedAt       time.Time
    ToDoLists       []ToDoList
    InventorySheet  InventorySheet
}
