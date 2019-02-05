package util

var winningCombinations = [8][3]int{
    {0, 1, 2}, // top row, straight across
    {0, 3, 6}, // left col, straight down
    {0, 4, 8}, // upper-left, lower-right diagonal
    {1, 4, 7}, // middle col, straight down
    {2, 5, 8}, // right col, straight down
    {2, 4, 6}, // upper-right, lower-left diagonal
    {3, 4, 5}, // middle row, straight across
    {6, 7, 8}, // bottom row, straight across
}

// HasPlayerWon will return true when the provided player occupies all of the
// cells for any of the defined winningCombinations
func HasPlayerWon(board [9]int, player int) bool {
    playerWon := false
    for _, cells := range winningCombinations {
        ownsAllCells := true
        for _, cell := range cells {
            if board[cell] != player {
                ownsAllCells = false
                break
            }
        }

        if ownsAllCells {
            playerWon = true
            break
        }
    }

    return playerWon
}
