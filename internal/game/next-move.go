package game

import (
    "strings"
)

// The value assigned to each player; synonymous to `X` and `O`. Note that while
// these values are (signed) ints, they do not have any impact on outcome
// calculations
const (
    Empty    int = 0
    Opponent int = 1
    Bot      int = 2
)

// Used for calculating desired outcomes (aka best move) with the presumptive
// goal being a Bot victory or, worst case, draw.
const (
    BotWins      int = 10
    Draw         int = 0
    OpponentWins int = -10
)

type move struct {
    Player int
    Index  int
    Score  int
}

var indent string

// Counter tracking the total number of potential outcomes
var outcomes uint16

func GetNextMove(board [9]int, player int) move {
    outcomes++

    if HasPlayerWon(board, Opponent) {
        return move{Score: OpponentWins}
    }

    if HasPlayerWon(board, Bot) {
        return move{Score: BotWins}
    }

    emptyCells := indicesOf(board, func(cell int) bool {
        return cell != Opponent && cell != Bot
    })
    if len(emptyCells) == 0 {
        return move{Score: Draw}
    }

    var moves []move
    for _, emptyCell := range emptyCells {
        m := move{Index: emptyCell, Player: player}
        orig := board[emptyCell]
        board[emptyCell] = player
        indent += "  "

        if player == Bot {
            worstMove := GetNextMove(board, Opponent)
            m.Score = worstMove.Score
        } else {
            bestMove := GetNextMove(board, Bot)
            m.Score = bestMove.Score
        }

        indent = strings.TrimSuffix(indent, "  ")
        board[emptyCell] = orig
        moves = append(moves, m)
    }

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

    return nextMove
}

func indicesOf(board [9]int, fn func(int) bool) []int {
    var spaces []int
    for i, v := range board {
        if fn(v) {
            spaces = append(spaces, i)
        }
    }
    return spaces
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
        if m.Score > highScore {
            highScore = m.Score
            bestMove = m
        }
    }
    return bestMove
}

func pickWorstMove(moves []move) move {
    var worstMove move
    lowScore := 100
    for _, m := range moves {
        if m.Score < lowScore {
            lowScore = m.Score
            worstMove = m
        }
    }
    return worstMove
}
