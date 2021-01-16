package main

import (
	"fmt"
	"os"
	"strconv"
)

type verbose struct {
	solutions   bool
	equations   bool
	oneequation bool
}

// NumbersSolution Struct to hold a solution
type NumbersSolution struct {
	found   bool
	answer  int
	steps   int
	history []op
	eq      string
}

type stats struct {
	operations int
	stored     int
	numbersets int
	beststeps  int
	closesteps int
	iterations int
}

var debug = verbose{solutions: false, equations: false, oneequation: false}
var counter stats

type op struct {
	x int
	y int
	r int
	o string
	u bool
}

var numbersSolution NumbersSolution

func printResults(results map[int]op) {
	// For operation print stuff

	for result, operation := range results {
		fmt.Println("result:", result, "from", operation.x, operation.o, operation.y)
	}
}

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkOperation(results map[int]op, r int, x int, y int, record op, answer int, history []op, numbers []int) {

	if r != x && r != y && r != 0 {
		results[r] = record
		history = append(history, record)

		// fmt.Println("Result Stored:", record.x, record.o, record.y, "=", r)

		if debug.equations {
			// printEq(append(history, record), r, counter.numbersets, found, steps, numbers)
		}

		if r == answer {
			numbersSolution.found = true

			if counter.beststeps == 0 || len(history) < counter.beststeps {
				numbersSolution.history = make([]op, len(history))
				numbersSolution.answer = r
				copy(numbersSolution.history, history)
				numbersSolution.steps = len(numbersSolution.history)
				numbersSolution.eq = getEq(numbersSolution)
				// fmt.Printf("eq: %v best: %v hl: %v, h: %v \n", getEq(numbersSolution), counter.beststeps, len(history), history)
				counter.beststeps = numbersSolution.steps
			}
		} else {
			if !numbersSolution.found {
				if abs(r-answer) < abs(numbersSolution.answer-answer) {
					numbersSolution.history = make([]op, len(history))
					numbersSolution.answer = r
					copy(numbersSolution.history, history)
					numbersSolution.steps = len(numbersSolution.history)
					numbersSolution.eq = getEq(numbersSolution)
					// fmt.Printf("eq: %v best: %v hl: %v, h: %v \n", getEq(numbersSolution), counter.beststeps, len(history), history)
				}
			}
		}
		// Print all found solutions
		if debug.solutions {
			// printEq(history, r, counter.numbersets, found, steps, numbers)
		}
	}
}

func operate(x int, y int, answer int, history []op, numbers []int) map[int]op {
	// Needs to be an integer
	// Needs to be a unique result
	// Can't be negative
	counter.operations++

	// r := make(map[int]struct{})
	results := make(map[int]op)
	var r int
	var opRecord op

	// Addition
	r = x + y
	opRecord = op{x: x, y: y, r: r, o: "+", u: false}
	// results[r] = op{x: x, y: y, o: "+"}
	checkOperation(results, r, x, y, opRecord, answer, history, numbers)

	// Multiplication
	r = x * y
	opRecord = op{x: x, y: y, r: r, o: "*", u: false}
	checkOperation(results, r, x, y, opRecord, answer, history, numbers)

	// Subtraction
	r = x - y
	if r >= 0 {
		opRecord = op{x: x, y: y, r: r, o: "-", u: false}
		checkOperation(results, r, x, y, opRecord, answer, history, numbers)
	}
	r = y - x
	if r >= 0 {
		opRecord = op{x: y, y: x, r: r, o: "-", u: false}
		checkOperation(results, r, x, y, opRecord, answer, history, numbers)
	}
	// Division
	if x%y == 0 {
		r = x / y
		opRecord = op{x: x, y: y, r: r, o: "/", u: false}
		checkOperation(results, r, x, y, opRecord, answer, history, numbers)
	}
	if y%x == 0 {
		r = y / x
		opRecord = op{x: y, y: x, r: r, o: "/", u: false}
		checkOperation(results, r, x, y, opRecord, answer, history, numbers)
	}

	return results
}

