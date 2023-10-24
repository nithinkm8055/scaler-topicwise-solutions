package Assignment

//Given a matrix of integers A of size N x M consisting of 0, 1 or 2.
//
//Each cell can have three values:
//
//The value 0 representing an empty cell.
//
//The value 1 representing a fresh orange.
//
//The value 2 representing a rotten orange.
//
//Every minute, any fresh orange that is adjacent (Left, Right, Top, or Bottom) to a rotten orange becomes rotten. Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1 instead.
//
//Note: Your solution will run on multiple test cases. If you are using global variables, make sure to clear them.

//Output Format
//Return the minimum number of minutes that must elapse until no cell has a fresh orange.
//
//If this is impossible, return -1 instead.

type Pair struct {
	x    int
	y    int
	time int
}
type Node struct {
	data *Pair
	next *Node
}
type Queue struct {
	head     *Node
	tail     *Node
	capacity int
}

func (q *Queue) enqueue(node *Node) {
	if q.isEmpty() {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = q.tail.next
	}
	q.capacity = q.capacity + 1
}

func (q *Queue) dequeue() {
	if !q.isEmpty() {
		q.head = q.head.next
		q.capacity = q.capacity - 1
	}
}

func (q *Queue) front() *Pair {
	if !q.isEmpty() {
		return q.head.data
	}
	return nil
}

func (q *Queue) isEmpty() bool {
	return q.capacity == 0
}

func solve(A [][]int) int {
	q := new(Queue)
	timeElapsed := -1
	// traverse through matrix and queue positions of rotten oranges
	// perform bfs with rotten oranges as starters in queue
	// traverse through matrix again to see if all oranges are picked

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {

			if A[i][j] == 2 {
				newNode := new(Node)
				newNode.data = &Pair{i, j, 0}
				q.enqueue(newNode)
			}
		}
	}

	// if queue is empty, means there are no rotten oranges

	for !q.isEmpty() {

		dx := []int{-1, 0, 1, 0}
		dy := []int{0, 1, 0, -1}

		pair := q.front()
		q.dequeue()

		for i := 0; i < len(dx); i++ {

			x := pair.x + dx[i]
			y := pair.y + dy[i]

			if x >= 0 && x < len(A) && y >= 0 && y < len(A[0]) && A[x][y] == 1 {
				A[x][y] = 2
				node := new(Node)
				node.data = &Pair{x, y, pair.time + 1}
				if pair.time+1 > timeElapsed {
					timeElapsed = pair.time + 1
				}
				q.enqueue(node)
			}
		}
	}

	// traverse to find there are no cells with fresh oranges (value 1)
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] == 1 {
				return -1
			}
		}
	}

	return timeElapsed
}
