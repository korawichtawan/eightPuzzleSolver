# eightPuzzleSolver
This is a Golang package to solve the 8 puzzle using state space search method.
# Example use case
This is an example of simple 8 puzzle solver web application [here](https://github.com/korawichtawan/eightPuzzleWebApp).
# How to use
import eight-puzzle solver 
```js
import ("github.com/korawichtawan/eightPuzzleSolver")
```
create puzzle board with 3*3 dimensional array.
```js
  board := [3][3]int{
    {1,2,3},
    {4,5,6},
    {0,7,8},
  }
```
call function Solve() from eightPuzzleSolver and pass board as parameter.
```js
  minMove,answer,err := eightPuzzleSolver.Solve(board)
```
# Result
minMove is int => minimum moves to solve the puzzle.<br />
answer is slice of string => list of moves to solve the puzzle("t": "top", "r": "right", "b":"bottom", "l":"left"). <br />
err is error => function Solve() return error when puzzle is invalid or unsolvable. <br />
```js
  fmt.Println(minMove)
  fmt.Println(answer)
  fmt.Println(err)
  //2
  //["r","r"]
  //nil
```
