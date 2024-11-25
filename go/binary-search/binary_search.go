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
	start := 0
	end := len(list) - 1
	mid := (start + end) / 2

	for {
		// found
		if list[mid] == key {
			return mid
		}
		// narrowed to one option and not found
		if start == end {
			return -1
		}
		// in "bottom" half
		if list[mid] > key {
			// adjust for even sized slices without a true midpoint
			if (mid-start)%2 == 1 {
				end = mid - 1
			} else {
				end = mid
			}
			mid = (start + end) / 2
		}
		// in "top" half
		if list[mid] < key {
			// adjust for even sized slices without a true midpoint
			if (end-mid)%2 == 1 {
				start = mid + 1
			} else {
				start = mid
			}
			mid = (start + end) / 2
		}
	}
}
