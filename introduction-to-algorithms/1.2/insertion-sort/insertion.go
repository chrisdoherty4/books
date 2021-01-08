package sort

func Insertion(vals []int) []int {
	if len(vals) == 0 {
		return vals
	}

	sortedVals := make([]int, len(vals))
	sortedVals[0] = vals[0]

	for i, v := range vals[1:] {
		insert(v, sortedVals, i)
	}

	return sortedVals
}

func insert(v int, sortedVals []int, totalAdded int) {
	i := totalAdded
	for ; i >= 0; i-- {
		if v < sortedVals[i] {
			sortedVals[i+1] = sortedVals[i]
		}
	}

	sortedVals[i+1] = v
}
