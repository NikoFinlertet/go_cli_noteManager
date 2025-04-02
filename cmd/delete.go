package cmd

import (
	"cli_noteManager/storage"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [text note]",
	Short: "Delete the note",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		store := &storage.JSONStorage{}
		notes, _ := store.LoadNotes()

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("The ID is integer")
			return
		}

		// Search note with id
		index := -1
		for i, note := range notes {
			if note.ID == id {
				index = i
				break
			}
		}

		// if note isn't found
		if index == -1 {
			fmt.Printf("The note with ID:%d is not exist.", id)
			return
		}

		// Delete note
		notes = append(notes[:index], notes[index+1:]...)
		store.SaveNotes(notes)

		fmt.Printf("The note with ID:%d has been deleted", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
