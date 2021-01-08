package sort_test

import (
	"math/rand"
	"testing"

	sort "github.com/chrisdoherty4/books/introduction-to-algorithms/1.2/insertion-sort"
	"github.com/stretchr/testify/assert"
)

func TestInsertion(t *testing.T) {
	vals := []int{5, 4, 3, 2, 1}
	sorted := sort.Insertion(vals)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, sorted)
}

var unsortedInput []int
var packageSorted []int

func init() {
	const totalUnsortedInput = 1000
	unsortedInput = make([]int, totalUnsortedInput)
	for i := 0; i < totalUnsortedInput; i++ {
		unsortedInput[i] = rand.Intn(1000)
	}
}

func BenchmarkInsertion(b *testing.B) {
	b.ReportAllocs()
	var sorted []int

	for i := 0; i < b.N; i++ {
		sorted = sort.Insertion(unsortedInput)
	}

	packageSorted = sorted
}
