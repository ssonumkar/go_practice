package main

import (
	"fmt"
	"practice.com/sagar/go_practice/data_structures/queue"
	"practice.com/sagar/go_practice/data_structures/stack"
)

func main() {

	fmt.Println("1.STACK\n2.QUEUE")
	var dOption int
	fmt.Scanf("%d", &dOption)
	switch dOption {
	case 1:
		{
			var s stack.Stack
			var ch int
			for {
				fmt.Println("1.PUSH\n2.POP\n3.TOP\n4.Display")
				fmt.Scanf("%d", &ch)
				switch ch {
				case 1:
					{
						var val string
						fmt.Scanf("%s", &val)
						s.Push(val)
					}
				case 2:
					{
						if s.Pop() {
							fmt.Println("Popped")
						} else {
							fmt.Println("Stack empty")
						}
					}
				case 3:
					fmt.Println(s.Top())

				case 4:
					s.PrintStack()
				}
				var op int
				fmt.Println("Press 1 to continue? ")
				fmt.Scanf("%d", &op)
				if op != 1 {
					break
				}
			}
		}
	case 2:
		{
			var q queue.Queue
			var ch int
			for {
				fmt.Println("1.ENQUEUE\n2.DEQUEUE\n3.QUEUE FRONT\n4.QUEUE REAR\n5.Display")
				fmt.Scanf("%d", &ch)
				switch ch {
				case 1:
					{
						var val string
						fmt.Scanf("%s", &val)
						q.Enqueue(val)
					}
				case 2:
					{
						if q.Dequeue() {
							fmt.Println("Dequeued")
						} else {
							fmt.Println("Queue empty")
						}
					}
				case 3:
					fmt.Println(q.GetFrontElement())

				case 4:
					fmt.Println(q.GetRearElement())

				case 5:
					fmt.Println(q)
				}
				var op int
				fmt.Println("Press 1 to continue? ")
				fmt.Scanf("%d", &op)
				if op != 1 {
					break
				}
			}
		}

	}

}
