package cmd

import (
	"cli_noteManager/storage"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Output all notes",
	Run: func(cmd *cobra.Command, args []string) {
		store := &storage.JSONStorage{}
		notes, err := store.LoadNotes()
		if err != nil {
			fmt.Println("The notes not exist.")
			return
		}

		for i, note := range notes {
			fmt.Printf(
				"The note with number %d in list: ID: %d, Content: %s \n",
				i,
				note.ID,
				note.Content,
			)
		}

		fmt.Println("List has been output.")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
