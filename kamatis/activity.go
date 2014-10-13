package kamatis

import (
    "time"
    "fmt"
)


type ActivityType struct {
    Name        string
    Description string 
    CreatedAt   time.Time
}

type Activity struct {
    StartedAt               time.Time
    EndedAt                 time.Time
    ToDoList                ToDoList
    Type                    ActivityType    
    ActivityDescription     string 
    ActualKamatis           int 
    InternalInterruptions   int
    ExternalInterruptions   int
    ActivityNotes           string
}

type ToDoList struct {
    Place   string //TODO: Use geolocation. 
    Date    time.Time
}

type User struct {
    FirstName string
    LastName  string
    CreatedAt time.Time
    ToDoLists []ToDoList
}

