package qsort

import "golang.org/x/exp/constraints"

func Sort[T constraints.Ordered](arr []T) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort[T constraints.Ordered](arr []T, low, high int) {
	if low >= high {
		return
	}

	pivot := partition(arr, low, high)
	quickSort(arr, low, pivot-1)
	quickSort(arr, pivot+1, high)
}

func partition[T constraints.Ordered](arr []T, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
