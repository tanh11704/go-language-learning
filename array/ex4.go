package array

func MergeSortedSlices(arr1, arr2 []int) []int {
	pointer1, pointer2 := 0, 0
	result := make([]int, 0, len(arr1)+len(arr2))

	for pointer1 < len(arr1) && pointer2 < len(arr2) {
		if arr1[pointer1] < arr2[pointer2] {
			result = append(result, arr1[pointer1])
			pointer1++
		} else {
			result = append(result, arr2[pointer2])
			pointer2++
		}
	}

	if pointer1 < len(arr1) {
		result = append(result, arr1[pointer1:]...)
	}

	if pointer2 < len(arr2) {
		result = append(result, arr2[pointer2:]...)
	}

	return result
}
