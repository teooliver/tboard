package data

import (
	"time"

	"github.com/google/uuid"
)

type DenormalizedBoard struct {
	id        uuid.UUID // uuid
	name      string
	columns   []DenormalizedColumn // column ids
	createdAt time.Time
	createdBy uuid.UUID // user id
}

type DenormalizedColumn struct {
	id        uuid.UUID
	name      string
	tasks     []Task
	createdAt time.Time
	createdBy uuid.UUID
}

func denormalizedBoard(board *Board) *DenormalizedBoard {
	// var denormalizedBoard DenormalizedBoard
	var columnsSlice = make([]Column, 0)

	for _, id := range board.columnRefs {
		c, found := getColumn(id)
		if !found {
			// Column not found
		}
		columnsSlice = append(columnsSlice, c)
	}
	for _, c := range columnsSlice {

		// add tasks
	}

	// denormalize Tasks

	return &DenormalizedBoard{
		id:   board.id,
		name: board.name,
		// columns: ,
	}
}

// If Column doesnt exist return false
func getColumn(columnRef uuid.UUID) (Column, bool) {
	var columns []Column = make([]Column, 0)

	for _, c := range columns {
		if c.id == columnRef {
			return c, true
		}
	}
	return Column{}, false
}
