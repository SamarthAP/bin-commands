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
	Use:   "ls [path]",
	Short: "List directory contents",
	Long:  "List directory contents",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			directoryContents(".")
		} else {
			directoryContents(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
