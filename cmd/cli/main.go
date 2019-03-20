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

    maxPlayer, minPlayer := getPlayers(os.Args[1])
    board := getBoard(os.Args[2:])

    nextMove := game.GetNextMove(board, maxPlayer, minPlayer, maxPlayer)
    board[nextMove.Index] = nextMove.Player
    runtime := time.Since(start)
    fmt.Printf("Runtime: %v\n", runtime)
    fmt.Printf("Move {player, index, score}: %v\n", nextMove)
    fmt.Printf("Final Board State: %v\n", board)
}

func getPlayers(input string) (int, int) {
    currPlayer, err := strconv.Atoi(input)
    if err != nil {
        log.Fatalf("cannot parse player: %s at position %d", os.Args[1], 1)
    }

    opponent := game.PlayerTwo
    if currPlayer == opponent {
        opponent = game.PlayerOne
    }

    return currPlayer, opponent
}

func getBoard(input []string) [9]int {
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
