package main

import "strings"

func cleanInput(s string) []string {
	s = strings.ToLower(s)
	fields := strings.Fields(s)
	return fields
}
