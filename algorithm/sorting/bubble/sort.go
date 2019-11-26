package bubble

func Sort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[j] > a[i] {
				tmp := a[j]
				a[j] = a[i]
				a[i] = tmp
			}
		}
	}
	return a
}
