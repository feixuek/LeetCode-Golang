import "strconv"

func robotSim(commands []int, obstacles [][]int) int {
	obsmap := make(map[string]bool)
	for _, obs := range obstacles {
		obsmap[strconv.Itoa(obs[0])+" "+strconv.Itoa(obs[1])] = true
	}
	steps := [][]int{
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
		[]int{-1, 0},
	}
	res, x, y, direction := 0, 0, 0, 0
	for _, cmd := range commands {
		switch cmd {
		case -1:
			direction++
			if direction == 4 {
				direction = 0
			}
		case -2:
			direction--
			if direction == -1 {
				direction = 3
			}
		default:
			for cmd > 0 {
				idx := strconv.Itoa(x+steps[direction][0]) + " " + strconv.Itoa(y+steps[direction][1])
				if exist, ok := obsmap[idx]; !ok || !exist {
					x += steps[direction][0]
					y += steps[direction][1]
				}
				cmd--
			}
		}
		tmp := x*x + y*y
		if tmp > res {
			res = tmp
		}
	}
	return res
}