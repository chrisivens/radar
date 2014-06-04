package nlp

import (
	"testing"
)

func TestTermFrequency(t *testing.T) {
	tokens := []string{"repeated", "unique", "solo", "repeated", "lonely"}
	tf := TermFrequency(tokens)

	if tf["repeated"] != 0.4 {
		t.Error("Term frequency miscalculated for word: repeated")
	}
}
