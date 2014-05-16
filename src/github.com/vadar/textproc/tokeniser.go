package textproc

import (
	"regexp"
	"strings"
)

func Tokenise(message string) (tokens, fields []string) {
	fields = strings.Fields(message)
	tokens = fields

	return
}

// TODO - Get the stopwords from the DB
func StopWords() []string {
	return stopwords
}

func RemoveStopwords(source string) string {
	var out string
	fields := strings.Fields(source)

	clean_fields := RemoveStopwordsFromSlice(fields)

	out = strings.Join(clean_fields, " ")

	return out
}

func RemoveStopwordsFromSlice(source []string) []string {
	stopwords := StopWords()

	var clean_fields []string = make([]string, len(source))

	for i := 0; i < len(source); i++ {
		for j := 0; j < len(stopwords); j++ {
			if strings.Contains(source[i], stopwords[j]) {
				re := regexp.MustCompile("(^|\\W+?)(" + stopwords[j] + ")($|\\W.*)")
				source[i] = re.ReplaceAllString(source[i], "$1$3")
			}
		}

		if source[i] != "" {
			clean_fields = append(clean_fields, source[i])
		}
	}

	return clean_fields
}
