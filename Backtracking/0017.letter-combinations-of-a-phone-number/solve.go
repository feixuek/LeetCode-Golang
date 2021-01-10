var (
	target = map[byte][]byte{
		'2': []byte("abc"),
		'3': []byte("def"),
		'4': []byte("ghi"),
		'5': []byte("jkl"),
		'6': []byte("mno"),
		'7': []byte("pqrs"),
		'8': []byte("tuv"),
		'9': []byte("wxyz"),
	}
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	res := make([]string, 0)
	backtrace([]byte(digits), &res, &[]byte{}, 0, 0)
	return res
}

func backtrace(digits []byte, res *[]string, trace *[]byte, start, innerStart int) {
	if len(*trace) == len(digits) {
		*res = append(*res, string(*trace))
		return
	}
	for i := start; i < len(digits); i++ {
		data := target[digits[i]]
		for j := innerStart; j < len(data); j++ {
			*trace = append(*trace, data[j])
			backtrace(digits, res, trace, i+1, innerStart)
			*trace = (*trace)[:len(*trace)-1]
		}
	}
}