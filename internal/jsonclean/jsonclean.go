package jsonclean

// Clean compacts a given JSON-formatted byte slice by removing all the
// redundand characters: ' ', '\n', '\r', '\t', keeping the field and
// the values intact. This function does not allocates new memory and
// its time complexity is O(n).
func Clean(data []byte) []byte {
	isOutsideQuotes := true
	var isGoodChar bool
	length := len(data)
	skipped := 0

	for i := 0; i < length; i++ {
		isOutsideQuotes = (data[i] == '"') != isOutsideQuotes

		if isOutsideQuotes {
			isGoodChar = data[i] != ' ' && data[i] != '\n' && data[i] != '\r' && data[i] != '\t'
		}

		if isOutsideQuotes && !isGoodChar {
			skipped++

			continue
		}

		data[i-skipped] = data[i]
	}

	return data[:length-skipped]
}
