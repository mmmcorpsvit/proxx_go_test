package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

// Field represents a two-dimensional field of cells.
type Field struct {
	cell [][]int
	//width, height int
}

// Game Config
const (
	GameFieldHeight     = 4
	GameFieldWidth      = 4
	GameFieldBlackHoles = 1
	GameFieldClicks     = 5
)

var (
	GameVisibleField *Field
	GameField        *Field
)

func GetRandomInt(value int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(value)))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return int(n)
}

func SetSurrounding(cell [][]int, x, y, dx, dy int) [][]int {
	// Check boundaries
	if x+dx >= 0 && y+dy >= 0 &&
		x+dx < GameFieldHeight && y+dy < GameFieldWidth &&
		// This is not Black Hole ?
		cell[x+dx][y+dy] != -1 {
		cell[x+dx][y+dy]++ // Inc counter
	}

	return cell
}

func NewField(gameFieldBlackHoles int) *Field {
	// TODO: add validator, h>2, w>2, black_holes_count > 0
	// TODO: check math lib to create matrix
	// https://stackoverflow.com/questions/39804861/what-is-a-concise-way-to-create-a-2d-slice-in-go
	f := make([][]int, GameFieldWidth)
	for i := range f {
		f[i] = make([]int, GameFieldHeight)
	}

	// add Black Holes
	blackHolesCounter := 0
	x := 0
	y := 0
	for blackHolesCounter < gameFieldBlackHoles {
		x = GetRandomInt(GameFieldWidth)
		y = GetRandomInt(GameFieldHeight)

		// update only this not Black Hole cell
		if f[x][y] != -1 {
			f[x][y] = -1 // There Black Hole
			blackHolesCounter++
		}
	}

	// compute cells
	for x := 0; x < GameFieldWidth; x++ {
		for y := 0; y < GameFieldHeight; y++ {
			if f[x][y] == -1 {
				SetSurrounding(f, x, y, -1, -1)
				SetSurrounding(f, x, y, -1, 0)
				SetSurrounding(f, x, y, -1, 1)

				SetSurrounding(f, x, y, 1, -1)
				SetSurrounding(f, x, y, 1, 0)
				SetSurrounding(f, x, y, 1, 1)

				SetSurrounding(f, x, y, 0, -1)
				SetSurrounding(f, x, y, 0, 1)
			}
		}
	}

	return &Field{cell: f}
	//return &Field{cell: f, width: width, height: height}
}

func Display(debug bool) {
	fmt.Print("   ")
	for x := 0; x < GameFieldHeight; x++ {
		fmt.Print(x)
		fmt.Print("  ")
	}
	fmt.Println()

	fmt.Print("   ")
	for x := 0; x < GameFieldHeight; x++ {
		fmt.Print("-")
		fmt.Print("  ")
	}
	fmt.Println()

	for x := 0; x < GameFieldWidth; x++ {
		s := ""
		for y := 0; y < GameFieldHeight; y++ {
			// TODO: do better output
			//fmt.Print(strconv.FormatInt(int64(f.s[i][k]), 10) + "   ")
			//value := 0
			//value = f.cell[x][y]
			v := "   "
			if debug == true || (debug == false && GameVisibleField.cell[x][y] == 1) {
				v = strings.Replace(
					fmt.Sprintf("%d  ", GameField.cell[x][y]), "-1", "H", -1)
			}

			//if debug == true {
			//	v = strings.Replace(
			//		fmt.Sprintf("%d  ", GameField.cell[x][y]), "-1", "H", -1)
			//}

			//else {
			//	v = "  "
			//}

			s = s + v
		}
		// fmt.Print( x + "ddd" + s)
		fmt.Print(x)
		fmt.Print("| ")
		fmt.Println(s)
	}
}

type GameVisibleCoord struct {
	x, y int
}

func IndexOf[T comparable](collection []T, el T) int {
	for i, x := range collection {
		if x == el {
			return i
		}
	}
	return -1
}

