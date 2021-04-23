package priorityQueue

import (
	"reflect"
	"testing"
)

func TestExtractMax(t *testing.T) {
	t.Run("returns node with maximum priority from a max heap", func(t *testing.T) {
		var heap = []int{0, 20, 5, 15, 4, 3, 1, 10}
		want := 20
		got := extractMax(&heap, len(heap)-1)

		if got != want {
			t.Errorf("wanted maximum priority node to be %v got %v", want, got)
		}
	})
	t.Run("max heapifies a heap after extracting maximum priority node", func(t *testing.T) {
		var heap = []int{0, 20, 5, 15, 4, 3, 1, 10}
		want := []int{0, 15, 5, 10, 4, 3, 1}
		extractMax(&heap, len(heap)-1)

		if !reflect.DeepEqual(heap, want) {
			t.Errorf("after extracting maximum priority node wanted heap to be %v got %v", want, heap)
		}
	})
}

func TestInsertNode(t *testing.T) {
	t.Run("returns new heap size after inserting a new node", func(t *testing.T) {
		var heap = []int{0, 20, 5, 15, 4, 3, 1, 10}
		want := 8
		got := insertNode(&heap, len(heap)-1, 30)

		if got != want {
			t.Errorf("after inserting, wanted new heap size to be %v got %v", want, got)
		}
	})
	t.Run("maintains the binary max heap property after inserting new node", func(t *testing.T) {
		var heap = []int{0, 20, 5, 15, 4, 3, 1, 10}
		want := []int{0, 30, 20, 15, 5, 3, 1, 10, 4}
		insertNode(&heap, len(heap)-1, 30)

		if !reflect.DeepEqual(heap, want) {
			t.Errorf("after inserting, wanted new heap to be %v got %v", want, heap)
		}
	})
}
