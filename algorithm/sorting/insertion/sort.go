package insertion

// Time complexity:
// Best case: when all elements are sorted omega(n), because only
// comparison loop needs to be executed
// Worst case: when all the elements are not sorted theta(n^2)
func Sort(a []int) []int {
	for i := 1; i < len(a); i++ {
		key := a[i]
		// compare to the previous location
		// till the situation is satisfied
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j = j - 1
		}
		a[j+1] = key
	}
	return a
}
