package main

import "log"

func main() {
	application, err := NewApplication()
	if err != nil {
		log.Fatal(err)
	}
	application.Run()
}
