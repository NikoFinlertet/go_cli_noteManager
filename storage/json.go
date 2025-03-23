package storage

import (
	"cli_noteManager/model"
	"encoding/json"
	"os"
)

const notesFile = "notes.json"

type JSONStorage struct{}

func (s *JSONStorage) LoadNotes() ([]model.Note, error) {
	file, err := os.ReadFile(notesFile)
	if err != nil {
		return []model.Note{}, nil
	}

	var notes []model.Note
	json.Unmarshal(file, &notes)

	return notes, nil
}

func (s *JSONStorage) SaveNotes(notes []model.Note) error {
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(notesFile, data, 0644)
}
