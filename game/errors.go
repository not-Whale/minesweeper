package game

import (
	"fmt"
)

type ErrCellIsAlreadyMarked struct {
	x, y int
}

type ErrCellIsAlreadyUnmarked struct {
	x, y int
}

type ErrCellIsAlreadyOpened struct {
	x, y int
}

type ErrGameOver struct {
	x, y int
}

type ErrOutOfRange struct {
	x, y int
}

type ErrMarkOpenedCell struct {
	x, y int
}

type ErrOpenMarkedCell struct {
	x, y int
}

type ErrYouWin struct{}

type ErrUnknownLevel int

func (e ErrCellIsAlreadyMarked) Error() string {
	return fmt.Sprintf("%sЯчейка [%v, %v] уже отмечена!%s", ErrorColor, e.x, e.y, ResetColor)
}

func (e ErrCellIsAlreadyUnmarked) Error() string {
	return fmt.Sprintf("%sЯчейка [%v, %v] и так не отмечена!%s", ErrorColor, e.x, e.y, ResetColor)
}

func (e ErrCellIsAlreadyOpened) Error() string {
	return fmt.Sprintf("%sЯчейка [%v, %v] уже открыта!%s", ErrorColor, e.x, e.y, ResetColor)
}

func (e ErrGameOver) Error() string {
	return fmt.Sprintf("%sВы проиграли! В ячейке [%v, %v] оказалась бомба!%s", ErrorColor, e.x, e.y, ResetColor)
}

func (e ErrOutOfRange) Error() string {
	return fmt.Sprintf("%sЯчейка [%v, %v] находится за пределами игрового поля!%s", ErrorColor, e.x, e.y, ResetColor)
}

func (e ErrMarkOpenedCell) Error() string {
	return fmt.Sprintf("%sЯчейка [%v, %v], которую Вы хотите отметить, открыта!%s", ErrorColor, e.x, e.y, ResetColor)
}

func (e ErrOpenMarkedCell) Error() string {
	return fmt.Sprintf("%sЯчейка [%v, %v], которую Вы хотите открыть, отмечена! Уберите метку и попробуйте еще раз!%s", ErrorColor, e.x, e.y, ResetColor)
}

func (e ErrUnknownLevel) Error() string {
	return fmt.Sprintf("%sНеизвестный уровень сложности \"%v\"!%s", ErrorColor, int(e), ResetColor)
}

func (e ErrYouWin) Error() string {
	return fmt.Sprintf("%sПоздравляю! Вы победили!%s", WinColor, ResetColor)
}
