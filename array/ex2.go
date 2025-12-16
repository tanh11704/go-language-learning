package array

func BinarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	left, right := 0, len(arr)

	isAscending := arr[0] <= arr[len(arr)-1]

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if isAscending {
			if arr[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if arr[mid] < target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}
