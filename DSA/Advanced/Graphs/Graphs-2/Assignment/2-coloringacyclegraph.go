package Assignment

//Given the number of vertices A in a Cyclic Graph.
//
//Your task is to determine the minimum number of colors required to color the graph so that no two Adjacent vertices have the same color.
//
//
//A cyclic graph with A vertices is a graph with A edges, such that it forms a loop. See example test case explanation for more details.

func _solve(A int) int {
	if A%2 == 0 {

		return 2
	}

	return 3
}
