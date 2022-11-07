package minesweeper

import (
	"errors"
	"log"
	"math"
	"math/rand"
	"time"
)

type Field struct {
	width, height int
	level         int
	bombs         int
	field         [][]Cell
	openedCells   int
	markedCells   int
	markBalance   int
}

type Cell struct {
	x, y        int
	bombsAround int
	isBomb      bool
	isOpened    bool
	isMarked    bool
}

func (f Field) checkVin() (ok bool) {
	if f.markedCells == f.bombs &&
		f.openedCells == f.width*f.height-f.bombs &&
		f.markBalance == 0 {
		ok = true
	}
	return
}

func (f Field) unmarkCell(cell Cell) {
	if !cell.isBomb {
		f.markBalance++
	}

	cell.isMarked = false
	f.markedCells--
}

func (f Field) markCell(cell Cell) {
	if !cell.isBomb {
		f.markBalance--
	}

	cell.isMarked = true
	f.markedCells++

	f.checkVin()
}

func (f Field) calcBombsAround(cell Cell) (count int) {
	for i := int(math.Min(float64(cell.y-1), 0)); i < int(math.Min(float64(cell.y+1), float64(f.height))); i++ {
		for j := int(math.Min(float64(cell.x-1), 0)); j < int(math.Max(float64(cell.x+1), float64(f.width))); j++ {
			if f.field[i][j].isBomb {
				count++
			}
		}
	}
	return
}

func (f Field) openCell(cell Cell) (bool, error) {
	if cell.isBomb {
		return false, errors.New("BOMB")
	}

	if cell.bombsAround = f.calcBombsAround(cell); cell.bombsAround == 0 {
		for i := int(math.Min(float64(cell.y-1), 0)); i < int(math.Min(float64(cell.y+1), float64(f.height))); i++ {
			for j := int(math.Min(float64(cell.x-1), 0)); j < int(math.Max(float64(cell.x+1), float64(f.width))); j++ {
				if i != cell.y || j != cell.x {
					_, _ = f.openCell(f.field[i][j])
				}
			}
		}
	}

	cell.isOpened = true
	f.openedCells++

	f.checkVin()
	return true, nil
}

func (f Field) initField(level int) (bool, error) {
	switch f.level = level; f.level {
	case 1:
		f.width = EasyWidth
		f.height = EasyHeight
		f.bombs = EasyBombs
	case 2:
		f.width = MediumWidth
		f.height = MediumHeight
		f.bombs = MediumBombs
	case 3:
		f.width = HardWidth
		f.height = HardHeight
		f.bombs = HardBombs
	default:
		return false, errors.New("UNKNOWN LEVEL")
	}
	return true, nil
}

func (f Field) initCells() {
	f.field = make([][]Cell, f.height)
	for i := 0; i < f.height; i++ {
		f.field[i] = make([]Cell, f.width)
		for j := 0; j < f.width; j++ {
			f.field[i][j].x = i
			f.field[i][j].y = j
		}
	}
}

func getBombsLocation(x, y, num int) ([]int, []int) {
	rand.Seed(time.Now().Unix())

	xs, ys := make([]int, num), make([]int, num)
	for i := 0; i < num; i++ {
		xs[i] = rand.Intn(x)
		ys[i] = rand.Intn(y)
	}
	return xs, ys
}

func (f Field) generateBombs() {
	xs, ys := getBombsLocation(f.width, f.width, f.bombs)

	for i := 0; i < f.bombs; i++ {
		f.field[xs[i]][ys[i]].isBomb = true
	}
}

func (f Field) init(level int) {
	_, err := f.initField(level)
	if err != nil {
		log.Fatal(err)
	}

	f.initCells()
	f.generateBombs()
}
