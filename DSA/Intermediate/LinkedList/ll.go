package LinkedList

import (
	"fmt"

	Assignment2 "github.com/nithinkm8055/scaler-topicwise-solutions/Intermediate/LinkedList/Assignment"
)

func LinkedList() {

	var limit int
	fmt.Scanf("%d", &limit)

	for i := 0; i < limit+100; i++ {
		var operation byte
		fmt.Scanf("%c", &operation)

		switch operation {
		case 'i':
			var position int
			var value int
			fmt.Scanf("%d", &position)
			fmt.Scanf("%d", &value)
			Assignment2.Insert_node(position, value)
		case 'p':
			Assignment2.Print_ll()
		case 'd':
			var position int
			fmt.Scanf("%d", &position)

			Assignment2.Delete_node(position)
		}

	}

}

func DesignLinkedList() {

	ll := Assignment2.DesignLinkedList([][]int{{1, 13, -1}, {3, 0, -1}, {3, 1, -1}, {2, 15, 0}, {3, 0, -1}, {1, 12, -1},
		{3, 0, -1}})
	temp := ll
	for temp.Next != nil {
		fmt.Print(temp.Value, "-->")
		temp = temp.Next
	}
	fmt.Print(temp.Value)

}
