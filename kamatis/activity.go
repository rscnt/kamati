//NOTE: using martini-contrib/binding for tags.
package kamatis

import (
	"fmt"
	"time"
)

type ActivityType struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Activity struct {
	StartedAt             time.Time    `json:"started_at"`
	EndedAt               time.Time    `json:"ended_at"`
	DeadLine              time.Time    `json:"deadline"`
	Type                  ActivityType `json:"activity_type"`
	Description           string       `json:"description"`
	StimatedKamatis       int          `json:"stimated_kamatis"`
	Done                  bool         `json:"done"`
	ActualKamatis         int          `json:"actual_kamatis"`
	InternalInterruptions int          `json:"internal_interruptions"`
	ExternalInterruptions int          `json:"external_interruptions"`
	Notes                 string       `json:"notes"`
}
