package heap

// Heap 通过数组切片存储二叉树节点
type Heap []int

func New() *Heap {
	return new(Heap)
}

func (h *Heap) Len() int {
	return len(*h)
}

// Less 比较元素大小
func (h *Heap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

// Swap 交换元素位置
func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Push 新增节点
func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.(int))
	i := h.Len() - 1 // 新增元素位置
	for {
		j := (i - 1) / 2 // 父节点位置
		// 如果是根节点或者父节点值小于子节点值，则退出循环
		if i == j || !h.Less(i, j) {
			break
		}
		// 否则交换子节点与父节点，直到父节点值小于子节点
		h.Swap(i, j)
		i = j
	}
}

// Pop 弹出堆顶元素并删除
func (h *Heap) Pop() (v interface{}) {
	// 将堆顶元素和最后一个元素交换位置，解决数组空洞问题
	n := h.Len() - 1
	h.Swap(0, n)
	// 剩余元素重新堆化
	i := 0
	for {
		j1 := 2*i + 1          // 左子节点
		if j1 >= n || j1 < 0 { // 数组越界，则退出
			break
		}
		j := j1
		// 如果i对应节点的右子节点比左子节点值小
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // 右子节点
		}
		// 左/右子节点值已经大于父节点值，则退出
		if !h.Less(j, i) {
			break
		}
		// 否则交换左/右子节点较小值对应节点与父节点的位置，继续循环，知道堆化成功
		h.Swap(i, j)
		i = j
	}
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1] // 注意此时最后一个节点（待删除节点）已经从切片中剔除，随时会被gc
	return v
}
