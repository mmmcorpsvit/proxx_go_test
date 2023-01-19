package main

import (
	"fmt"
)

// Field represents a two-dimensional field of cells.
type Field struct {
	s    [][]int
	w, h int
}

func NewField(h, w, k int) *Field {
	// TODO: add validator
	// TODO: check math lib to create matrix
	// https://stackoverflow.com/questions/39804861/what-is-a-concise-way-to-create-a-2d-slice-in-go
	s := make([][]int, h)
	for i := range s {
		s[i] = make([]int, w)
	}
	return &Field{s: s, w: w, h: h}
}

func PrintField(f *Field) {
	for i := 0; i < f.h; i++ {
		for k := 0; k < f.w; k++ {
			// TODO: do better output
			//fmt.Print(strconv.FormatInt(int64(f.s[i][k]), 10) + "   ")
			fmt.Print(fmt.Sprintf("%d   ", f.s[i][k]))
		}
		fmt.Println()
	}
}

// String returns the game board as a string.
//func (l *Life) String() string {
//	var buf bytes.Buffer
//	for y := 0; y < l.h; y++ {
//		for x := 0; x < l.w; x++ {
//			b := byte(' ')
//			if l.a.Alive(x, y) {
//				b = '*'
//			}
//			buf.WriteByte(b)
//		}
//		buf.WriteByte('\n')
//	}
//	return buf.String()
//}

func main() {
	data := NewField(10, 10, 5)
	PrintField(data)
	//_ = data
	//fmt.Println("Hello, 世界")
}
