type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	helper := make(map[int]*Employee)
	for i := 0; i < len(employees); i++ {
		helper[employees[i].Id] = employees[i]
	}
	return cal(helper, id)
}

func cal(employees map[int]*Employee, id int) int {
	res := 0
	e, ok := employees[id]
	if !ok {
		return res
	}
	res += e.Importance
	for i := 0; i < len(e.Subordinates); i++ {
		res += cal(employees, e.Subordinates[i])
	}
	return res
}