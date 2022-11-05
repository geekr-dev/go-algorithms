package heap_test

import (
	"fmt"
	"testing"

	"github.com/geekr-dev/go-algorithms/tree/heap"
)

func verify(t *testing.T, h *heap.Heap, i int) {
	t.Helper()
	n := h.Len()
	j1 := 2*i + 1 // 左子节点
	j2 := 2*i + 2 // 右子节点
	if j1 < n {
		if h.Less(j1, i) { // 左子节点小于父节点值
			t.Errorf("堆特性已破坏: [%d] = %d > [%d] = %d", i, (*h)[i], j1, (*h)[j1])
			return
		}
		verify(t, h, j1)
	}
	if j2 < n {
		if h.Less(j2, i) { // 右子节点小于父节点值
			t.Errorf("堆特性已破坏: [%d] = %d > [%d] = %d", i, (*h)[i], j1, (*h)[j2])
			return
		}
		verify(t, h, j2)
	}
}

func TestHeapPush(t *testing.T) {
	h := heap.New()
	for i := 10; i > 0; i-- {
		h.Push(i)
	}
	verify(t, h, 0)
}

func TestHeapPop(t *testing.T) {
	h := heap.New()
	// 先插入
	for i := 10; i > 0; i-- {
		h.Push(i)
	}
	// 再弹出
	for h.Len() > 0 {
		n := h.Pop()
		fmt.Printf("%d ", n)
		verify(t, h, 0)
	}
	fmt.Println()
}
