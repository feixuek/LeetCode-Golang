import "fmt"

var (
	data = []int{1, 2, 4, 8, 1, 2, 4, 8, 16, 32}
)

func readBinaryWatch(num int) []string {
	res := []string{}
	backtrace(num, 0, 0, 0, &res)
	return res
}

func backtrace(num, hour, min, start int, res *[]string) {
	if num == 0 {
		if min < 10 {
			*res = append(*res, fmt.Sprintf("%d:0%d", hour, min))
		} else {
			*res = append(*res, fmt.Sprintf("%d:%d", hour, min))
		}
		return
	}
	for i := start; i < 10 && num > 0; i++ {
		if i < 4 && hour+data[i] < 12 {
			backtrace(num-1, hour+data[i], min, i+1, res)
		}
		if i >= 4 && min+data[i] < 60 {
			backtrace(num-1, hour, min+data[i], i+1, res)
		}
	}
}