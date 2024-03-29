package heap

import (
	"reflect"
	"testing"
)

func TestNodeFinders(t *testing.T) {
	t.Run("returns index of left child node", func(t *testing.T) {

		got := LeftNodeIndex(1)
		want := 2

		if got != want {
			t.Errorf("wanted left node index %v got %v", want, got)
		}
	})
	t.Run("returns index of right child node", func(t *testing.T) {
		got := RightNodeIndex(1)
		want := 3

		if got != want {
			t.Errorf("wanted right node index %v got %v", want, got)
		}
	})
	t.Run("returns index of parent node", func(t *testing.T) {
		got := ParentIndex(3)
		want := 1

		if got != want {
			t.Errorf("wanted right node index %v got %v", want, got)
		}
	})
}

func TestMaxHeapify(t *testing.T) {
	var heap = []int{0, 10, 5, 20, 4, 3, 1, 15}

	want := []int{0, 20, 5, 15, 4, 3, 1, 10}
	MaxHeapify(&heap, len(heap)-1, 1)

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

func TestMinHeapify(t *testing.T) {
	var heap = []int{0, 10, 5, 20}

	want := []int{0, 5, 10, 20}
	minHeapify(&heap, len(heap)-1, 1)

	if !reflect.DeepEqual(want, heap) {
		t.Errorf("wanted heap %v got %v", want, heap)
	}
}

func TestBuildMinHeap(t *testing.T) {
	var heap = []int{0, 10, 5, 20, 4, 3, 1, 15}

	want := []int{0, 1, 3, 10, 4, 5, 20, 15}
	buildMinHeap(&heap, len(heap)-1)

	if !reflect.DeepEqual(want, heap) {
		t.Errorf("wanted built heap to be %v got %v", want, heap)
	}
}

func TestHeapSortAscending(t *testing.T) {
	var heap = []int{0, 10, 5, 20, 4, 3, 1, 15}

	want := []int{0, 1, 3, 4, 5, 10, 15, 20}
	heapSortAscending(&heap, len(heap)-1)

	if !reflect.DeepEqual(want, heap) {
		t.Errorf("after sorting using max heapify wanted %v got %v", want, heap)
	}
}

func TestHeapSortDescending(t *testing.T) {
	var heap = []int{0, 10, 5, 20, 4, 3, 1, 15}

	want := []int{0, 20, 15, 10, 5, 4, 3, 1}
	heapSortDescending(&heap, len(heap)-1)

	if !reflect.DeepEqual(want, heap) {
		t.Errorf("after sorting using min heapify wanted %v got %v", want, heap)
	}
}
