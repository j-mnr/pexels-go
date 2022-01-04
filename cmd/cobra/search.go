package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() { RootCmd.AddCommand(searchCmd) }

var searchCmd = &cobra.Command{
	Use:       "search { photos | videos | collections }",
	Short:     "Search pexels for photos/videos/collections",
	ValidArgs: []string{"photos", "videos", "collections"},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "photos":
			fmt.Println("You choose photo!")
		case "videos":
			fmt.Println("You choose video!")
		case "collections":
			fmt.Println("You choose collections!")
		default:
			cmd.Usage()
		}
	},
}
