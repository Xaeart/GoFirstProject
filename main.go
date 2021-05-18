package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generateProblem generates a n-sized array of integers
// containing values from 0 to n-1, and the duplicated integer
func generateProblem(n int) ([]int, int) {
	if n < 2 {
		panic("n must be >= 2")
	}
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(n, func(i, j int) { out[i], out[j] = out[j], out[i] })

	// Duplicate a number
	idx := rand.Intn(n)
	num := out[idx]
	idx2 := rand.Intn(n)
	for idx == idx2 {
		idx2 = rand.Intn(n)
	}
	out[idx2] = num
	return out, num
}

// Examples
//  Input	        Output
//	[4 2 3 5 4 0]   4
//	[2 5 3 1 1 4]   1
//	[0 4 0 5 1 3]   0
//	[4 0 3 1 2 3]   3

func solveProblem1(p []int) int {

	// ...
	// Implement the solution here
	// ...

	return -1
}

func solveProblem2(p []int) int {
	return -1
}
func solveProblem3(p []int) int {
	return -1
}

func main() {
	for i := 0; i < 10; i++ {
		// Generate a problem
		problem, solution := generateProblem(6)
		fmt.Printf("%v\t", problem)

		// Solve the problem
		s := solveProblem1(problem)
		fmt.Printf("%d", s)

		// Solve the problem 2nd way
		answer2 := solveProblem2(problem)
		fmt.Printf("%d", s)

		// Solve the problem 3rd way.
		answer3 := solveProblem3(problem)
		fmt.Printf("%d", s)

		// Print the result
		if s == solution {
			fmt.Printf("\tCorrect!")
		} else {
			fmt.Printf("\tWRONG!")
		}

		if answer2 == solution {
			fmt.Printf("\tCorrect!")
		} else {
			fmt.Printf("\tWRONG!")
		}

		if answer3 == solution {
			fmt.Printf("\tCorrect!")
		} else {
			fmt.Printf("\tWRONG!")
		}
		fmt.Println()
	}
}