func SetSurroundingEmptyVisible(cell [][]int, slice []GameVisibleCoord, x, y, dx, dy int) []GameVisibleCoord {
	// Check boundaries
	if x+dx >= 0 && y+dy >= 0 &&
		x+dx < GameFieldHeight && y+dy < GameFieldWidth &&
		// This is not Black Hole ?
		cell[x+dx][y+dy] == 0 &&
		// coord not yet visible ?
		IndexOf(slice, GameVisibleCoord{x: x + dx, y: y + dy}) == -1 {
		//cell[x+dx][y+dy]++ // Inc counter
		GameVisibleField.cell[x+dx][y+dy] = 1

		slice = append(slice, GameVisibleCoord{x: x + dx, y: y + dy})
		fmt.Println(fmt.Sprintf("debug set %d, %d", x+dx, y+dy))
	}

	return slice
}

func Click(x, y int, debug bool) {
	// if cell already was clicked - ignore
	if GameVisibleField.cell[y][x] == 1 {
		if debug == true {
			fmt.Println("You already clicked at this point. Just ignore click")
		}
		return
	}

	// do cell visible any way
	//if GameField.cell[y][x] != 1 {
	GameVisibleField.cell[y][x] = 1
	//}

	// click at Black Hole, just write warning
	if GameField.cell[y][x] == -1 && debug == false {
		fmt.Println("You click at Black Hole. Game Over!")
		os.Exit(0)
		return
	}

	f := GameVisibleField.cell
	var slice = make([]GameVisibleCoord, 0)

	if GameField.cell[y][x] == 0 {
		// show all empty cells
		//old_array:= [...]int

		slice = append(slice, GameVisibleCoord{x, y})

		slice = SetSurroundingEmptyVisible(f, slice, x, y, -1, -1)
		slice = SetSurroundingEmptyVisible(f, slice, x, y, -1, 0)
		slice = SetSurroundingEmptyVisible(f, slice, x, y, -1, 1)

		slice = SetSurroundingEmptyVisible(f, slice, x, y, 1, -1)
		slice = SetSurroundingEmptyVisible(f, slice, x, y, 1, 0)
		slice = SetSurroundingEmptyVisible(f, slice, x, y, 1, 1)

		slice = SetSurroundingEmptyVisible(f, slice, x, y, 0, -1)
		slice = SetSurroundingEmptyVisible(f, slice, x, y, 0, 1)

		fmt.Println(slice)
	}

}

//func Bye() {
//	fmt.Println()
//	fmt.Println()
//	fmt.Println("Developed by MMM_Corp, test task special for Data Science UA, 2023")
//	fmt.Println("Skype: mmm_ogame")
//}

func main() {
	//fmt.Println("*********************************")
	//fmt.Println("*            PROXX              *")
	//fmt.Println("*                               *")
	//fmt.Println("* Legend:                       *")
	//fmt.Println("*    H   - Black Hole           *")
	//fmt.Println("*    0   - Visible Cell         *")
	//fmt.Println("*        - Hidden Cell          *")
	//fmt.Println("*    1-8 - Surrounding Cell     *")
	//fmt.Println("*********************************")
	//fmt.Println()

	GameField = NewField(GameFieldBlackHoles) // Generated field
	GameVisibleField = NewField(0)            // Visible Field

	fmt.Println()
	fmt.Println("Debug Game Field")
	Display(true)
	fmt.Println()
	fmt.Println("Simulate few clicks at random places")

	//x := 0
	//y := 0
	for i := 0; i < GameFieldClicks; i++ {
		x := GetRandomInt(GameFieldHeight)
		y := GetRandomInt(GameFieldWidth)
		fmt.Println()
		fmt.Println(fmt.Sprintf("Clicked at: %d, %d", x, y))
		Click(x, y, false)
		Display(false)
	}

	//defer Bye()
}
