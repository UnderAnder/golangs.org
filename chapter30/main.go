package main

import (
	"errors"
	"fmt"
	"strings"
)

const rows, columns = 9, 9

var (
	ErrBounds   = errors.New("outside of grid boundaries")
	ErrDigit    = errors.New("invalid digit")
	RowRules    = errors.New("Повторение в строке")
	ColumnRules = errors.New("Повторение в столбце")
	ZoneRules   = errors.New("Повторение числа в зоне")
	FixedRules  = errors.New("Это число нельзя менять")
)

// Cell является зоной Судоку
type Cell struct {
	digit int8
	fixed bool
}

// Grid является сеткой Судоку
type Grid [rows][columns]Cell

// SudokuError массив ошибок
type SudokuError []error

func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if g.inRow(row, digit) {
		errs = append(errs, RowRules)
	}
	if g.inColumn(column, digit) {
		errs = append(errs, ColumnRules)
	}
	if g.inZone(row, column, digit) {
		errs = append(errs, ZoneRules)
	}
	if g.isFixed(row, column) {
		errs = append(errs, FixedRules)
	}
	if len(errs) > 0 {
		return errs
	}

	g[row][column].digit = digit
	return nil
}

func (g *Grid) inZone(row, column int, digit int8) bool {
	startRow, startColumn := row/3*3, column/3*3
	for r := startRow; r < startRow+3; r++ {
		for c := startColumn; c < startColumn+3; c++ {
			if g[r][c].digit == digit {
				return true
			}
		}
	}
	return false
}

func (g Grid) isFixed(row, column int) bool {
	return g[row][column].fixed
}

func (g *Grid) inColumn(column int, digit int8) bool {
	for r := 0; r < rows; r++ {
		if g[r][column].digit == digit {
			return true
		}
	}
	return false
}

func (g *Grid) inRow(row int, digit int8) bool {
	for c := 0; c < columns; c++ {
		if g[row][c].digit == digit {
			return true
		}
	}
	return false
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func validDigit(d int8) bool {
	return d >= 1 && d <= 9
}

// Error возвращает одну или несколько ошибок через запятые.
func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error()) // Конвертирует ошибки в строки
	}
	return strings.Join(s, ", ")
}

func NewSudoku(d [rows][columns]int8) *Grid {
	var grid Grid
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			digit := d[i][j]
			if digit != 0 {
				grid[i][j].digit = digit
				grid[i][j].fixed = true
			}
		}
	}
	return &grid
}

func main() {
	s := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	err := s.Set(0, 2, 4)
	if err != nil {
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
	}

	for _, row := range s {
		fmt.Println(row)
	}
}
