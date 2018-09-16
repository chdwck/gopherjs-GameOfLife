package main

import (
	"reflect"
	"testing"
)

// Any live cell with fewer than two live neighbors dies, as if by under population.
// Any live cell with two or three live neighbors lives on to the next generation.
// Any live cell with more than three live neighbors dies, as if by overpopulation.
// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
func TestCheckAlive(t *testing.T) {
	b := initBoard(3, 3, true)
	board := &b
	(*board)[1][1].Alive = true
	if checkAlive(board, 1, 1) {
		t.Error("Expcted no surivors")
	}
	(*board)[1][2].Alive = true
	(*board)[1][0].Alive = true
	if !checkAlive(board, 1, 1) {
		t.Error("Expcted surivors")
	}
	(*board)[0][1].Alive = true
	(*board)[2][1].Alive = true
	if checkAlive(board, 1, 1) {
		t.Error("Expcted no surivors")
	}
}

func TestInitBoard(t *testing.T) {
	board := initBoard(500, 500, false)
	if len(board) != 500 {
		t.Errorf("Board was not properly initialized")
	}
	for i, row := range board {
		if len(row) != 500 {
			t.Errorf("Board row %d not properly initialized", i)
		}
		for _, cell := range row {
			cellType := reflect.TypeOf(cell).Kind()
			if cellType != reflect.Struct {
				t.Errorf("Expected struct got: %q", cellType)
			}
			aliveType := reflect.TypeOf(cell.Alive).Kind()
			if aliveType != reflect.Bool {
				t.Errorf("Expected bool got: %q", aliveType)
			}
			xType := reflect.TypeOf(cell.X).Kind()
			if xType != reflect.Int {
				t.Errorf("Expected number got: %q", xType)
			}
			yType := reflect.TypeOf(cell.Y).Kind()
			if yType != reflect.Int {
				t.Errorf("Expected number got: %q", yType)
			}

		}
	}
}
