package twofer

import (
	"strings"
)

//If the given name is "Alice", the result should be "One for Alice, one for me."
// If no name is given, the result should be "One for you, one for me."
func ShareWith(name string) string {
	var res string
	if strings.Contains(name, "Alice") {
		res = "One for " + name + ", one for me."
	} else if name == "" {
		res = "One for you, one for me."
	} else {
		res = "One for " + name + ", one for me."
	}
	return res
}
