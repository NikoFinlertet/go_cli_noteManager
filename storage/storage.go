package storage

import (
	"cli_noteManager/model"
)

type Storage interface {
	LoadNotes() ([]model.Note, error)
	SaveNotes([]model.Note, error)
}
