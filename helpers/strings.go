package helpers

import "strings"

func RemoveUnderscores(name string) string {
	words := strings.Split(name, "_")
	return strings.Join(words, " ")
}
