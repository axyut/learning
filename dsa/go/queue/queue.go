package queue

import (
	"fmt"
	"strconv"
)

var max_size = 5
var queues = make([]int, max_size)

// var front = 0
var rear = -1 // top

func Run() {
	fmt.Println("\nQUEUE")
begin:
	fmt.Printf("\n1.Enqueue 2.Dequeue 3. Display 0.Exit\n: ")
	var num int
	fmt.Scanln(&num)
	switch num {
	case 0:
		goto end
	case 1:
		enqueue()
		display()
		goto begin
	case 2:
		dequeue_and_organize()
		display()
		goto begin
	case 3:
		display()
		goto begin
	default:
		fmt.Println("Enter 1,2,3 or 0.")
		goto begin
	}

end:
}

func isFull() bool {
	return rear == max_size-1 // array starts 0
}

func isEmpty() bool {
	return rear == -1
}

func enqueue() {
	if isFull() {
		fmt.Println("Queue Full. Cannot enqueue.")
	} else {
		fmt.Printf("Enqueue: ")
		var input int
		fmt.Scanf("%d", &input)

		rear += 1
		queues[rear] = input
		fmt.Println("Enqueued:", input)
	}
}

//	func dequeue() {
//		if isEmpty() {
//			fmt.Println("Queue Empty. Cannot dequeue.")
//		} else {
//			fmt.Printf("Dequeued: %d\n", queues[front])
//			queues[front] = 0
//			front += 1
//		}
//	}
func dequeue_and_organize() {
	if isEmpty() {
		fmt.Println("Queue Empty. Cannot dequeue.")
	} else {
		fmt.Printf("Dequeued: %d\n", queues[0])
		for i := 0; i <= rear; i++ {
			if i+1 > max_size-1 {
				queues[max_size-1] = 0
			} else {
				queues[i] = queues[i+1]
			}
		}
		rear -= 1
	}
}

func display() {
	for _, v := range queues {
		for range len(strconv.Itoa(v)) + 3 {
			fmt.Printf("-")
		}
	}
	fmt.Printf("-")
	fmt.Printf("\n")

	// main
	for _, v := range queues {
		fmt.Printf("| %d ", v)
	}

	fmt.Printf("|\n")
	for _, v := range queues {
		for range len(strconv.Itoa(v)) + 3 {
			fmt.Printf("-")
		}
	}
	fmt.Printf("-\n")
}
