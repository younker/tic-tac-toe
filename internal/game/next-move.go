package game

import (
    "strings"
)

// The value assigned to each player; synonymous to `X` and `O`. Note that while
// these values are (signed) ints, they do not have any impact on outcome
// calculations
const (
    Empty     int = 0
    PlayerOne int = 1
    PlayerTwo int = 2
)

// Used for calculating the desired outcome with the goal being a maxPlayer
// victory or, worst case, draw.
const (
    maxPlayerWins int = 10
    Draw          int = 0
    minPlayerWins int = -10
)

type move struct {
    Player int
    Index  int
    Score  int
}

var indent string

// Counter tracking the total number of potential outcomes
var outcomes uint16

func GetNextMove(board [9]int, maxPlayer int, minPlayer, currPlayer int) move {
    outcomes++

    if HasPlayerWon(board, minPlayer) {
        return move{Score: minPlayerWins}
    }

    if HasPlayerWon(board, maxPlayer) {
        return move{Score: maxPlayerWins}
    }

    emptyCells := indicesOf(board, func(cell int) bool {
        return cell == Empty
    })
    if len(emptyCells) == 0 {
        return move{Score: Draw}
    }

    var moves []move
    for _, emptyCell := range emptyCells {
        m := move{Index: emptyCell, Player: currPlayer}
        orig := board[emptyCell]
        board[emptyCell] = currPlayer
        indent += "  "

        if currPlayer == maxPlayer {
            worstMove := GetNextMove(board, maxPlayer, minPlayer, minPlayer)
            m.Score = worstMove.Score
        } else {
            bestMove := GetNextMove(board, maxPlayer, minPlayer, maxPlayer)
            m.Score = bestMove.Score
        }

        indent = strings.TrimSuffix(indent, "  ")
        board[emptyCell] = orig
        moves = append(moves, m)
    }

    // The player we want to win is known as the the "maximizing player". Below,
    // when the maxPlayer is the current player, our next move needs to have the
    // highest chance (score) of success. Conversely, if we are calculating the
    // next move for our opponent, or the minimzing player, then we will want to
    // select the move with the lowest change (score) of success.
    var nextMove move
    if currPlayer == maxPlayer {
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
//   {2, 2, 1, 1, 2, 2, 0, 0, 1}
//   {1, 1, -1, -1, 1, 1, 0, 0, -1}
// This algorithm will select the following end-of-game sequence:
//   {2,5}, {1,7}, {2,8}
//   {2,5}, {1,8}, {2,7}
// Both result in the maxPlayer win but really, the only move we needed to make
// was {2,7}. To solve this we could factor in the number of steps required to
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
