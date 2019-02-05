package main

import (
    "fmt"
    "strings"
    "time"

    "github.com/younker/tic-tac-toe/util"
)

// The value assigned to each player; synonymous to `X` and `O`. Note that while
// these values are (signed) ints, they do not have any impact on outcome
// calculations
const (
    Bot      int = 1
    Empty    int = 0
    Opponent int = -1
)

// Used for calculating desired outcomes (aka best move) with the presumptive
// goal being a Bot victory or, worst case, draw.
const (
    BotWins      int = 10
    Draw         int = 0
    OpponentWins int = -10
)

// Counter tracking the total number of potential outcomes
var outcomes uint16

type move struct {
    player int
    index  int
    score  int
}

func main() {
    start := time.Now()

    // Placeholder: partially completed game with limited number of moves used
    // for initial development. Board visual:
    // x x o
    // o x
    board := [9]int{1, 1, -1, -1, 1, 1, 0, 0, -1}

    fmt.Printf("Initial Board State: %v\n", board)
    nextMove := getNextMove(board, Bot)
    runtime := time.Since(start)
    fmt.Printf("Runtime: %v\n", runtime)
    fmt.Printf("Possible Outcomes: %v\n", outcomes)
    fmt.Printf("Move {player, index, score}: %v\n", nextMove)
}

var indent string

func getNextMove(board [9]int, player int) move {
    outcomes++

    if util.HasPlayerWon(board, Opponent) {
        fmt.Printf("%v! Game Over. Opponent wins\n", indent)
        return move{score: OpponentWins}
    }

    if util.HasPlayerWon(board, Bot) {
        fmt.Printf("%v! Game Over. Bot wins\n", indent)
        return move{score: BotWins}
    }

    emptyCells := util.IndicesOf(board, func(cell int) bool {
        return cell != Opponent && cell != Bot
    })
    if len(emptyCells) == 0 {
        fmt.Printf("%v! Game Over. Draw\n", indent)
        return move{score: Draw}
    }
    fmt.Printf("%v  emptyCells: %v\n", indent, emptyCells)

    var moves []move
    for _, emptyCell := range emptyCells {
        m := move{index: emptyCell, player: player}
        orig := board[emptyCell]
        board[emptyCell] = player
        indent += "  "
        fmt.Printf("%v+ move %v to cell %v\n", indent, player, emptyCell)

        if player == Bot {
            worstMove := getNextMove(board, Opponent)
            m.score = worstMove.score
        } else {
            bestMove := getNextMove(board, Bot)
            m.score = bestMove.score
        }

        indent = strings.TrimSuffix(indent, "  ")
        board[emptyCell] = orig
        moves = append(moves, m)
    }
    fmt.Printf("%v  moves: %v\n", indent, moves)

    // The player we want to win is known as the the "maximizing player". In our
    // case, the Bot is our selected champion so our next move needs to have the
    // highest chance (score) of success. Conversely, if we are calculating the
    // next move for our opponent, or the minimzing player, then we will want to
    // select the minimum score.
    var nextMove move
    if player == Bot {
        nextMove = pickBestMove(moves)
    } else {
        nextMove = pickWorstMove(moves)
    }

    fmt.Printf("%v  next move: %v\n", indent, nextMove)
    return nextMove
}

// Of all the possible moves, pick the (first) move with the highest score.
// This is good enough for a first pass but present some problems. For
// example, given the following board:
//   {1, 1, -1, -1, 1, 1, 0, 0, -1}
// This algorithm will select the following end-of-game sequence:
//   {1,5}, {-1,7}, {1,8}
//   {1,5}, {-1,8}, {1,7}
// Both result in the Bot winning but really, the only move we needed to make
// was {1,7}. To solve this we could factor in the number of steps required to
// win scoring quicker wins higher rather than selecting the first "best move".
func pickBestMove(moves []move) move {
    var bestMove move
    highScore := -100
    for _, m := range moves {
        if m.score > highScore {
            highScore = m.score
            bestMove = m
        }
    }
    return bestMove
}

func pickWorstMove(moves []move) move {
    var worstMove move
    lowScore := 100
    for _, m := range moves {
        if m.score < lowScore {
            lowScore = m.score
            worstMove = m
        }
    }
    return worstMove
}
