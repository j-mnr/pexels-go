package cmd

import (
	"fmt"
	"os"

	"github.com/JayMonari/pexels-go"
	"github.com/spf13/cobra"
)

var query string

var searchCmd = &cobra.Command{
	Use:       "search {photo video collection} -q query",
	Short:     "Search for photos, videos or collections",
	ValidArgs: []string{"photos", "videos", "collections"},
	Run: func(cmd *cobra.Command, args []string) {
		c, err := pexels.New(pexels.WithAPIKey(os.Getenv("PEXELS_API_KEY")))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		switch args[0][0] {
		case 'p':
			res, err := c.SearchPhotos(&pexels.PhotoSearchParams{Query: query,
				General: pexels.General{}})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)
		case 'v':
		case 'c':
		}
	},
}

func init() {
	searchCmd.Flags().StringVarP(&query, "query", "q", "", "What to search for")
	RootCmd.AddCommand(searchCmd)
}
