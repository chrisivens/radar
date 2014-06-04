package nlp

func TermFrequency(tokens []string) map[string]float32 {
	tf_map := make(map[string]float32)
	doc_length := len(tokens)

	for i := 0; i < doc_length; i++ {
		token := tokens[i]

		if _, ok := tf_map[token]; !ok {
			occ := float32(occurences(token, tokens)) / float32(doc_length)
			tf_map[token] = occ
		}
	}

	return tf_map
}

func occurences(token string, collection []string) int {
	count := 0
	for i := range collection {
		if collection[i] == token {
			count++
		}
	}

	return count
}
