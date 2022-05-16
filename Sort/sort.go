package Sort

import "fmt"

func Sort() {
	{ // Selection Sort
		fmt.Println(selectionSort([]int{64, 25, 12, 22, 11}))
		fmt.Println(selectionSort([]int{-1, 2, 0, 0, 0, -2}))
	}

	{ // Bubble Sort
		fmt.Println(bubbleSort([]int{64, 25, 22, 12, 11}))
		fmt.Println(bubbleSort([]int{-1, 2, 0, 0, 0, -2}))
	}

	{ // Insertion Sort
		fmt.Println(insertionSort([]int{64, 25, 22, 12, 11}))
		fmt.Println(insertionSort([]int{-1, 2, 0, 0, 0, -2}))
	}
}
