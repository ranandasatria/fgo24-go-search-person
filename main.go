package main

import (
	"playground/utils"
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
	utils.SearchPerson(users, "Gle")
	utils.SearchPerson(users, "Cle")
	utils.SearchPerson(users, "asd")
}
