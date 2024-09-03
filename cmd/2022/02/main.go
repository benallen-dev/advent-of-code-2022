package main

import (
	"bufio"
	"os"
	"fmt"
)

func waitForEnter() {
    fmt.Print("Press Enter to continue...")
    reader := bufio.NewReader(os.Stdin)
    _, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("\nError reading input:", err)
    }
}

func main() {
	rounds, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	total := 0
	for i, round := range rounds {
		fmt.Printf("%04d: %d + %d = %d\n", i, total, round.Score(), total + round.Score())
		total += round.Score()
		//waitForEnter()
	}

	// Part 01, total score if everything happens as described
	// No doubt part 2 will be "oh shit they're trying to trick you" or something
	fmt.Println(total)

	// Nope, we bamboozled ourselves!
	// X = lose
	// Y = draw
	// Z = win

	// The whole round datastructure is wrong, and I could have known because
	// why use ABC XYZ when they refer to the same thing?

}
