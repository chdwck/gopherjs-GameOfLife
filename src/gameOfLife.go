package main

import (
	"math/rand"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

// Cell is one of our cells
type Cell struct {
	X     int
	Y     int
	Alive bool
}

const boardSize = 50
const cellWH = 1000 / boardSize

func main() {
	canvas := js.Global.Get("document").Call("getElementById", "canvas")
	context := canvas.Call("getContext", "2d")
	board := initBoard(boardSize, boardSize, false)
	for {
		gameOfLife(context, &board)
	}
}

func gameOfLife(context *js.Object, board *[][]Cell) {
	renderBoard(context, board)
	getNextIteration(board)
	time.Sleep(100 * time.Millisecond)
}

func renderBoard(context *js.Object, board *[][]Cell) {
	context.Call("clearRect", 0, 0, 1000, 1000)
	for _, row := range *board {
		for _, cell := range row {
			drawCell(context, cell)
		}
	}
}

func drawCell(context *js.Object, cell Cell) {
	if cell.Alive {
		context.Set("fillStyle", "#FFFFFF")
	} else {
		context.Set("fillStyle", "#000000")
	}
	context.Call("fillRect", cell.X*cellWH, cell.Y*cellWH, cellWH, cellWH)
}

func getNextIteration(board *[][]Cell) {
	for y, row := range *board {
		for x := range row {
			(*board)[y][x].Alive = checkAlive(board, x, y)
		}
	}
}

func initBoard(x int, y int, dead bool) [][]Cell {
	board := make([][]Cell, y)
	for i := range board {
		for j := 0; j < x; j++ {
			cell := Cell{}
			cell.Y = i
			cell.X = j
			if dead {
				cell.Alive = false
			} else {
				cell.Alive = rand.Intn(15) == 2
			}
			board[i] = append(board[i], cell)
		}
	}
	return board
}

// Any live cell with fewer than two live neighbors dies, as if by under population.
// Any live cell with two or three live neighbors lives on to the next generation.
// Any live cell with more than three live neighbors dies, as if by overpopulation.
// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

func checkAlive(board *[][]Cell, x int, y int) bool {
	neighborCount := 0
	startY := y - 1
	startX := x - 1
	yMax := len(*board)
	xMax := len((*board)[0])
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if startY+i < 0 || startX+j < 0 || startY+i >= yMax || startX+j >= xMax || (startY+i == y && startX+j == x) {
				continue
			}
			if (*board)[startY+i][startX+j].Alive {
				neighborCount++
			}
		}
	}
	if !(*board)[y][x].Alive {
		return neighborCount == 3
	}
	return neighborCount >= 2 && neighborCount <= 3
}
