package main

import (
    "fmt"
    "log"
    "os"
    "strconv"
    "time"

    "github.com/younker/tic-tac-toe/internal/game"
)

func main() {
    start := time.Now()

    board := parseArgs(os.Args[1:])

    // TODO: [jason@] introduce more sophisticated args parsing so that we can
    // have flags to modify output (ie debug). For now, this will work.
    // debug := false

    fmt.Printf("Initial Board State: %v\n", board)
    nextMove := game.GetNextMove(board, game.Bot)
    board[nextMove.Index] = nextMove.Player
    runtime := time.Since(start)
    fmt.Printf("Runtime: %v\n", runtime)
    fmt.Printf("Move {player, index, score}: %v\n", nextMove)
    fmt.Printf("Final Board State: %v\n", board)
}

func parseArgs(input []string) [9]int {
    var board [9]int
    for i, cell := range input {
        n, err := strconv.Atoi(cell)
        if err != nil {
            log.Fatalf("cannot parse input: %s at position %d", cell, i)
        }

        board[i] = n
    }

    return board
}
