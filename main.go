package main

import "log"

func main() {
	if err := runWordCounter(); err != nil {
		log.Fatal(err)
	}
}
