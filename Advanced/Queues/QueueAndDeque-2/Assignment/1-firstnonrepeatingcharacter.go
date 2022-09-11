package Assignment

type QueueNode struct {
	next *QueueNode
	val  byte
}

func newQueueNode(val byte) *QueueNode {
	node := new(QueueNode)
	node.val = val
	node.next = nil
	return node
}

// Solution Approach
// insert to tail everytime a new character comes,
// if character is already present in LL, remove it

// keep a charMap to keep track of characters
func FirstNonRepeatingCharacter(A string) string {

	charMap := make(map[byte]int)
	ans := make([]byte, 0)

	dummyNode := newQueueNode('0')
	var head = dummyNode

	var tail = dummyNode

	for i := range A {

		if _, ok := charMap[A[i]]; ok {
			// remove node from LL
			temp := head
			for temp.next != nil {
				if temp.next.val == A[i] {
					temp.next = temp.next.next
				} else {
					temp = temp.next
				}
			}
			tail = temp
		} else {
			nNode := newQueueNode(A[i])
			tail.next = nNode
			tail = tail.next
		}

		if head.next != nil {
			ans = append(ans, head.next.val)
		} else {
			ans = append(ans, '#')
		}

		charMap[A[i]]++
	}

	return string(ans)
}
