package textproc

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestStopwords(t *testing.T) {
	source_string := "the man and the dog are cake eaters"

	cleaned := RemoveStopwords(source_string)

	fmt.Println(cleaned)
	stopwords := StopWords()

	if !strings.Contains(cleaned, "man") {
		t.Error("Stopwords removed a word it shouldn't have done: man")
	}

	if !strings.Contains(cleaned, "dog") {
		t.Error("Stopwords removed a word it shouldn't have done: dog")
	}

	if !strings.Contains(cleaned, "cake") {
		t.Error("Stopwords removed a word it shouldn't have done: cake")
	}

	if !strings.Contains(cleaned, "eaters") {
		t.Error("Stopwords removed a word it shouldn't have done: eaters")
	}

	for i := 0; i < len(stopwords); i++ {
		re := regexp.MustCompile("(^|\\W+?)(" + stopwords[i] + ")($|\\W.*)")

		if re.MatchString(cleaned) {
			t.Error("Stopword " + stopwords[i] + " found in string")
		}
	}
}
