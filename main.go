package main

import (
	"goplay/Adhoc"
	"goplay/Arrays"
	"goplay/LinkedList"
	"goplay/Maps"
	"goplay/Maps/Day1"
	"goplay/Maths"
	"goplay/Modular"
	"goplay/Sort"
	"goplay/Sorting"
	"goplay/StacksnQueues"
	"goplay/Strings"
	"goplay/Trees"
)

func main() {

	{ // Selection Sort, Bubble Sort, Insertion Sort
		Sort.Sort()
	}

	{
		Arrays.Arrays()
	}

	{
		Modular.Modular()
	}

	{
		Adhoc.Adhoc()
	}
	{
		Maths.Maths()
	}

	{
		Sorting.Sorting()
	}

	{
		Strings.Strings()
	}

	{
		//Contest.Contest()
	}

	{
		Day1.Maps()
		Maps.Maps()
	}

	{
		LinkedList.LinkedList()
		LinkedList.DesignLinkedList()
	}

	{
		StacksnQueues.StacksAndQueues()
	}

	{
		Trees.Trees()
	}

}
