package minesweeper

import (
	"fmt"
)

type errCellIsAlreadyMarked struct {
	x, y int
}

type errCellIsAlreadyNotMarked struct {
	x, y int
}

type errGameOver struct {
	x, y int
}

type errOutOfRange struct {
	x, y int
}

type errYouWin struct{}

type errUnknownLevel int

func (e errCellIsAlreadyMarked) Error() string {
	return fmt.Sprintf("Ячейка [%v, %v] уже отмечена!", e.x, e.y)
}

func (e errCellIsAlreadyNotMarked) Error() string {
	return fmt.Sprintf("Ячейка [%v, %v] и так не отмечена!", e.x, e.y)
}

func (e errGameOver) Error() string {
	return fmt.Sprintf("Вы проиграли! В ячейке [%v, %v] оказалась бомба!", e.x, e.y)
}

func (e errOutOfRange) Error() string {
	return fmt.Sprintf("Ячейка [%v, %v] находится за пределами игрового поля!", e.x, e.y)
}

func (e errUnknownLevel) Error() string {
	return fmt.Sprintf("Неизвестный уровень сложности %v!", int(e))
}

func (e errYouWin) Error() string {
	return "Поздравляю! Вы победили!"
}
