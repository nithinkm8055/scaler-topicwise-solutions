package Homework

//▄▄        ▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄       ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄    ▄               ▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄
//▐░░▌      ▐░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌     ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌  ▐░▌             ▐░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░▌
//▐░▌░▌     ▐░▌▐░█▀▀▀▀▀▀▀█░▌ ▀▀▀▀█░█▀▀▀▀      ▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌▐░▌   ▐░▌           ▐░▌ ▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌
//▐░▌▐░▌    ▐░▌▐░▌       ▐░▌     ▐░▌          ▐░▌          ▐░▌       ▐░▌▐░▌    ▐░▌         ▐░▌  ▐░▌          ▐░▌       ▐░▌
//▐░▌ ▐░▌   ▐░▌▐░▌       ▐░▌     ▐░▌          ▐░█▄▄▄▄▄▄▄▄▄ ▐░▌       ▐░▌▐░▌     ▐░▌       ▐░▌   ▐░█▄▄▄▄▄▄▄▄▄ ▐░▌       ▐░▌
//▐░▌  ▐░▌  ▐░▌▐░▌       ▐░▌     ▐░▌          ▐░░░░░░░░░░░▌▐░▌       ▐░▌▐░▌      ▐░▌     ▐░▌    ▐░░░░░░░░░░░▌▐░▌       ▐░▌
//▐░▌   ▐░▌ ▐░▌▐░▌       ▐░▌     ▐░▌           ▀▀▀▀▀▀▀▀▀█░▌▐░▌       ▐░▌▐░▌       ▐░▌   ▐░▌     ▐░█▀▀▀▀▀▀▀▀▀ ▐░▌       ▐░▌
//▐░▌    ▐░▌▐░▌▐░▌       ▐░▌     ▐░▌                    ▐░▌▐░▌       ▐░▌▐░▌        ▐░▌ ▐░▌      ▐░▌          ▐░▌       ▐░▌
//▐░▌     ▐░▐░▌▐░█▄▄▄▄▄▄▄█░▌     ▐░▌           ▄▄▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄▄▄▐░▐░▌       ▐░█▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄█░▌
//▐░▌      ▐░░▌▐░░░░░░░░░░░▌     ▐░▌          ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌        ▐░░░░░░░░░░░▌▐░░░░░░░░░░▌
//▀        ▀▀  ▀▀▀▀▀▀▀▀▀▀▀       ▀            ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀          ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀

//Here are a few edge cases to consider:
//
//When the rectangle's dimensions are both 0 (x = 0 and y = 0).
//When the radius of the circles is 0 (R = 0).
//When the number of circles is 0 (N = 0).
//When the coordinates of the circles are outside the rectangle's boundaries.
//When the rectangle's dimensions are equal to the radius of the circles (x = R and y = R).
//Make sure your solution can handle these edge cases appropriately.

//Problem Description
//There is a rectangle with left bottom as (0, 0) and right up as (x, y).
//
//There are N circles such that their centers are inside the rectangle.
//
//Radius of each circle is R. Now we need to find out if it is possible that we can move from (0, 0) to (x, y) without touching any circle.
//
//Note : We can move from any cell to any of its 8 adjecent neighbours and we cannot move outside the boundary of the rectangle at any point of time.
//
//
//
//Problem Constraints
//0 <= x , y, R <= 100
//
//1 <= N <= 1000
//
//Center of each circle would lie within the grid
//
//
//
//Input Format
//1st argument given is an Integer x , denoted by A in input.
//
//2nd argument given is an Integer y, denoted by B in input.
//
//3rd argument given is an Integer N, number of circles, denoted by C in input.
//
//4th argument given is an Integer R, radius of each circle, denoted by D in input.
//
//5th argument given is an Array A of size N, denoted by E in input, where A[i] = x cordinate of ith circle
//
//6th argument given is an Array B of size N, denoted by F in input, where B[i] = y cordinate of ith circle
//
//
//
//Output Format
//Return YES or NO depending on weather it is possible to reach cell (x,y) or not starting from (0,0).
type Pair struct {
	x       int
	y       int
	visited int
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
func _solve(x int, y int, N int, R int, E []int, F []int) string {

	// coordinates (x,y) denote destination
	// N is the number of circle
	// R is the radius of circle, extends to all eight sides from the cell of a matrix

	// traverse in a matrix from (0,0) to (x,y)
	// E is an array with x coordinate of circle centre
	// F is an array with y coordinate of circle centre

	// 1. Initialize x*y matrix
	matrix := make([][]int, x)

	for i := range matrix {
		matrix[i] = make([]int, y)
	}

	// 2. Mark position of circle in matrix (denoted by 1)
	for i := range E {

		xC := E[i] - 1
		yC := F[i] - 1

		if xC >= 0 && xC < len(matrix) && yC >= 0 && yC < len(matrix[0]) {
			matrix[xC][yC] = 1
			mark_circle(xC, yC, R, matrix)
		}

	}

	// 3. Search for path from 0,0 to x,y
	q := new(Queue)
	node := new(Node)
	node.data = &Pair{0, 0, 1}
	q.enqueue(node)

	for !q.isEmpty() {
		dx := []int{-1, -1, 0, 1, 1, 1, 0, -1}
		dy := []int{0, 1, 1, 1, 0, -1, -1, -1}

		pair := q.front()
		q.dequeue()

		for i := range dx {
			xC := pair.x + dx[i]
			yC := pair.y + dy[i]

			if xC >= 0 && xC < len(matrix) && yC >= 0 && yC < len(matrix[0]) && matrix[xC][yC] == 0 {
				n := new(Node)
				n.data = &Pair{xC, yC, 1}
				if xC == x-1 && yC == y-1 {
					return "YES"
				}
				q.enqueue(n)

			}

		}
	}

	return "NO"
}

func mark_circle(x, y, R int, matrix [][]int) {

	dx := []int{-R, -R, 0, R, R, R, 0, -R}
	dy := []int{0, R, R, R, 0, -R, -R, -R}

	for i := range dx {

		xC := x + dx[i]
		yC := y + dy[i]

		if xC >= 0 && xC < len(matrix) && yC >= 0 && yC < len(matrix[0]) {
			matrix[xC][yC] = 1
		}
	}
}
