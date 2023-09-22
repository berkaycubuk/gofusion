package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Usage: go run . <command> [args...]")
		os.Exit(1)
	}

	command := os.Args[0]

	if strings.Compare("new:pkg", command) == 1 {
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . new:pkg package-name")
			os.Exit(1)
		}

		packageName := os.Args[2]

		err := os.Mkdir("pkg/"+packageName, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("%s successfully created!", packageName)
	}

	for _, v := range os.Args {
		fmt.Println(v)
	}

	fmt.Println(len(os.Args))
}
