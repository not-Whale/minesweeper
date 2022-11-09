package cli

import (
	"fmt"
	"os"

	. "github.com/not-Whale/minesweeper/game"
)

func printField(game Game) {
	for i := -1; i < game.Height; i++ {
		for j := -1; j < game.Width; j++ {
			if i == -1 {
				fmt.Printf("%s", BlueColor)
				if j == -1 {
					fmt.Printf("  ")
				} else {
					fmt.Printf("%v ", j)
				}
				fmt.Printf("%s", ResetColor)
			} else if j == -1 {
				fmt.Printf("%s%v %s", BlueColor, i, ResetColor)
			} else {
				if game.Field[i][j].IsOpened {
					fmt.Printf("%s%v %s", YellowColor, game.Field[i][j].BombsAround, ResetColor)
				} else if game.Field[i][j].IsMarked {
					fmt.Printf("%s! %s", RedColor, ResetColor)
				} else {
					fmt.Printf("? ")
				}
			}
		}
		fmt.Printf("\n")
	}
}

func readLevel() (level int) {
	fmt.Printf("Выберите уровень сложности:\n")
	fmt.Printf("1. Легкий\n2. Нормальный\n3. Сложный\n")
	_, err := fmt.Scan(&level)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

func readAction() (action int) {
	fmt.Printf("Что сделать?\n")
	fmt.Printf("1. Открыть ячейку\n2. Отметить бомбу\n")
	fmt.Printf("3. Снять отметку\n4. Выйти\n")
	_, err := fmt.Scan(&action)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

func readCoordinates() (x int, y int) {
	fmt.Printf("Введите координаты ячейки в формате: номер_строки номер_столбца\n")
	_, err := fmt.Scan(&x, &y)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

func openHandler(game Game) {
	x, y := readCoordinates()
	err := game.OpenCell(game.Field[x][y])
	if err != nil {
		switch err.(type) {
		case ErrOutOfRange, ErrCellIsAlreadyOpened, ErrOpenMarkedCell:
			fmt.Println(err)
		case ErrGameOver, ErrYouWin:
			game.OpenAll()
			printField(game)
			fmt.Println(err)
			os.Exit(0)
		default:
			fmt.Println("Неизвестная ошибка!")
			os.Exit(1)
		}
	}
}

func markHandler(game Game) {
	x, y := readCoordinates()
	err := game.MarkCell(game.Field[x][y])
	if err != nil {
		switch err.(type) {
		case ErrOutOfRange, ErrCellIsAlreadyMarked, ErrMarkOpenedCell:
			fmt.Println(err)
		case ErrYouWin:
			game.OpenAll()
			printField(game)
			fmt.Println(err)
			os.Exit(0)
		default:
			fmt.Println("Неизвестная ошибка!")
			os.Exit(1)
		}
	}
}

func unmarkHandler(game Game) {
	x, y := readCoordinates()
	err := game.UnmarkCell(game.Field[x][y])
	if err != nil {
		switch err.(type) {
		case ErrOutOfRange, ErrCellIsAlreadyUnmarked:
			fmt.Println(err)
		default:
			fmt.Println("Неизвестная ошибка!")
			os.Exit(1)
		}
	}
}

func StartConsoleGame() {
	game := Game{}
	level := readLevel()
	game.Init(level)

	for {
		printField(game)
		switch action := readAction(); action {
		case 1:
			openHandler(game)
		case 2:
			markHandler(game)
		case 3:
			unmarkHandler(game)
		case 4:
			os.Exit(0)
		default:
			fmt.Printf("%sНеизвестная команда! Попробуйте еще раз...%s", ErrorColor, ResetColor)
		}
	}
}
