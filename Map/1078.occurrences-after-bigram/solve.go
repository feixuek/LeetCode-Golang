import "strings"

func findOcurrences(text string, first string, second string) []string {
	var res []string
	words := strings.Split(text, " ")
	for i := 0; i < len(words)-1; i++ {
		if words[i] == first && words[i+1] == second {
			if i+2 < len(words) {
				res = append(res, words[i+2])
			}
		}
	}
	return res
}