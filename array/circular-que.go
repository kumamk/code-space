package array

type circularQue struct {
	size      int
	head      int
	count     int // to store no of next operation
	windowSum int // to store updated sum
	que       []int
}

func NewCircularQue(size int) *circularQue {
	que := make([]int, size)
	return &circularQue{size: size, que: que, head: 0, count: 0, windowSum: 0}
}

func (cq *circularQue) Next(val int) float32 {
	cq.count++
	// remove ele
	tail := (cq.head + 1) % cq.size
	cq.windowSum = cq.windowSum - cq.que[tail] + val
	// add ele
	cq.head = (cq.head + 1) % cq.size
	cq.que[cq.head] = val

	return float32(cq.windowSum) / float32(min(cq.count, cq.size))
}
