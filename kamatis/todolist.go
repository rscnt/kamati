package kamatis

import (
    "time"
    "fmt"
)

//Things ToDo "today" sheet.
type ToDoList struct {
    //TODO: Use geolocation [?].
    Place                   string  
    Date                    time.Time
    ToDoList                []Activity 
    ToDoUrgentUnplanned     []Activity
    AvailableKamatis        int
}
