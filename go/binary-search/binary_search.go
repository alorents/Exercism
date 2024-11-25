package binarysearch

func SearchInts(list []int, key int) int {
	// empty list
	if len(list) == 0 {
		return -1
	}
	// search value outside of range
	if key < list[0] || key > list[len(list)-1] {
		return -1
	}

	for start, end := 0, len(list)-1; start <= end; {
		mid := (start + end) / 2
		switch {
		case list[mid] == key:
			// found
			return mid
		case list[mid] < key:
			// search "top" half
			start = mid + 1
		case list[mid] > key:
			// search "bottom" half
			end = mid - 1
		}
	}
	return -1
}
