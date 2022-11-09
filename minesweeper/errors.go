package minesweeper

import (
	"fmt"
)

type errCellIsMarked struct {
	x, y int
}

type errCellIsNotMarked struct {
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

func (e errCellIsMarked) Error() string {
	return fmt.Sprintf("ЯЧЕЙКА [%v, %v] УЖЕ ОТМЕЧЕНА", e.x, e.y)
}

func (e errCellIsNotMarked) Error() string {
	return fmt.Sprintf("ЯЧЕЙКА [%v, %v] И ТАК НЕ ОТМЕЧЕНА", e.x, e.y)
}

func (e errGameOver) Error() string {
	return fmt.Sprintf("ВЫ ПРОИГРАЛИ: В ЯЧЕЙКЕ [%v, %v] ОКАЗАЛАСЬ БОМБА", e.x, e.y)
}

func (e errOutOfRange) Error() string {
	return fmt.Sprintf("ЯЧЕЙКА [%v, %v] НАХОДИТСЯ ЗА ПРЕДЕЛЯМИ ИГРОВОГО ПОЛЯ", e.x, e.y)
}

func (e errUnknownLevel) Error() string {
	return fmt.Sprintf("НЕИЗВЕСТНЫЙ УРОВЕНЬ СЛОЖНОСТИ %v", int(e))
}

func (e errYouWin) Error() string {
	return "ВЫ ПОБЕДИЛИ"
}
