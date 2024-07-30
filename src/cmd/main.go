package main

import (
	"flag"
	"fmt"
	"os"
)

func contactArrIntoString(arr []string) string {
	var result string
	for i, v := range arr {
		if i == 0 {
			result = v
		} else {
			result += " " + v
		}
	}
	return result
}

func main() {
	createCommandFlag := flag.Bool("c", false, "Explain the command")

	flag.Parse()

	switch *createCommandFlag {
	case true:
		fmt.Println(contactArrIntoString(os.Args[2:]))
		break
	case false:
		fmt.Println(contactArrIntoString(os.Args[1:]))
		break
	}
}
