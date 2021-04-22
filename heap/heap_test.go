package heap

import (
	"reflect"
	"testing"
)

func TestNodeFinders(t *testing.T) {
	t.Run("returns index of left child node", func(t *testing.T) {

		got := leftNodeIndex(1)
		want := 2

		if got != want {
			t.Errorf("wanted left node index %v got %v", want, got)
		}
	})
	t.Run("returns index of right child node", func(t *testing.T) {
		got := rightNodeIndex(1)
		want := 3

		if got != want {
			t.Errorf("wanted right node index %v got %v", want, got)
		}
	})
	t.Run("returns index of parent node", func(t *testing.T) {
		got := parentIndex(3)
		want := 1

		if got != want {
			t.Errorf("wanted right node index %v got %v", want, got)
		}
	})
}

func TestMaxHeapify(t *testing.T) {
	var heap = []int{0, 10, 5, 20, 4, 3, 1, 15}

	want := []int{0, 20, 5, 15, 4, 3, 1, 10}
	maxHeapify(&heap, 8, 1)

	if !reflect.DeepEqual(want, heap) {
		t.Errorf("wanted heap %v got %v", want, heap)
	}
}

func TestBuildMaxHeap(t *testing.T) {
	var heap = []int{0, 10, 5, 20, 4, 3, 1, 15}

	want := []int{0, 20, 5, 15, 4, 3, 1, 10}
	buildMaxHeap(&heap, len(heap)-1)

	if !reflect.DeepEqual(want, heap) {
		t.Errorf("wanted built heap to be %v got %v", want, heap)
	}
}
