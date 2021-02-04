//Show command line arguments with index of argument using strconv string ot int conversion
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
