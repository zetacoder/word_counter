package main

import (
	"bufio"
	"log"
	"os"
)

func processFromStdin() string {
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		//log.Printf("Read line: %s\n", line)
		text += line + " "
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return text
}
