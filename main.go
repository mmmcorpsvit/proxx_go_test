package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const (
	GameFieldWidth      = 10
	GameFieldHeight     = 10
	GameFieldBlackHoles = 5
	GameFieldClicks     = 5
)

type ShiftCoordinateElement struct {
	x, y int
}

// ShiftCoordinate array shift to get adjustment, see keyboard num block how comment
var ShiftCoordinate = [...]ShiftCoordinateElement{
	{x: -1, y: -1}, // 7
	{x: 0, y: -1},  // 8
	{x: 1, y: -1},  // 9
	{x: 1, y: 0},   // 3
	{x: 1, y: 1},   // 2
	{x: 0, y: 1},   // 1

	{x: -1, y: 1}, // 4
	{x: -1, y: 0}, // 7
}

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
		dx < GameFieldWidth && dy < GameFieldHeight &&
		// This is not Black Hole ?
		cell[dx][dy] != -1 {
		cell[dx][dy]++ // Inc cell counter
	}
}

func CreateFieldMatrix(width, height int) [][]int {
	f := make([][]int, width)
	for i := range f {
		f[i] = make([]int, height)
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
			x := GetRandomInt(GameFieldWidth)
			y := GetRandomInt(GameFieldHeight)

			// update only this not Black Hole cell
			if f[x][y] != -1 {
				f[x][y] = -1 // There Black Hole
				blackHolesCounter++

				// compute cells
				for _, element := range ShiftCoordinate {
					SetSurrounding(f, x, y, element.x, element.y)
				}
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

	for y := 0; y < GameFieldHeight; y++ {
		s := ""
		for x := 0; x < GameFieldWidth; x++ {
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
		fmt.Print(y)
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

func SetSurroundingEmptyVisible(cell [][]int, slice []GameVisibleCoord, realClick bool, x, y, dx, dy int) []GameVisibleCoord {
	//func SetSurroundingEmptyVisible(cell [][]int, x, y, dx, dy int) {
	dx = x + dx
	dy = y + dy

	fmt.Println("********************")

	// Check boundaries
	if dx >= 0 && dy >= 0 &&
		dx < GameFieldHeight && dy < GameFieldWidth &&
		// This is not Black Hole ?
		cell[dx][dy] != -1 {

		// coord not yet visible ?
		//if IndexOf(slice, GameVisibleCoord{y: dx, x: dy}) == -1 {
		fmt.Printf("debug SetSurroundingEmptyVisible set %d, %d", dy, dx)
		//cell[x+dx][y+dy]++ // Inc counter
		GameFieldVisible.cell[dx][dy] = 1

		//if realClick == false {
		if cell[dx][dy] == 0 {
			fmt.Printf("debug CELL CLICK  %d, %d", dy, dx)
			Click(dx, dy, false) // recurse, but can be another way - use shift slice for loop
		}

		//slice = append(slice, GameVisibleCoord{y: dx, x: dy})
		//fmt.Printf("debug set %d, %d", dy, dx)
		//}
	}

	return slice
}

func Click(x, y int, realClick bool) {
	//fmt.Printf("DEBUG Click set %d, %d\n", x, y)

	// cell already was clicked, ignore click and show message
	if GameFieldVisible.cell[x][y] == 1 && realClick == true {
		fmt.Printf("You already do Click, ignoring... %d, %d\n", x, y)
		return
	}

	// number cell, show
	if GameField.cell[x][y] > 0 {
		GameFieldVisible.cell[x][y] = 1
		return
	}

	// Black Hole cell, just write warning and break app
	if GameField.cell[x][y] == -1 && realClick == true {
		fmt.Println("You click at Black Hole. Game Over!")
		os.Exit(0)
		return
	}

	if GameField.cell[x][y] == 0 {
		//var slice = make([]GameVisibleCoord, 0)
		var visited = make(map[GameVisibleCoord]bool, 0)
		//firstRun := true
		//slice = append(slice, GameVisibleCoord{x, y})
		visited[GameVisibleCoord{x, y}] = true
		//oldLen := len(slice)

		//GameFieldVisible.cell[x][y] = 1

		//for firstRun == true || len(slice) > 0 {
		for len(visited) > 0 {
			//fmt.Println(visited)
			//fmt.Println("********************")
			//oldLen := len(slice)

			//if firstRun == false {

			// TODO: WTF ??? how extract first element (any) from map ??? Ugly !!!
			for k := range visited {
				//fmt.Println("First Element with loop", visited[k])
				x = k.x
				y = k.y
				break
			}

			//}

			//if GameFieldVisible.cell[x][y] == 0 {
			delete(visited, GameVisibleCoord{x, y})
			GameFieldVisible.cell[x][y] = 1

			for _, element := range ShiftCoordinate {
				//SetSurrounding(f, x, y, element.x, element.y)
				dx := x + element.x
				dy := y + element.y

				// Check boundaries
				// TODO: create func check boundaries
				if dx >= 0 && dy >= 0 &&
					dx < GameFieldWidth && dy < GameFieldHeight {

					//if GameField.cell[x][y] == 0 && IndexOf(slice, GameVisibleCoord{x: dx, y: dy}) == -1 {
					if GameField.cell[dx][dy] == 0 && GameFieldVisible.cell[dx][dy] == 0 {
						//slice = append(slice, GameVisibleCoord{dx, dy})
						visited[GameVisibleCoord{dx, dy}] = true
					}

					GameFieldVisible.cell[dx][dy] = 1
				}
			}

			//slice = SetSurroundingEmptyVisible(f, slice, realClick, x, y, element.x, element.y)

			//}

			//if IndexOf(slice, GameVisibleCoord{y: x, x: y}) > 0 {
			//	slice = slice[1:]
			//}

			//slice = slice[1:]

			//GameFieldVisible.cell[x][y] = 1

			//if len(slice) == oldLen {
			//	break
			//	return
			//}

			//firstRun = false
			//fmt.Println(visited)
			//Display(false)
			//return
		}
	}

	//if (GameField.cell[x][y] == 0 && realClick == true) || (realClick == false && GameFieldVisible.cell[x][y] == 0) {
	//	//if GameField.cell[x][y] == 0 && realClick == true {
	//	// show all empty cells
	//	//old_array:= [...]int
	//	GameFieldVisible.cell[x][y] = 1
	//
	//	//slice = append(slice, GameVisibleCoord{y, x})
	//	var slice = make([]GameVisibleCoord, 0)
	//	f := GameField.cell
	//
	//	for _, element := range ShiftCoordinate {
	//		//SetSurrounding(f, x, y, element.x, element.y)
	//		slice = SetSurroundingEmptyVisible(f, slice, realClick, x, y, element.x, element.y)
	//	}
	//
	//	fmt.Println(slice)
	//	//return
	//}

}

//func Bye() {
//	fmt.Println()
//	fmt.Println()
//	fmt.Println("Developed by MMM_Corp, test task special for Data Science UA, 2023")
//	fmt.Println("Skype: mmm_ogame")
//}

func main() {
	fmt.Println("*********************************")
	fmt.Println("*            PROXX              *")
	fmt.Println("*                               *")
	fmt.Println("* Legend:                       *")
	fmt.Println("*    H   - Black Hole           *")
	fmt.Println("*    0   - Visible Cell         *")
	fmt.Println("*        - Hidden Cell          *")
	fmt.Println("*    1-8 - Surrounding Cell     *")
	fmt.Println("*********************************")

	if //goland:noinspection ALL
	GameFieldBlackHoles > GameFieldWidth*GameFieldHeight {
		fmt.Printf("Error, possible maximum Black Holes Count: %v", GameFieldWidth*GameFieldHeight)
		os.Exit(0)
	}

	GameField = NewField(GameFieldBlackHoles) // Generated field
	GameFieldVisible = NewField(0)            // Visible Field

	fmt.Println("Generated Game Field")
	Display(true)
	//time.Sleep(1)
	fmt.Println("Simulate few clicks at random places")
	//os.Exit(0)

	for i := 0; i < GameFieldClicks; i++ {
		x := GetRandomInt(GameFieldWidth)
		y := GetRandomInt(GameFieldHeight)

		fmt.Println()
		fmt.Println("Clicked at: ", x, y)
		Click(x, y, true)
		Display(false)
	}

	//_ = 1
	//fmt.Println(GameFieldVisible)
	//defer Bye()
}
