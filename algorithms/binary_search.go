package algorithms

import "errors"

func BinarySearch(array []int, target int) (int, error) {
	left := 0
	right := len(array) - 1

	for left <= right {
		middleIndex := right - (right-left)/2

		if array[middleIndex] == target {
			return middleIndex, nil
		} else if right > array[middleIndex] {
			left++
		} else {
			right--
		}
	}

	return -1, errors.New("not found")
}
