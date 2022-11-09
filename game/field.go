package game

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

type Game struct {
	Field         [][]Cell
	Width, Height int
	level         int
	bombs         int
	openedCells   int
	markedCells   int
	markBalance   int
}

type Cell struct {
	x, y        int
	BombsAround int
	isBomb      bool
	IsOpened    bool
	IsMarked    bool
}

func (game *Game) OpenAll() {
	for i := 0; i < game.Height; i++ {
		for j := 0; j < game.Width; j++ {
			if !game.Field[i][j].isBomb && !game.Field[i][j].IsOpened {
				_ = game.OpenCell(game.Field[i][j])
			}
			if game.Field[i][j].isBomb && !game.Field[i][j].IsMarked {
				_ = game.MarkCell(game.Field[i][j])
			}
		}
	}
}

func (game *Game) isWin() (ok bool) {
	if game.markedCells == game.bombs &&
		// game.openedCells == game.Width*game.Height-game.bombs &&
		game.markBalance == 0 {
		ok = true
	}
	return
}

func (game *Game) UnmarkCell(cell Cell) error {
	if cell.y >= game.Height || cell.x >= game.Width ||
		cell.y < 0 || cell.x < 0 {
		return ErrOutOfRange{cell.x, cell.y}
	}

	if !cell.IsMarked {
		return ErrCellIsAlreadyUnmarked{cell.x, cell.y}
	}

	if !cell.isBomb {
		game.markBalance++
	}

	game.Field[cell.y][cell.x].IsMarked = false
	game.markedCells--

	return nil
}

func (game *Game) MarkCell(cell Cell) error {
	if cell.y >= game.Height || cell.x >= game.Width ||
		cell.y < 0 || cell.x < 0 {
		return ErrOutOfRange{cell.x, cell.y}
	}

	if cell.IsMarked {
		return ErrCellIsAlreadyMarked{cell.x, cell.y}
	}

	if cell.IsOpened {
		return ErrMarkOpenedCell{cell.x, cell.y}
	}

	if !cell.isBomb {
		game.markBalance--
	}

	game.Field[cell.y][cell.x].IsMarked = true
	game.markedCells++

	if game.isWin() {
		return ErrYouWin{}
	}

	return nil
}

func (game *Game) calcBombsAround(cell Cell) (count int) {
	for i := int(math.Max(float64(cell.y-1), 0)); i <= int(math.Min(float64(cell.y+1), float64(game.Height-1))); i++ {
		for j := int(math.Max(float64(cell.x-1), 0)); j <= int(math.Min(float64(cell.x+1), float64(game.Width-1))); j++ {
			if game.Field[i][j].isBomb {
				count++
			}
		}
	}
	return
}

func (game *Game) OpenCell(cell Cell) error {
	if cell.y >= game.Height || cell.x >= game.Width ||
		cell.y < 0 || cell.x < 0 {
		return ErrOutOfRange{cell.x, cell.y}
	}

	if cell.IsOpened {
		return ErrCellIsAlreadyOpened{cell.x, cell.y}
	}

	if cell.IsMarked {
		return ErrOpenMarkedCell{cell.x, cell.y}
	}

	if cell.isBomb {
		return ErrGameOver{cell.x, cell.y}
	}

	game.openedCells++
	game.Field[cell.y][cell.x].IsOpened = true
	game.Field[cell.y][cell.x].BombsAround = game.calcBombsAround(cell)

	if game.Field[cell.y][cell.x].BombsAround == 0 {
		for i := int(math.Max(float64(cell.y-1), 0)); i <= int(math.Min(float64(cell.y+1), float64(game.Height-1))); i++ {
			for j := int(math.Max(float64(cell.x-1), 0)); j <= int(math.Min(float64(cell.x+1), float64(game.Width-1))); j++ {
				if !game.Field[i][j].IsOpened {
					_ = game.OpenCell(game.Field[i][j])
				}
			}
		}
	}

	if game.isWin() {
		return ErrYouWin{}
	}

	return nil
}

func (game *Game) initField(level int) error {
	switch game.level = level; game.level {
	case 1:
		game.Width = EasyWidth
		game.Height = EasyHeight
		game.bombs = EasyBombs
	case 2:
		game.Width = MediumWidth
		game.Height = MediumHeight
		game.bombs = MediumBombs
	case 3:
		game.Width = HardWidth
		game.Height = HardHeight
		game.bombs = HardBombs
	default:
		return ErrUnknownLevel(game.level)
	}
	return nil
}

func (game *Game) initCells() {
	game.Field = make([][]Cell, game.Height)
	for i := 0; i < game.Height; i++ {
		game.Field[i] = make([]Cell, game.Width)
		for j := 0; j < game.Width; j++ {
			game.Field[i][j].x = j
			game.Field[i][j].y = i
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
	widthCoords, heightCoords := generateBombsCoordinates(game.Width, game.Height, game.bombs)
	for i := 0; i < game.bombs; i++ {
		game.Field[heightCoords[i]][widthCoords[i]].isBomb = true
	}
}

func (game *Game) Init(level int) {
	err := game.initField(level)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	game.initCells()
	game.generateBombs()
}
