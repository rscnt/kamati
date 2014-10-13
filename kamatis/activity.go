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
    DeadLine                time.Time
    Type                    ActivityType    
    ActivityDescription     string 
    StimatedKamatis         int
    Done                    bool
    ActualKamatis           int 
    InternalInterruptions   int
    ExternalInterruptions   int
    ActivityNotes           string
}
