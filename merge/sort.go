package merge

import "golang.org/x/exp/constraints"

func Sort[T constraints.Ordered](arr []T) {
	copy(arr, sort(arr))
}

func sort[T constraints.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2

	return merge(sort(arr[:middle]), sort(arr[middle:]))
}

func merge[T constraints.Ordered](left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}

		// todo: consider index-based approach instead of re-slicing every time
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}
