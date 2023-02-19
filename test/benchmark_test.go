package test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/kucherenkovova/gosort/bubble"
	"github.com/kucherenkovova/gosort/merge"
	"github.com/kucherenkovova/gosort/qsort"
)

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := rand.Perm(100000)
		bubble.Sort(arr)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := rand.Perm(100000)
		merge.Sort(arr)
	}
}

func BenchmarkQsort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := rand.Perm(100000)
		qsort.Sort(arr)
	}
}

func BenchmarkStandardSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := rand.Perm(100000)
		sort.Ints(arr)
	}
}
