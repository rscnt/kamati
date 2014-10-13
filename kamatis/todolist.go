package kamatis

import (
	"fmt"
	"time"
)

//Things ToDo "today" sheet.
type ToDoList struct {
	//TODO: Use geolocation [?].
	Place               string     `json:"place"`
	Date                time.Time  `json:"created_at"`
	EndedAt             time.Time  `json:"ended_at"`
	ToDoList            []Activity `json:"todolist"`
	ToDoUrgentUnplanned []Activity `json:"urgent_unplanned"`
	AvailableKamatis    int        `json:"available_kamatis"`
}
