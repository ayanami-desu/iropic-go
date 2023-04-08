package utils

func Intersect(ls1 []uint, ls2 []uint) []uint {
	var ls []uint
	i, j := 0, 0
	length1, length2 := len(ls1), len(ls2)
	for i < length1 && j < length2 {
		if ls1[i] > ls2[j] {
			j++
		} else if ls1[i] < ls2[j] {
			i++
		} else {
			ls = append(ls, ls1[i])
			i++
			j++
		}
	}
	return ls
}
