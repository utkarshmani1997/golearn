package merge

func Sort(a []int) []int {
	return sort(a)
}

// Time complexity of sort is O(nlogn)
// since T(n) = 2T(n/2) + O(n)
// Space complexity = klogn (stack) + O(n) (merge) ~ O(n)
func sort(a []int) []int {
	if len(a) == 1 {
		return a
	}
	middle := len(a) / 2
	//	left, right := make([]int, middle), make([]int, len(a)-middle)
	left, right := a[:middle], a[middle:]
	return merge(sort(left), sort(right))
}

func Merge(l, r []int) []int {
	return merge(l, r)
}

// Time complexity if merge operation is O(n)
// Space complexity is also O(n) (result(n) and i, j, k (constant))
func merge(left, right []int) []int {
	l, r := len(left), len(right)
	length := l + r
	result := make([]int, length)
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	// copy remaining elements
	for i < l {
		result[k] = left[i]
		i++
		k++
	}

	for j < r {
		result[k] = right[j]
		j++
		k++
	}

	return result
}
