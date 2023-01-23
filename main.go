package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

// Game Config
const (
	GameFieldWidth      = 10
	GameFieldHeight     = 5
	GameFieldBlackHoles = 3
	GameFieldClicks     = 2
)

// Field represents a two-dimensional field of cells.
type Field struct {
	cell [][]int
	//width, height int
}

var (
	GameField        *Field
	GameFieldVisible *Field
)

func GetRandomInt(value int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(value)))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return int(n)
}

func SetSurrounding(cell [][]int, x, y, dx, dy int) {
	dx = x + dx
	dy = y + dy
	if dx >= 0 && dy >= 0 && // Check boundaries
		dx < GameFieldHeight && dy < GameFieldWidth &&
		// This is not Black Hole ?
		cell[dx][dy] != -1 {
		cell[dx][dy]++ // Inc cell counter
	}
}

func CreateFieldMatrix(width, height int) [][]int {
	f := make([][]int, GameFieldHeight)
	for i := range f {
		f[i] = make([]int, GameFieldWidth)
	}
	return f
}

// NewField Create new game field. If gameFieldBlackHoles = 0, just create empty GameFieldVisible matrix
func NewField(gameFieldBlackHoles int) *Field {
	f := CreateFieldMatrix(GameFieldWidth, GameFieldHeight)

	if gameFieldBlackHoles > 0 {
		// add Black Holes
		blackHolesCounter := 0
		for blackHolesCounter < gameFieldBlackHoles {
			x := GetRandomInt(GameFieldHeight)
			y := GetRandomInt(GameFieldWidth)

			// update only this not Black Hole cell
			if f[x][y] != -1 {
				f[x][y] = -1 // There Black Hole
				blackHolesCounter++

				// compute cells
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

	//fmt.Println(f)
	return &Field{cell: f}
}

func Display(debug bool) {
	fmt.Print("   ")
	for x := 0; x < GameFieldWidth; x++ {
		fmt.Print(x)
		fmt.Print("  ")
	}
	fmt.Println()

	fmt.Print("   ")
	for x := 0; x < GameFieldWidth; x++ {
		fmt.Print("-")
		fmt.Print("  ")
	}
	fmt.Println()

	for x := 0; x < GameFieldHeight; x++ {
		s := ""
		for y := 0; y < GameFieldWidth; y++ {
			// TODO: do better output
			//fmt.Print(strconv.FormatInt(int64(f.s[i][k]), 10) + "   ")
			//value := 0
			//value = f.cell[x][y]
			v := "   "
			if debug == true || (debug == false && GameFieldVisible.cell[x][y] == 1) {
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

//func SetSurroundingEmptyVisible(cell [][]int, slice []GameVisibleCoord, x, y, dx, dy int) []GameVisibleCoord {
//	dx = x + dx
//	// Check boundaries
//	if x+dx >= 0 && y+dy >= 0 &&
//		x+dx < GameFieldHeight && y+dy < GameFieldWidth &&
//		// This is not Black Hole ?
//		cell[x+dx][y+dy] != -1 &&
//		// coord not yet visible ?
//		IndexOf(slice, GameVisibleCoord{x: x + dx, y: y + dy}) == -1 {
//		//cell[x+dx][y+dy]++ // Inc counter
//		GameFieldVisible.cell[x+dx][y+dy] = 1
//		Click(x+dx, y+dy, true) // recurse, but can be another way - use shift slice for loop
//
//		slice = append(slice, GameVisibleCoord{x: x + dx, y: y + dy})
//		//fmt.Println(fmt.Sprintf("debug set %d, %d", x+dx, y+dy))
//	}
//
//	return slice
//}

func SetSurroundingEmptyVisible(cell [][]int, slice []GameVisibleCoord, x, y, dx, dy int) []GameVisibleCoord {
	dx = x + dx
	dy = y + dy
	// Check boundaries
	if dx >= 0 && dy >= 0 &&
		dx < GameFieldHeight && dy < GameFieldWidth &&
		// This is not Black Hole ?
		cell[dx][dy] != -1 &&
		// coord not yet visible ?
		IndexOf(slice, GameVisibleCoord{x: dx, y: dy}) == -1 {
		//cell[x+dx][y+dy]++ // Inc counter
		GameFieldVisible.cell[dx][dy] = 1
		//Click(dx, dy, true) // recurse, but can be another way - use shift slice for loop

		slice = append(slice, GameVisibleCoord{x: dx, y: dy})
		//fmt.Println(fmt.Sprintf("debug set %d, %d", x+dx, y+dy))
	}

	return slice
}

func Click(x, y int, debug bool) {
	// if cell already was clicked - ignore
	if GameFieldVisible.cell[x][y] == 1 {
		if debug == true {
			//fmt.Println("You already clicked at this point. Just ignore click")
		}
		return
	}

	// do cell visible any way
	//if GameField.cell[y][x] != 1 {
	GameFieldVisible.cell[x][y] = 1
	//}

	// click at Black Hole, just write warning
	if GameField.cell[x][y] == -1 && debug == false {
		fmt.Println("You click at Black Hole. Game Over!")
		os.Exit(0)
		return
	}

	f := GameFieldVisible.cell
	var slice = make([]GameVisibleCoord, 0)

	if GameField.cell[x][y] == 0 {
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

		//fmt.Println(slice)
		return
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
	//fmt.Println("*********************************\n")

	if //goland:noinspection ALL
	GameFieldBlackHoles > GameFieldWidth*GameFieldHeight {
		fmt.Printf("Error, possible maximum Black Holes Count: %v", GameFieldWidth*GameFieldHeight)
		os.Exit(0)
	}

	GameField = NewField(GameFieldBlackHoles) // Generated field
	GameFieldVisible = NewField(0)            // Visible Field

	fmt.Println("\nGenerated Game Field")
	Display(true)
	fmt.Println("\nSimulate few clicks at random places")

	for i := 0; i < GameFieldClicks; i++ {
		y := GetRandomInt(GameFieldHeight)
		x := GetRandomInt(GameFieldWidth)
		fmt.Println()
		fmt.Printf("Clicked at: %d, %d \n", x, y)
		Click(y, x, false)
		Display(false)
	}

	_ = 1
	fmt.Println(GameFieldVisible)
	//defer Bye()
}
