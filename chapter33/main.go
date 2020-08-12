package main

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"sync"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(3)
	stop  = command(4)
)

// MarsGrid представляет сетку какой-то поверхности
// Марса. Может использовать конкурентно другой
// горутиной.
type MarsGrid struct {
	mu     sync.Mutex
	bounds image.Rectangle
	cells  [][]cell
}

type cell struct {
	lifesig int
	occupier *Occupier
}

// Occupier представляет занятую клетку сетки
// Может использоваться конкурентно другой сеткой
type Occupier struct {
	grid *MarsGrid
	pos  image.Point
}

func (o *Occupier) Pos() image.Point {
	return o.pos
}

// Move перемещает occupier на другую клетку сетки
// Сообщает, было ли перемещение успешным
// Может не получиться, если пытается выйти за пределы
// сетки или потому что пытается перместиться в клетку, что
// уже занята. Если проваливается, occupier остается на прежнем месте
func (o *Occupier) Move(p image.Point) bool {
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()
	newCell := o.grid.cell(p)
	if newCell == nil || newCell.occupier != nil {
		return false
	}
	//снимаем с текущей
	o.grid.cell(o.pos).occupier = nil
	//устанавливаем на новую
	newCell.occupier = o
	o.pos = p
	return true
}

// Occupy занимает ячейку в данной точке сетки. Он
// возвращает nil, если точка уже занята или точка
// за пределами сетки. В противном случае возвращается значение, что
// можно переместить в другое место сетки.
func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	g.mu.Lock()
	defer g.mu.Unlock()
	cell := g.cell(p)
	if cell == nil || cell.occupier != nil {
		return nil
	}
	cell.occupier = &Occupier{
		grid: g,
		pos:  p,
	}
	return cell.occupier
}

// возвращает клетку по ее координатам
func (g *MarsGrid) cell(p image.Point) *cell {
	if !p.In(g.bounds) {
		return nil
	}
	return &g.cells[p.Y][p.X]
}

func (g *MarsGrid) Size() image.Point {
	return g.bounds.Max
}

func NewMarsGrid(size image.Point) *MarsGrid {
	grid := &MarsGrid{
		bounds: image.Rectangle{
			Max: size,
		},
		cells: make([][]cell, size.Y),
	}
	for y := range grid.cells {
		grid.cells[y] = make([]cell, size.X)
		for x := range grid.cells[y] {
            cell := &grid.cells[y][x]
            cell.lifesig = rand.Intn(1000)
        }
	}
	return grid
}

type RoverDriver struct {
	name     string
	commandc chan command
	occupier *Occupier
}

func startDriver(name string, grid *MarsGrid) *RoverDriver {
	var occupier *Occupier
	for occupier == nil {
		startPoint := image.Point{
			X: rand.Intn(grid.Size().X),
			Y: rand.Intn(grid.Size().Y),
		}
		occupier = grid.Occupy(startPoint)
	}
	return NewRoverDriver(name, occupier)
}

func NewRoverDriver(name string, occupier *Occupier) *RoverDriver {
	r := &RoverDriver{
		name:     name,
		commandc: make(chan command),
		occupier: occupier,
	}
	go r.drive()
	return r
}

// drive ответственен за вождение марсохода. Ожидается
// что он начнется в горутине.
func (r *RoverDriver) drive() {
	log.Printf("%s начальная позиция %v", r.name, r.occupier.Pos())
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc: // Ждет новых команд на командном канале
			switch c {
			case right: // поворот направо
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left: // поворот налево
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			case start: // начало движения
				direction = image.Point{X: 1, Y: 0}
				log.Printf("начало движения %v", direction)
			case stop:
				direction = image.Point{X: 0, Y: 0}
				log.Printf("остановка %v", direction)
			}
			log.Printf("%s новое направление %v", r.name, direction)

		case <-nextMove:
			nextMove = time.After(updateInterval)
			newPos := r.occupier.Pos().Add(direction)
			if r.occupier.Move(newPos) {
				log.Printf("%s перемещение на %v", r.name, newPos)
				break
			}
			log.Printf("%s заблокирован при попытке перемещения из %v в %v", r.name, r.occupier.Pos(), newPos)
			// Случайно выбирается одно из других случайных направлений
			// Далее мы попробуем передвинуться в новое направление
			dir := rand.Intn(3) + 1
			for i := 0; i < dir; i++ {
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			}
			log.Printf("%s новое случайное направление %v", r.name, direction)
		}
	}
}

// Left поворачивает марсоход налево (90° против часовой стрелки).
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// Right поворачивает марсоход направо (90° по часовой стрелке).
func (r *RoverDriver) Right() {
	r.commandc <- right
}

func (r *RoverDriver) Start() {
	r.commandc <- start
}

func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

func main() {
	size := image.Point{X: 20, Y: 20}
	grid := NewMarsGrid(size)

	rover := make([]*RoverDriver, 8)
	for i := range rover {
		rover[i] = startDriver(fmt.Sprint("Марсоход ", i), grid)
	}
	time.Sleep(6 * time.Second)
}