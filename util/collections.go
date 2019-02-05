package util

// filter board values for
func IndicesOf(board [9]int, fn func(int) bool) []int {
    var spaces []int
    for i, v := range board {
        if fn(v) {
            spaces = append(spaces, i)
        }
    }
    return spaces
}
