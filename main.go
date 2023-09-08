package main

import "fmt"

func main() {
	if err := runWordCounter(); err != nil {
		fmt.Println(err.Error())
	}
}
