package textproc

import "strings"

func Tokenise(message string) []string {
	split_message := strings.Fields(message)
	return split_message
}
