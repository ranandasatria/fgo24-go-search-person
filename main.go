package main

import (
	"fmt"
	"strings"
)

func main() {
	users := []string{
		"Leanne Graham", 
		"Ervin Howell", 
		"Clementine Bauch", 
		"Patricia Lebsack", 
		"Chelsey Dietrich", 
		"Mrs. Dennis Schulist", 
		"Kurtis Weissnat", 
		"Nicholas Runolfsdottir V", 
		"Glena Reichert", 
		"Clementina DuBuque"}
	searchPerson(users, "asd")
}

func searchPerson(data []string, query string) {
	for _, person := range data {
		if strings.Contains(strings.ToLower(person), strings.ToLower(query)) {
			fmt.Println(person)
		} else {
			fmt.Println("Users not found")
			break
		}
	}
}