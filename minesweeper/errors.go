package minesweeper

import (
	"fmt"
)

type errCellIsAlreadyMarked struct {
	x, y int
}

type errCellIsAlreadyUnmarked struct {
	x, y int
}

type errCellIsAlreadyOpened struct {
	x, y int
}

type errGameOver struct {
	x, y int
}

type errOutOfRange struct {
	x, y int
}

type errMarkOpenedCell struct {
	x, y int
}

type errOpenMarkedCell struct {
	x, y int
}

type errYouWin struct{}

type errUnknownLevel int

func (e errCellIsAlreadyMarked) Error() string {
	return fmt.Sprintf("%sЯчейка [%v, %v] уже отмечена!%s", ErrorColor, e.x, e.y, ResetColor)
}

func (e errCellIsAlreadyUnmarked) Error() string {
	return fmt.Sprintf("%sЯчейка [%v, %v] и так не отмечена!%s", ErrorColor, e.x, e.y, ResetColor)
}

func (e errCellIsAlreadyOpened) Error() string {
	return fmt.Sprintf("%sЯчейка [%v, %v] уже открыта!%s", ErrorColor, e.x, e.y, ResetColor)
}

func (e errGameOver) Error() string {
	return fmt.Sprintf("%sВы проиграли! В ячейке [%v, %v] оказалась бомба!%s", ErrorColor, e.x, e.y, ResetColor)
}

func (e errOutOfRange) Error() string {
	return fmt.Sprintf("%sЯчейка [%v, %v] находится за пределами игрового поля!%s", ErrorColor, e.x, e.y, ResetColor)
}

func (e errMarkOpenedCell) Error() string {
	return fmt.Sprintf("%sЯчейка [%v, %v], которую Вы хотите отметить, открыта!%s", ErrorColor, e.x, e.y, ResetColor)
}

func (e errOpenMarkedCell) Error() string {
	return fmt.Sprintf("%sЯчейка [%v, %v], которую Вы хотите открыть, отмечена! Уберите метку и попробуйте еще раз!%s", ErrorColor, e.x, e.y, ResetColor)
}

func (e errUnknownLevel) Error() string {
	return fmt.Sprintf("%sНеизвестный уровень сложности \"%v\"!%s", ErrorColor, int(e), ResetColor)
}

func (e errYouWin) Error() string {
	return fmt.Sprintf("%sПоздравляю! Вы победили!%s", WinColor, ResetColor)
}
