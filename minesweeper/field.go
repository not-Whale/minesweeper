package minesweeper

import (
	"errors"
	"log"
	"math"
	"math/rand"
	"time"
)

type Game struct {
	field         [][]Cell
	width, height int
	level         int
	bombs         int
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

func (game Game) isWin() (ok bool) {
	if game.markedCells == game.bombs &&
		game.openedCells == game.width*game.height-game.bombs &&
		game.markBalance == 0 {
		ok = true
	}
	return
}

func (game Game) unmarkCell(cell Cell) {
	if !cell.isBomb {
		game.markBalance++
	}

	cell.isMarked = false
	game.markedCells--
}

func (game Game) markCell(cell Cell) (bool, error) {
	if cell.isMarked {
		return false, errors.New("ЯЧЕЙКА УЖЕ ОТМЕЧЕНА")
	}

	if !cell.isBomb {
		game.markBalance--
	}

	cell.isMarked = true
	game.markedCells++

	if game.isWin() {
		return true, errors.New("ПОБЕДА")
	}

	return true, nil
}

func (game Game) calcBombsAround(cell Cell) (count int) {
	for i := int(math.Min(float64(cell.y-1), 0)); i < int(math.Min(float64(cell.y+1), float64(game.height))); i++ {
		for j := int(math.Min(float64(cell.x-1), 0)); j < int(math.Max(float64(cell.x+1), float64(game.width))); j++ {
			if game.field[i][j].isBomb {
				count++
			}
		}
	}
	return
}

func (game Game) openCell(cell Cell) (bool, error) {
	if cell.isBomb {
		return false, errors.New("БОМБА")
	}

	if cell.bombsAround = game.calcBombsAround(cell); cell.bombsAround == 0 {
		for i := int(math.Min(float64(cell.y-1), 0)); i < int(math.Min(float64(cell.y+1), float64(game.height))); i++ {
			for j := int(math.Min(float64(cell.x-1), 0)); j < int(math.Max(float64(cell.x+1), float64(game.width))); j++ {
				if i != cell.y || j != cell.x {
					_, _ = game.openCell(game.field[i][j])
				}
			}
		}
	}

	cell.isOpened = true
	game.openedCells++

	game.isWin()
	return true, nil
}

func (game Game) initField(level int) (bool, error) {
	switch game.level = level; game.level {
	case 1:
		game.width = EasyWidth
		game.height = EasyHeight
		game.bombs = EasyBombs
	case 2:
		game.width = MediumWidth
		game.height = MediumHeight
		game.bombs = MediumBombs
	case 3:
		game.width = HardWidth
		game.height = HardHeight
		game.bombs = HardBombs
	default:
		return false, errors.New("UNKNOWN LEVEL")
	}
	return true, nil
}

func (game Game) initCells() {
	game.field = make([][]Cell, game.height)
	for i := 0; i < game.height; i++ {
		game.field[i] = make([]Cell, game.width)
		for j := 0; j < game.width; j++ {
			game.field[i][j].x = i
			game.field[i][j].y = j
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

func (game Game) generateBombs() {
	xs, ys := getBombsLocation(game.width, game.width, game.bombs)

	for i := 0; i < game.bombs; i++ {
		game.field[xs[i]][ys[i]].isBomb = true
	}
}

func (game Game) init(level int) {
	_, err := game.initField(level)
	if err != nil {
		log.Fatal(err)
	}

	game.initCells()
	game.generateBombs()
}
