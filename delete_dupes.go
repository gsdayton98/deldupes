// Package deldupes removes duplicate entries from a slice.
// This version returns a new copy of the slice sorted with duplicated entries removed.
package deldupes

import (
	"container/heap"
)

// HeapHelper provides a Heap interface to a slice.
// HeapHelper is exported for testing purposes.  Don't use it in normal use.
type HeapHelper []int

// Len returns the length of the container underlying the heap structure
func (self HeapHelper) Len() int           { return len(self) }

// Less provides the comparison function for building the heap.
// This Less implements a max-heap (an arbitrary design decision with no remarkable benefit).
func (self HeapHelper) Less(i, j int) bool { return self[j] < self[i] } // max heap

// Swap provides the implementation of swapping elements.  The Go heap implementation
// does not have access to the individual elements of the container, so it needs a
// Swap method to manipulate the container.
func (self HeapHelper) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

// Push is not actually used, but the Heap interface demands it.
// In other situations, if implemented, pushes a new element onto the end of the
// container.  This is a push on the container, not the heap structure.
func (self *HeapHelper) Push(x interface{}) {
	panic("Push function in deldupes isn't supposed to be used")
}

// Pop removes the end of the container and returns it.  This a pop operation
// on the container underlying the heap -- not the heap pop.  Heap.Pop swaps
// the end of the heap with the first element, then calls this function to
// shrink the heap.
func (self *HeapHelper) Pop() interface{} {
	old := *self
	n := len(old) - 1
	x := old[n]
	*self = old[0:n]
	return x
}


// DeleteDuplicates heap sorts a slice and returns a new slice with duplicate
// entries removed.
func DeleteDuplicates(theArray []int) []int {
	// Need to get a mutable copy of the heap
	heapHelper := make(HeapHelper, len(theArray))
	i := 0
	for _, value := range theArray {
		heapHelper[i] = value
		i++
	}
	heap.Init(&heapHelper)

	//  How to do this in-place?
	result := make([]int, 0)

	if len(theArray) > 0 {
		prevElement := heap.Pop(&heapHelper).(int)

		result = append(result, prevElement)

		for heapHelper.Len() > 0 {
			currentElement := heap.Pop(&heapHelper).(int)
			if currentElement != prevElement {
				result = append(result, currentElement)
				prevElement = currentElement
			}
		}
	}
	return result
}
