func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, v := range bills {
		switch v {
		case 5:
			five++
		case 10:
			if five == 0 {
				return false
			}
			five--
			ten++
		case 20:
			switch {
			case five > 0 && ten > 0:
				five--
				ten--
			case five >= 3:
				five -= 3
			default:
				return false
			}
		}
	}
	return true
}