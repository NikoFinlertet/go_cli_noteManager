package cmd

import (
	"cli_noteManager/model"
	"cli_noteManager/storage"
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [text note]",
	Short: "Добавить новую заметку",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		store := &storage.JSONStorage{}
		notes, _ := store.LoadNotes()
		newNote := model.Note{
			ID:      len(notes) + 1,
			Content: args[0],
		}

		notes = append(notes, newNote)
		store.SaveNotes(notes)
		fmt.Println("Заметка добавлена!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
