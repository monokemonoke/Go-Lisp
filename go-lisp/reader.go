package golisp

import "strings"

func Read(program string) []string {
	text := strings.Replace(program, "(", " ( ", -1)
	text = strings.Replace(text, ")", " ) ", -1)
	items := strings.Split(text, " ")
	var newItems []string
	for _, item := range items {
		if item != "" {
			newItems = append(newItems, item)
		}
	}
	return newItems
}
