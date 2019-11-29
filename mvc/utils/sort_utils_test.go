package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	els := []int{9, 8, 7, 6, 5}

	els = BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els), "Should have 5 elements returned")
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func TestBubbleSortBestCase(t *testing.T) {
	els := []int{5, 6, 7, 8, 9}

	els = BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els), "Should have 5 elements returned")
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func TestSortWorstCase(t *testing.T) {
	els := []int{9, 8, 7, 6, 5}

	els = Sort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els), "Should have 5 elements returned")
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func TestSortWorstLargeNumberCase(t *testing.T) {
	els := getElements(10001)

	els = Sort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 10001, len(els), "Should have 10001 elements returned")

	for i := 0; i < len(els)-1; i++ {
		assert.LessOrEqual(t, els[i], els[i+1])
	}

}

func TestSortBestCase(t *testing.T) {
	els := []int{5, 6, 7, 8, 9}

	els = Sort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els), "Should have 5 elements returned")
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func TestBubbleSortNilSlice(t *testing.T) {

	BubbleSort(nil)

	//assert.Nil(els)
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	els := getElements(5)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els), "Should have 5 elements returned")
	assert.EqualValues(t, 4, els[0])
	assert.EqualValues(t, 3, els[1])
	assert.EqualValues(t, 2, els[2])
	assert.EqualValues(t, 1, els[3])
	assert.EqualValues(t, 0, els[4])
}

func BenchmarkBubbleSort10Elements(b *testing.B) {
	els := getElements(10)

	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort1000Elements(b *testing.B) {
	els := getElements(1000)

	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort100000Elements(b *testing.B) {
	els := getElements(100000)

	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkGoSort10Elements(b *testing.B) {
	els := getElements(10)

	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkGoSort1000Elements(b *testing.B) {
	els := getElements(1000)

	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkGoSort100000Elements(b *testing.B) {
	els := getElements(100000)

	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkSort10Elements(b *testing.B) {
	els := getElements(10)

	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkSort1000Elements(b *testing.B) {
	els := getElements(1000)

	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkSort100000Elements(b *testing.B) {
	els := getElements(100000)

	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}
