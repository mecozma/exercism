/*
Package twofer returns the string "One for X, one for me" where X is a
dynamic value for just "you".
*/
package twofer

import "fmt"

// ShareWith function returns a string based on the provided name.
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
