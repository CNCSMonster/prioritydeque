// a double end priority queue based on container/heap
package prioritydeque

import "container/heap"

type priorityDequeNode struct {
	v        any
	maxIndex int //index in max heap
	minIndex int //index in min heap
}

type priorityque struct {
	Arr    []*priorityDequeNode
	LessFn func(i, j int) bool
	SwapFn func(i, j int)
}

func (pq priorityque) Len() int           { return len(pq.Arr) }
func (pq priorityque) Swap(i, j int)      { pq.SwapFn(i, j) }
func (pq priorityque) Less(i, j int) bool { return pq.LessFn(i, j) }

func (pq *priorityque) Push(v any) {
	pq.Arr = append(pq.Arr, v.(*priorityDequeNode))
}

func (pq *priorityque) Pop() any {
	out := pq.Arr[pq.Len()-1]
	pq.Arr = pq.Arr[:pq.Len()-1]
	return out
}

type PriorityDeque struct {
	minHeap priorityque
	maxHeap priorityque
}

// to use PriorityDeque ,you need to call this function to get a prepared one.
func New(Less func(v1, v2 any) bool) *PriorityDeque {
	out := &PriorityDeque{}
	out.maxHeap.LessFn = func(i, j int) bool {
		return !Less(out.maxHeap.Arr[i].v, out.maxHeap.Arr[j].v)
	}

	out.maxHeap.SwapFn = func(i, j int) {
		out.maxHeap.Arr[i], out.maxHeap.Arr[j] = out.maxHeap.Arr[j], out.maxHeap.Arr[i]
		// swap max index
		out.maxHeap.Arr[i].minIndex, out.maxHeap.Arr[j].minIndex = out.maxHeap.Arr[j].minIndex, out.maxHeap.Arr[i].minIndex
	}

	out.minHeap.LessFn = func(i, j int) bool {
		return Less(out.minHeap.Arr[i].v, out.minHeap.Arr[j].v)
	}
	out.minHeap.SwapFn = func(i, j int) {
		out.minHeap.Arr[i], out.minHeap.Arr[j] = out.minHeap.Arr[j], out.minHeap.Arr[i]
		out.minHeap.Arr[i].maxIndex, out.minHeap.Arr[j].maxIndex = out.minHeap.Arr[j].maxIndex, out.minHeap.Arr[i].maxIndex
	}

	return out
}

func FromSlice(Less func(v1, v2 any) bool, v ...any) *PriorityDeque {
	out := New(Less)
	for _, v := range v {
		out.Push(v)
	}
	return out
}

// push any value in the priorityque,which will take O(logn) time
func (priorityDeque *PriorityDeque) Push(v any) {
	toPush := &priorityDequeNode{
		v:        v,
		maxIndex: priorityDeque.Len(),
		minIndex: priorityDeque.Len(),
	}
	heap.Push(&priorityDeque.maxHeap, toPush)
	heap.Push(&priorityDeque.minHeap, toPush)
}

// remove and return the min value,which will take O(logn) time
func (priorityDeque *PriorityDeque) PopMin() any {
	minPDN := heap.Pop(&priorityDeque.minHeap).(*priorityDequeNode)
	maxIndex := minPDN.maxIndex
	heap.Remove(&priorityDeque.maxHeap, maxIndex)
	return minPDN.v
}

// remove and return the max value,taking O(logn) time
func (priorityDeque *PriorityDeque) PopMax() any {
	maxPDN := heap.Pop(&priorityDeque.maxHeap).(*priorityDequeNode)
	minIndex := maxPDN.minIndex
	heap.Remove(&priorityDeque.minHeap, minIndex)
	return maxPDN.v
}

// get current max value,taking O(1) time
func (priorityDeque PriorityDeque) Max() any {
	return priorityDeque.maxHeap.Arr[0].v
}

// get current min value,taking O(1) time
func (priorityDeque PriorityDeque) Min() any {
	return priorityDeque.minHeap.Arr[0].v
}

// this method will replace Max value with the given value,and Fix the heap from the place
// which equals to pop max value and push the given value in it,but will work faster,which takes O(logn) time
// the old max value will be returned
func (priorityDeque *PriorityDeque) ReplaceMax(v any) any {
	maxPDN := priorityDeque.maxHeap.Arr[0]
	out := maxPDN.v
	maxPDN.v = v
	minIndex := maxPDN.minIndex
	heap.Fix(&priorityDeque.maxHeap, 0)
	heap.Fix(&priorityDeque.minHeap, minIndex)
	return out
}

// this method will replace Min value with the given value,and Fix the heap from the place
// which equals to pop min value and push the given value in it,but will work faster,which takes O(logn) time
// the old min value will be returned
func (priorityDeque *PriorityDeque) ReplaceMin(v any) any {
	minPDN := priorityDeque.minHeap.Arr[0]
	out := minPDN.v
	minPDN.v = v
	maxIndex := minPDN.maxIndex
	heap.Fix(&priorityDeque.minHeap, 0)
	heap.Fix(&priorityDeque.maxHeap, maxIndex)
	return out
}

// get num of elements still in prioritydeque,taking O(1) time (which depends on a efficient built-in function len)
func (priorityDeque PriorityDeque) Len() int {
	return priorityDeque.minHeap.Len()
}
