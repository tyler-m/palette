package main

import (
	"fmt"
	"os"

	"github.com/tyler-m/palette/cmd"
)

func main() {
	err := cmd.RootCommand.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
