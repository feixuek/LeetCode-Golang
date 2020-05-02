type RecentCounter struct {
	Data []int
}

func Constructor() RecentCounter {
	var rc RecentCounter
	return rc
}

func (rc *RecentCounter) Ping(t int) int {
	rc.Data = append(rc.Data, t)
	start := t - 3000
	ret := 0
	for i := 0; i < len(rc.Data); i++ {
		if rc.Data[i] >= start && rc.Data[i] <= t {
			ret++
		}
	}
	return ret
}

