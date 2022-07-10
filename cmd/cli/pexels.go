package main

import (
	"fmt"
	"os"

	"github.com/JayMonari/pexels-go/cmd"
)

func main() {
	check(cmd.RootCmd.Execute())
}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
