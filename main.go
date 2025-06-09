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
	searchPerson(users, "clemen")
}

func searchPerson(users []string, query string) {
	for _, person := range users {
		if strings.Contains(strings.ToLower(person), strings.ToLower(query)) {
			fmt.Println(person)
		}
	}
}