package nqueen

import (
	"math/rand"
)

// N-Queen structure is a board of integers where the index represents
// the column of the queen in the board and the content represents the
// line of the queen.
type queen struct {
	board []int
}

// Make a new N-Queen object with a board of mixed values after
// a sequential initialization.
func Make(size int) queen {
	q := queen{board: make([]int, size)}

	for i := 0; i < size; i++ {
		q.board[i] = i
	}

	q.mixBoard()

	return q
}

// duplicates replicates a N-Queen object by creating a new one
// and copying its contents to the new one.
func (q *queen) duplicate() queen {
	newBoard := make([]int, len(q.board))

	for i := 0; i < len(q.board); i++ {
		newBoard[i] = q.board[i]
	}

	newQueen := queen{board: newBoard}
	return newQueen
}

// randInt generates a random integer from 0 to the size of the board
func (q queen) randInt() int {
	return rand.Intn(len(q.board))
}

// swapTwo swaps two random queens from the N-Queens board
func (q *queen) swapTwo() {
	first := q.randInt()
	second := q.randInt()

	q.board[first], q.board[second] = q.board[second], q.board[first]
}

// mixBoard mixes the N-Queens board by swapping two random queens as many times
// as the size of the board.
func (q *queen) mixBoard() {
	for i := 0; i < len(q.board); i++ {
		q.swapTwo()
	}
}

// areThreats checkes if two given queens threat each other.
func (q *queen) areThreats(first int, second int) bool {
	return q.board[first]-first == q.board[second]-second ||
		q.board[first]+first == q.board[second]+second ||
		q.board[first] == q.board[second]
}

// Heuristic function returns the number of threats in a board of N-Queens.
func (q *queen) Heuristic() int {
	threats := 0

	for i := 0; i < len(q.board); i++ {
		for j := i + 1; j < len(q.board); j++ {
			if q.areThreats(i, j) {
				threats++
			}
		}
	}
	return threats
}

// Objective function checks if a given board is a solution to the problem,
// that is, if its heuristic is 0.
func (q *queen) Objective() bool {
	return q.Heuristic() == 0
}

// Sucessor generates a possible list of successors and selects the first
// one found where its heuristic is smaller or equal than the current one.
func (q *queen) Successor() *queen {
	var (
		listSize         = len(q.board) * 5
		successors       = make([]queen, listSize)
		currentHeuristic = q.Heuristic()
	)

	for i := 0; i < listSize; i++ {
		newSuccessor := q.duplicate()
		newSuccessor.swapTwo()
		successors[i] = newSuccessor
	}

	for _, s := range successors {
		if s.Heuristic() <= currentHeuristic {
			return &s
		}
	}

	return nil
}

// Solve function is the function that solves the problem. This is an
// implementation of the hill-climbing algorithm with restarts.
func (q *queen) Solve() queen {
	current := Make(len(q.board))

	for {
		for i := 0; i < len(q.board)*3; i++ {
			successor := current.Successor()

			if successor != nil {
				current = *successor
				break
			} else {
				continue
			}
		}

		if current.Objective() {
			return current
		}
	}
}
