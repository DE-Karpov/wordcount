package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	row := os.Args[1]
	rows := strings.Fields(row)
	fmt.Println(len(rows))
}
