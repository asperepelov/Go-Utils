package Utils

func Contains(a []string, x string) bool {
	for _, n := range a {
		if n == x {
			return true
		}
	}
	return false
}

func ContainsInt(arr []int, i int) bool {
	for _, n := range arr {
		if n == i {
			return true
		}
	}
	return false
}
