// Test deldupes/DeleteDuplicates
package deldupes

import (
	"sort"
	"testing"
	"github.com/stretchr/testify/assert"
)

// TestBasic tests the basic functionality of DeleteDuplicates.
// Test cases cover a normal array with no duplicates, empty slices,
// duplicates at the beginning, middle, end, and small slices.
func TestBasic(t *testing.T) {
	var testCases = []struct {
		sample   []int
		expected []int
	}{
		{[]int{3, 2, 1, 4}, []int{1, 2, 3, 4}},             // normal no duplicates
		{[]int{}, []int{}},                                 // empty input
		{[]int{1}, []int{1}},                               // single element input
		{[]int{2, 1}, []int{1, 2}},                         // two element input no dups
		{[]int{2, 2}, []int{2}},                            // two element input all dups
		{[]int{3, 2, 1}, []int{1, 2, 3}},                   // three element input no dups
		{[]int{3, 2, 1, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}}, // even number no dups
		{[]int{3, 3, 1, 4, 5, 6}, []int{1, 3, 4, 5, 6}},    // even number 1 dup beginning
		{[]int{3, 2, 1, 4, 6, 6}, []int{1, 2, 3, 4, 6}},    // even number 1 dup end
		{[]int{3, 1, 1, 1, 6, 6}, []int{1, 3, 6}},          // even number 1 dup end
	}

	for _, testCase := range testCases {
		actual := DeleteDuplicates(testCase.sample)
		sort.Ints(actual)
		assert.Equal(t, testCase.expected, actual, "expected not matched")
	}
}

// TestUnusedFunction checks that application will panic if deldupes.Push is ever called.  Panic on calling
// deldupes.Push is reasonable because it could only occur on sever corruption.
func TestUnusedFunction(t *testing.T) {
	errorSynthesizer := make(HeapHelper, 42)
	assert.Panics(t, func() { errorSynthesizer.Push(0x666) }, "Internal function Push should panic")

}

