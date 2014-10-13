package kamatis

import (
	"fmt"
	"time"
)

//Things unplanned activities with a deadline, and activities that cannot
//be done on unplanned.
type InventorySheet struct {
	Activities []Activity `json:"activities"`
}

/**
 * A recordsheet, is just a view of ToDoLists and InventorySheet.
 * type RecordsSheet struct {
 *   CreatedAt   time.Time
 * }
 */
