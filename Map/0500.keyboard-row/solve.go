import "strings"

var qRow = map[byte]bool{
	'q': true,
	'w': true,
	'e': true,
	'r': true,
	't': true,
	'y': true,
	'u': true,
	'i': true,
	'o': true,
	'p': true,
}

var aRow = map[byte]bool{
	'a': true,
	's': true,
	'd': true,
	'f': true,
	'g': true,
	'h': true,
	'j': true,
	'k': true,
	'l': true,
}

var zRow = map[byte]bool{
	'z': true,
	'x': true,
	'c': true,
	'v': true,
	'b': true,
	'n': true,
	'm': true,
}

func isRow(word string, row map[byte]bool) bool {
	for _, b := range []byte(word) {
		if _, ok := row[b]; !ok {
			return false
		}
	}
	return true
}

func findWords(words []string) []string {
	var ret []string
	for i := 0; i < len(words); i++ {
		word := strings.ToLower(words[i])
		if isRow(word, qRow) || isRow(word, aRow) || isRow(word, zRow) {
			ret = append(ret, words[i])
		}
	}
	return ret
}