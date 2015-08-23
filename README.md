# deleteDuplicates-go

Delete duplicates from an array slice using a modified heap sort.

On the master branch, DeleteDuplicates() makes a max-heap, then interatively
pops the max off the heap, but pushes it onto the result slice only if it is
not a duplicate.  Should have O(n*log n) scale performance.
