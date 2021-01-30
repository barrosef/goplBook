package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	for i, value := range os.Args {
		fmt.Println(strconv.Itoa(i) + " " + value + " ")
	}
}
