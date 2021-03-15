package slice

func RemoveInt(s []int, elem int) []int {
	index := -1
	for i, v := range s {
		if v == elem {
			index = i
			break
		}
	}
	new := append(s[:index], s[index:]...)
	return new
}

func RemoveI64(s []int64, elem int64) []int64 {
	index := -1
	for i, v := range s {
		if v == elem {
			index = i
			break
		}
	}
	new := append(s[:index], s[index:]...)
	return new
}
