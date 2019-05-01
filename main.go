package console

import (
	"bufio"
	"os"
	"regexp"
)

// WaitForInput nil->string
// Function for reading from console input
// Delays entire program until the user hits enter in the console
// Returns string user entered including the newline
func WaitForInput() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		if text != "" {
			return text
		}
	}
}

// FixString string->string
// Function to remove all alphanumeric characters from a string
// Takes input s as string and returns string containing only alphanumeric characters
func FixString(s string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(s, "")
}
