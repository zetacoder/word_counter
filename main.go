package main

import (
	"fmt"
	"log"
)

func main() {
	if err := runWordCounter(); err != nil {
		fmt.Println(err.Error())
	}

	log.Println("Program finished")
}
