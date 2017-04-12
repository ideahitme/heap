package heap

import "fmt"

func ExampleMaxHeap() {
	maxHeap := NewMaxHeap()
	maxHeap.Insert(1)
	maxHeap.Insert(2)
	fmt.Println(maxHeap.Top())
	maxHeap.Insert(5)
	fmt.Println(maxHeap.Pop())
	fmt.Println(maxHeap.Pop())
	fmt.Println(maxHeap.Pop())
	fmt.Println(maxHeap.Pop())
	//Output:
	//2 <nil>
	//5 <nil>
	//2 <nil>
	//1 <nil>
	//0 heap is empty

}

func ExampleMinHeap() {
	minHeap := NewMinHeap()
	minHeap.Insert(6)
	fmt.Println(minHeap.Top())
	minHeap.Insert(2)
	fmt.Println(minHeap.Top())
	minHeap.Insert(5)
	fmt.Println(minHeap.Pop())
	fmt.Println(minHeap.Pop())
	fmt.Println(minHeap.Pop())
	fmt.Println(minHeap.Pop())
	//Output:
	//6 <nil>
	//2 <nil>
	//2 <nil>
	//5 <nil>
	//6 <nil>
	//0 heap is empty
}
