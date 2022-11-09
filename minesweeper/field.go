package minesweeper

import (
	"fmt"
	"math"
	"math/rand"
	"os"
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

func (game *Game) openAll() {
	for i := 0; i < game.height; i++ {
		for j := 0; j < game.width; j++ {
			if !game.field[i][j].isBomb && !game.field[i][j].isOpened {
				_ = game.openCell(game.field[i][j])
			}
			if game.field[i][j].isBomb && !game.field[i][j].isMarked {
				_ = game.markCell(game.field[i][j])
			}
		}
	}
}

func (game *Game) isWin() (ok bool) {
	if game.markedCells == game.bombs &&
		// game.openedCells == game.width*game.height-game.bombs &&
		game.markBalance == 0 {
		ok = true
	}
	return
}

func (game *Game) unmarkCell(cell Cell) error {
	if cell.y >= game.height || cell.x >= game.width ||
		cell.y < 0 || cell.x < 0 {
		return errOutOfRange{cell.x, cell.y}
	}

	if !cell.isMarked {
		return errCellIsNotMarked{cell.x, cell.y}
	}

	if !cell.isBomb {
		game.markBalance++
	}

	game.field[cell.y][cell.x].isMarked = false
	game.markedCells--

	return nil
}

func (game *Game) markCell(cell Cell) error {
	if cell.y >= game.height || cell.x >= game.width ||
		cell.y < 0 || cell.x < 0 {
		return errOutOfRange{cell.x, cell.y}
	}

	if cell.isMarked {
		return errCellIsMarked{cell.x, cell.y}
	}

	if !cell.isBomb {
		game.markBalance--
	}

	game.field[cell.y][cell.x].isMarked = true
	game.markedCells++

	if game.isWin() {
		return errYouWin{}
	}

	return nil
}

func (game *Game) calcBombsAround(cell Cell) (count int) {
	for i := int(math.Max(float64(cell.y-1), 0)); i <= int(math.Min(float64(cell.y+1), float64(game.height-1))); i++ {
		for j := int(math.Max(float64(cell.x-1), 0)); j <= int(math.Min(float64(cell.x+1), float64(game.width-1))); j++ {
			if game.field[i][j].isBomb {
				count++
			}
		}
	}
	return
}

func (game *Game) openCell(cell Cell) error {
	if cell.y >= game.height || cell.x >= game.width ||
		cell.y < 0 || cell.x < 0 {
		return errOutOfRange{cell.x, cell.y}
	}

	if cell.isBomb {
		return errGameOver{cell.x, cell.y}
	}

	game.openedCells++
	game.field[cell.y][cell.x].isOpened = true
	game.field[cell.y][cell.x].bombsAround = game.calcBombsAround(cell)

	if game.field[cell.y][cell.x].bombsAround == 0 {
		for i := int(math.Max(float64(cell.y-1), 0)); i <= int(math.Min(float64(cell.y+1), float64(game.height-1))); i++ {
			for j := int(math.Max(float64(cell.x-1), 0)); j <= int(math.Min(float64(cell.x+1), float64(game.width-1))); j++ {
				if !game.field[i][j].isOpened {
					_ = game.openCell(game.field[i][j])
				}
			}
		}
	}

	if game.isWin() {
		return errYouWin{}
	}

	return nil
}

func (game *Game) initField(level int) error {
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
		return errUnknownLevel(game.level)
	}
	return nil
}

func (game *Game) initCells() {
	game.field = make([][]Cell, game.height)
	for i := 0; i < game.height; i++ {
		game.field[i] = make([]Cell, game.width)
		for j := 0; j < game.width; j++ {
			game.field[i][j].x = j
			game.field[i][j].y = i
		}
	}
}

func generateBombsCoordinates(width, height, bombs int) ([]int, []int) {
	rand.Seed(time.Now().Unix())

	widthCoords, heightCoords := make([]int, bombs), make([]int, bombs)
	for i := 0; i < bombs; i++ {
		widthCoords[i] = rand.Intn(width)
		heightCoords[i] = rand.Intn(height)
	}

	uniqueStart := 0
	for uniqueStart != bombs {
		for i := uniqueStart; i < bombs; i++ {
			for j := i + 1; j < bombs; j++ {
				if widthCoords[i] == widthCoords[j] && heightCoords[i] == heightCoords[j] {
					widthCoords[j] = rand.Intn(width)
					heightCoords[j] = rand.Intn(height)
					j--
				}
			}
			uniqueStart++
		}
	}

	return widthCoords, heightCoords
}

func (game *Game) generateBombs() {
	widthCoords, heightCoords := generateBombsCoordinates(game.width, game.height, game.bombs)
	for i := 0; i < game.bombs; i++ {
		game.field[heightCoords[i]][widthCoords[i]].isBomb = true
	}
}

func (game *Game) init(level int) {
	err := game.initField(level)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	game.initCells()
	game.generateBombs()
}
