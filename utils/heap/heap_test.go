package heap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	h := &Heap[int]{compare: func(a, b int) int { return a - b }}
	tests := []struct {
		name            string
		actionFunc      func(*testing.T)
		expectedSize    int
		expectedPeek    int
		expectedPeekErr error
	}{
		{
			name: "TestHeap_Insert",
			actionFunc: func(_ *testing.T) {
				for _, num := range []int{5,3,7,1,2} {
					h.Insert(num)
				}
			},
			expectedSize: 5,
			expectedPeek: 7,
		},
		{
			name: "TestHeap_Pop",
			actionFunc: func(t *testing.T) {
				expecteds := []int{7, 5}
				for _, expected := range expecteds {
					if node, _ := h.Pop(); node != expected {
						t.Errorf("Expected %d but got %d", expected, node)
					}
				}
			},
			expectedSize: 3,
			expectedPeek: 3,
		},
		{
			name: "TestHeap_PopAll",
			actionFunc: func(_ *testing.T) {
				expecteds := []int{3, 2, 1}
				for _, expected := range expecteds {
					if node, err := h.Pop(); node != expected || err !=nil {
						t.Errorf("Expected %d but got %d", expected, node)
					}
				}
			},
			expectedSize:    0,
			expectedPeekErr: EmptyHeapError,
		},
		{
			name:            "TestHeap_Noop",
			actionFunc:      func(_ *testing.T) {},
			expectedSize:    0,
			expectedPeekErr: EmptyHeapError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.actionFunc(t)
			node, err := h.Peek()
			if err != tt.expectedPeekErr {
				t.Errorf("Heap peek unexpected error, got: %v, want: %v", err, tt.expectedPeekErr)
			}
			if err == nil && node != tt.expectedPeek {
				t.Errorf("Heap peek incorrect, got: %d, want: %d.", node, tt.expectedPeek)
			}
			if h.Size() != tt.expectedSize {
				t.Errorf("Heap size incorrect, got: %d, want: %d.", h.size, tt.expectedSize)
			}
		})
	}
}
