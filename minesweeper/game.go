package minesweeper

import (
	"fmt"
	"os"
)

func printField(game Game) {
	for i := -1; i < game.height; i++ {
		for j := -1; j < game.width; j++ {
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
				if game.field[i][j].isOpened {
					fmt.Printf("%s%v %s", YellowColor, game.field[i][j].bombsAround, ResetColor)
				} else if game.field[i][j].isMarked {
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

func StartConsoleGame() {
	game := Game{}
	level := readLevel()
	game.init(level)

	for {
		printField(game)
		switch action := readAction(); action {
		case 1:
			x, y := readCoordinates()
			err := game.openCell(game.field[x][y])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case 2:
			x, y := readCoordinates()
			err := game.markCell(game.field[x][y])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case 3:
			x, y := readCoordinates()
			err := game.unmarkCell(game.field[x][y])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case 4:
			os.Exit(0)
		default:
			fmt.Printf("Неизвестная команда! Попробуйте еще раз...")
		}
	}
}
