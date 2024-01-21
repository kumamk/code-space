package array

//https://leetcode.com/problems/sliding-window-maximum/

type deque struct {
	list []int
}

func (d *deque) empty() bool {
	return len(d.list) == 0
}

func (d *deque) push(ele int) {
	d.list = append(d.list, ele)
}

func (d *deque) getFront() int {
	return d.list[0]
}

func (d *deque) popFront() {
	d.list = d.list[1:]
}

func (d *deque) getBack() int {
	return d.list[len(d.list)-1]
}

func (d *deque) popBack() {
	d.list = d.list[:len(d.list)-1]
}

func MaxSlidingWindow(nums []int, k int) []int {
	var result []int
	dq := &deque{}

	for i := range nums {

		for !dq.empty() && dq.getFront() <= i-k {
			dq.popFront()
		}

		for !dq.empty() && nums[i] > nums[dq.getBack()] {
			dq.popBack()
		}

		dq.push(i)

		if i >= k-1 {
			result = append(result, nums[dq.getFront()])
		}
	}
	return result
}
