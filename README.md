# Tic Tac Toe

A simple go program meant to take in the `current player` and state of the board. The board state is represented by `0` (empty cell), `1` (player 1) & `2`s (player 2). Based on the board state, the program will return the next next move for the current player.

The decision is made using the [minimax](https://en.wikipedia.org/wiki/Minimax) algorithm (currently without alpha beta pruning) which is designed to minimize negative outcome while maximizing opportunity for success. In short, the best you can do against this program is tie.

For this reason, we refer to the current player as the `maxPlayer` (aka the player we want to win) and the opponent as the `minPlayer`.

# Binaries

There are currently 2 binaries:
1. `cli`: has basic command line parsing allowing for use like so:

```bash
$ go run cmd/cli/main.go 2 1 1 2 2 2 1 1 0 0
[1 1 2 2 2 1 1 2 0]
```

2. `server`: a compressed binary deployed to aws lambda (using [scripts/deploy.sh](src/github.com/younker/tic-tac-toe/scripts/deploy.sh)) where it is available via HTTP:

```bash
$ curl https://vpppfv00l4.execute-api.us-east-1.amazonaws.com/prod/get-move -X POST -d '{"player": 2, board":[0, 0, 1, 1, 2, 2, 0, 0, 1]}'
HTTP/2 200

{
  "player": 2,
  "board": [2, 0, 1, 1, 2, 2, 0, 0, 1]
}
```

# Board

This current definition for `board` is simply an array of 9 integers where the values are as follows:
- `0`: represents and empty cell, available for either player to claim
- `1`: represents all cells owned by player 1
- `2`: represents all cells owned by player 2

For example, in a game where player 1 is assigned to `x` and player 2 is assigned `o`, a board which looks like this:
```
   |   | x
---|---|---
 x | o | o
---|---|---
   |   | x
```

Is represented with `[0, 0, 1, 1, 2, 2, 0, 0, 1]`. If player 2 has the next move, you could get their move by running the following:

```
$ go run cmd/cli/main.go 2 0 0 1 1 2 2 0 0 1
[2 0 1 1 2 2 0 0 1]
```

The result would represent the new state of the board where player 2's latest move can be seen in the upper-left cell:
```
 o |   | x
---|---|---
 x | o | o
---|---|---
   |   | x
```

# Development

For posterity (since this is my first app with golang and have thus forgotten about this several times); local documentation, ideal when working on the ferry!

```
$ godoc -http=:6060
$ open http://localhost:6060
```

# TODO
- [ ] flag for verbose mode (should be quite by default, esp. in lambda)
- [ ] better organize `/internal` such that discrete functions are isolated (and more testable... which brings me to)
- [ ] add some tests (of course)
- [ ] add alpha beta pruning to reduce the number of outcomes explored
- [ ] needs more robust arguments for `cmd/cli`. Ideally we would take a player (whose turn it is) so bots could play each other
- [-] provide more friendly usage. Currently this requires the consumer to know who `0`, `1` and `2` are
- [ ] update to work with API Gateway via [websockets](https://www.youtube.com/watch?v=3SCdzzD0PdQ)
- [ ] `internal/next-move#pick(Best|Worst)Move` could be cleaner if you passed in a compare function. Here is a chicken scratch to jog your memory:

```go
func pickBestMove(moves []move, compare func(a move, b move) move) move {
    return pickMove(moves, func(a, b move) bool {
        return a.score > b.score
    })
}
```