func operateOnNumbers(numbers []int, answer int, history []op) {
	counter.numbersets++
	//fmt.Println("operateOnNumbersCalled:", counter.numbersets, numbers)

	// if len(history) == 1 {
	// 	printDebugNumbers(history, numbers, counter.numbersets)
	// }

	// if there are two or more numbers
	if len(numbers) >= 2 {
		results := make(map[int]op)

		// For every order of those numbers
		for i := 0; i < (len(numbers) - 1); i++ {
			// fmt.Println(numbers)

			// Select the first number and operate on each of the others
			for i, number := range numbers {
				if i > 0 {
					results = operate(numbers[0], number, answer, history, numbers)

					// remove the first and other number
					nn := make([]int, len(numbers))
					// newnumbers = append(numbers[:i])
					copy(nn, numbers)
					// fmt.Println(nn)
					// fmt.Println(nn[i])
					// fmt.Println(i)
					// fmt.Println(number)
					nn[i] = nn[len(nn)-1]
					nn = nn[:len(nn)-1] // Truncate slice.
					nn[0] = nn[len(nn)-1]
					nn = nn[:len(nn)-1] // Truncate slice.
					// fmt.Println(nn)
					// for each result,
					for result, operations := range results {
						// adjust history
						nh := make([]op, len(history))
						// newnumbers = append(numbers[:i])
						copy(nh, history)

						// add the result into the new numbers
						// nn = append(nn, result)
						// nh = append(nh, operations)
						// operateOnNumbers(nn, answer, nh)

						operateOnNumbers(append(nn, result), answer, append(nh, operations))
						// fmt.Println("appending result:", result)
						// fmt.Println("going to call with", append(nn, result))
						// Call operateOnNumbers
						// fmt.Println(nn)

					}
					// I think we should now start on number the second number
					// printDebugNumbers(history, numbers, counter.numbersets)
					// os.Exit(33)
				}

			}
			var tmp int
			tmp, numbers = numbers[0], numbers[1:]
			numbers = append(numbers, tmp)
			// fmt.Println(numbers)
		}

		// printResults(results)
	} else {
		// printEq(history, numbers[0], counter.numbersets, false, 0)
	}
}

func getEq(solution NumbersSolution) string {
	var eq string
	// fmt.Println("Length of history:", len(history))
	for _, operation := range solution.history {
		eq = eq + "(" + strconv.Itoa(operation.x) + operation.o + strconv.Itoa(operation.y) + "=" + strconv.Itoa(operation.r) + ")"
	}
	return eq
}

// PrintSolution Print a solution
func printSolution(solution NumbersSolution) {
	if solution.found {
		fmt.Println("Best Solution:", solution.eq, "in", solution.steps, "steps")
	} else {
		fmt.Println("Closest Solution:", solution.eq, "in", solution.steps, "steps")
	}

}
func printDebugNumbers(history []op, numbers []int, numberset int) {
	var eq string
	// fmt.Println("Length of history:", len(history))
	for _, operation := range history {
		eq = eq + "(" + strconv.Itoa(operation.x) + operation.o + strconv.Itoa(operation.y) + "=" + strconv.Itoa(operation.r) + ")"
	}

	fmt.Println("DEBUG:", numberset, "numberlength", len(numbers), "numbers", numbers, "historyl:", len(history), "history", eq)

	// if numberset >= 10 {
	// 	os.Exit(23)
	// }

	// if len(history) == 1 {
	// 	fmt.Println("DEBUG:", numberset, "numberlength", len(numbers), "numbers", numbers, "historyl:", len(history), "history", eq)
	// }

}

func printEq(history []op, solution int, numberset int, found bool, steps int, numbers []int) {
	var eq string
	if found {
		for _, operation := range history {
			if operation.u {
				eq = eq + "(" + strconv.Itoa(operation.x) + operation.o + strconv.Itoa(operation.y) + "=" + strconv.Itoa(operation.r) + ")"
			}
		}
		fmt.Println("Solution Found", numberset, eq, "=", solution, "in", steps, "steps")
	} else {
		for _, operation := range history {
			eq = eq + "(" + strconv.Itoa(operation.x) + operation.o + strconv.Itoa(operation.y) + "=" + strconv.Itoa(operation.r) + ")"
		}
		fmt.Println("Equation", numberset, eq, "=", solution, numbers)
	}

}

// NumbersGame Solve the numbers game
func NumbersGame(numbers []int, answer int) NumbersSolution {
	var history []op
	numbersSolution.found = false
	numbersSolution.answer = 0
	counter = stats{operations: 0, stored: 0, numbersets: 0, beststeps: 0, closesteps: 0, iterations: 0}
	operateOnNumbers(numbers, answer, history)
	return numbersSolution
}

func main() {
	args := os.Args
	numbers := make([]int, 6)

	for i := 1; i <= 6; i++ {
		numbers[i-1], _ = strconv.Atoi(args[i])
	}
	answer, _ := strconv.Atoi(args[7])
	// fmt.Println(numbers, answer)
	solution := NumbersGame(numbers, answer)
	printSolution(solution)
}
