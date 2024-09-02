package main

import (
	"container/heap"
	"fmt"
)




func main() {
	// Read input from file
	lines := readInput("input.txt")
	// just a big pile of lines, I want to group them

	fmt.Printf("Lines: %v+\n", lines)

	// Build a MaxHeap, my money's on the second question being "Which elf has the most calories?"
	foo := make(MaxHeap, len(lines))

	for i, cals := range lines {
		sum := 0
		for _, c := range cals {
			sum += c
		}

		foo[i] = &Elf{
			id: i,
			calories: cals,
			total: sum,
			index: i,
		}

		i++
	}

	// Now we have a MaxHeap of elves, let's see who has the most calories
	heap.Init(&foo)

	// Pop the top elf
	topElf := heap.Pop(&foo).(*Elf)

	fmt.Printf("Elf %d has the most calories: %v\n", topElf.id + 1, topElf.total)

	// Ah, I was wrong, the second part is "add the top 3 together" - ezpz, just pop 3 times

	secondElf := heap.Pop(&foo).(*Elf)
	thirdElf := heap.Pop(&foo).(*Elf)

	fmt.Printf("The top 3 elves have %d calories\n", topElf.total + secondElf.total + thirdElf.total)

}
