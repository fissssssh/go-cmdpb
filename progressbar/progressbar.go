package progressbar

import (
	"errors"
	"fmt"
	"strings"
)

type ProgressBar interface {
	Value() float64
	Min() float64
	Max() float64
	BlankChar() rune
	FillChar() rune
	Length() int

	SetValue(float64)
	SetMin(float64) error
	SetMax(float64) error
	SetBlankChar(rune)
	SetFillChar(rune)
	SetLength(int)
}

type CmdProgressBar interface {
	ProgressBar
	Print(label string)
}

func NewCmdProgressBar() CmdProgressBar {
	return &cmdProgressBar{
		value:     0,
		min:       0,
		max:       100,
		blankChar: '_',
		fillChar:  '*',
		length:    25,
	}
}

type cmdProgressBar struct {
	value     float64
	min       float64
	max       float64
	blankChar rune
	fillChar  rune
	length    int
}

func (b *cmdProgressBar) Value() float64 {
	return b.value
}

func (b *cmdProgressBar) Min() float64 {
	return b.min
}

func (b *cmdProgressBar) Max() float64 {
	return b.max
}

func (b *cmdProgressBar) BlankChar() rune {
	return b.blankChar
}

func (b *cmdProgressBar) FillChar() rune {
	return b.fillChar
}

func (b *cmdProgressBar) Length() int {
	return b.length
}

func (b *cmdProgressBar) SetValue(value float64) {
	if value < b.min {
		b.value = b.min
	} else if value > b.max {
		b.value = b.max
	} else {
		b.value = value
	}
}

func (b *cmdProgressBar) SetMin(value float64) error {
	if value > b.max {
		return errors.New("progressbar min value can not be gather than max value")
	}
	b.min = value
	return nil
}

func (b *cmdProgressBar) SetMax(value float64) error {
	if value < b.min {
		return errors.New("progressbar max value can not be less than min value")
	}
	b.max = value
	return nil
}

func (b *cmdProgressBar) SetBlankChar(c rune) {
	b.blankChar = c
}
func (b *cmdProgressBar) SetFillChar(c rune) {
	b.fillChar = c
}
func (b *cmdProgressBar) SetLength(l int) {
	b.length = l
}
func (b *cmdProgressBar) Print(label string) {
	percent := (b.value - b.min) / (b.max - b.min)
	fillCount := int(percent * float64(b.length))
	blankCount := b.length - fillCount
	bar := strings.Builder{}
	for i := 0; i < fillCount; i++ {
		bar.WriteRune(b.fillChar)
	}
	for i := 0; i < blankCount; i++ {
		bar.WriteRune(b.blankChar)
	}
	fmt.Printf("\r[%s] %s %8.2f%%", &bar, label, percent*100)
}
