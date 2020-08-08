package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type safeWriter struct {
	w   io.Writer
	err error // Место для хранения первой ошибки
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return // Пропускает запись, если раньше была ошибка
	}
	_, sw.err = fmt.Fprintln(sw.w, s) // Записывает строку и затем хранить любую ошибку
}

/* func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		f.Close()
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	//f.Close()
	return err
} */

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors, handle them gracefully.")
	sw.writeln("Don't panic.")
	sw.writeln("Make the zero value useful.")
	sw.writeln("The bigger the interface, the weaker the abstraction.")
	sw.writeln("interface{} says nothing.")
	sw.writeln("Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.")
	sw.writeln("Documentation is for users.")
	sw.writeln("A little copying is better than a little dependency.")
	sw.writeln("Clear is better than clever.")
	sw.writeln("Concurrency is not parallelism.")
	sw.writeln("Don't communicate by sharing memory, share memory by communicating.")
	sw.writeln("Channels orchestrate; mutexes serialize.")

	return sw.err // Возвращает ошибку в случае ее возникновения
}

const rows, columns = 9, 9

var (
	ErrBounds = errors.New("outside of grid boundaries")
	ErrDigit  = errors.New("invalid digit")
)

// Grid является сеткой Судоку
type Grid [rows][columns]int8

func (g *Grid) Set(row, column int, digit int8) error {
/* 	if !inBounds(row, column) {
		return ErrBounds
	}
	if !validDigit(digit) {
		return ErrDigit
	} */

	var errs SudokuError
    if !inBounds(row, column) {
        errs = append(errs, ErrBounds)
    }
    if !validDigit(digit) {
        errs = append(errs, ErrDigit)
    }
    if len(errs) > 0 {
        return errs
	}
	
	g[row][column] = digit
	return nil
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

type SudokuError []error

// Error возвращает одну или несколько ошибок через запятые.
func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error()) // Конвертирует ошибки в строки
	}
	return strings.Join(s, ", ")
}

func main() {
	/* 	files, err := ioutil.ReadDir(".")
	   	if err != nil {
	   		fmt.Println(err)
	   		os.Exit(1)
	   	}

	   	for _, file := range files {
	   		fmt.Println(file.Name())
	   	}

	   	errr := proverbs("proverbs.txt")
	   	if errr != nil {
	   		fmt.Println(err)
	   		os.Exit(1)
	   	} */
	///////////////////////////

	var g Grid
	err := g.Set(10, 0, 11)
	if err != nil {
/* 		switch err {
		case ErrBounds, ErrDigit:
			fmt.Println("Les erreurs de paramètres hors limites.")
		default:
			fmt.Println(err)
		}*/

		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}

		os.Exit(1) 
	}
}
