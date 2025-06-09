package utils

import (
	"fmt"
	"strings"
)

func SearchPerson(data []string, query string) {
	found := false
	for _, person := range data {
		if strings.Contains(strings.ToLower(person), strings.ToLower(query)) {
			fmt.Println(person)
			found = true
		}
	}
	if !found {
		fmt.Println("Users not found")
	}
}