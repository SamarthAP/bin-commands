package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

func directoryContents(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name() + "/")
		} else {
			fmt.Println(file.Name())
		}
	}
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List directory contents",
	Long:  "List directory contents",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path")
		if path.Value.String() != "" {
			directoryContents(path.Value.String())
		} else {
			directoryContents(".")
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().StringP("path", "p", "", "List directory contents of specified path")
}
