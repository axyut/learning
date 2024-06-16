package queue

import "fmt"

var max_size = 5
var queues [5]int
var front = 0
var rear = -1 // top

func Run() {
	fmt.Println("\n------- QUEUE -------")
begin:
	fmt.Printf("\n1.Enqueue 2.Dequeue 3. Display 0.Exit\nYour Input: ")
	var num int
	fmt.Scanln(&num)
	switch num {
	case 0:
		goto end
	case 1:
		enqueue()
		goto begin
	case 2:
		dequeue()
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
		fmt.Scanln(&input)

		rear += 1
		queues[rear] = input
		fmt.Println(input, " Enqueued.")
	}
}

func dequeue() {
	if isEmpty() {
		fmt.Println("Queue Empty. Cannot dequeue.")
	} else {
		fmt.Printf("delete %d?", queues[front])
		queues[front] = 0
		front += 1
	}
}

func display() {
	for i, v := range queues {
		fmt.Println(i, v)
	}
}
