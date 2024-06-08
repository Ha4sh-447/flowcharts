package diagrams

import (
	"fmt"

	"github.com/Ha4sh-447/diagramFromText/draw"
)

const HEIGHT = 3
const ARROWLEN = 3

func BoX(s Shape, c *draw.Canvas) {
	// X and Y will always be the topleft corner of the box
	x := c.Cursor.X
	y := c.Cursor.Y
	width := len(s.Content)

	// NO it won't be top left
	// it will center of the top border

	//After rendering the updated point of Cursor will be
	// as the middle of bottom border
	fmt.Println("BOX - x ", x)
	fmt.Println("BOX - y ", y)

	fmt.Println("CUUNTENT LENINTHINOOO: ", len(s.Content))

	fmt.Println("BOX - UPDATED x ", x-len(s.Content)/2, x+len(s.Content)/2, y)

	if s.Right != nil && s.Right.Type == LeftArrow {
		// x = x+width +1
		// y -> y+HEIGHT/2 && y -> y-HEIGHT/2
		c.Grid[y-HEIGHT/2-1][x-1] = "+"
		c.Grid[y+HEIGHT/2][x-1] = "+"
		c.Grid[y+HEIGHT/2][x-width-2] = "+"
		c.Grid[y-HEIGHT/2-1][x-width-2] = "+"

		for i := x - width - 1; i < x-1; i++ {
			c.Grid[y-HEIGHT/2-1][i] = "-"
			c.Grid[y+HEIGHT/2][i] = "-"
		}

		index := 0
		for i := x - width - 1; i < x-1 && index < width; i++ {
			c.Grid[y-1][i] = string(s.Content[index])
			index++
		}

		for i := y - HEIGHT/2; i < y+HEIGHT/2; i++ {
			c.Grid[i][x-1] = "|"
			c.Grid[i][x-width-2] = "|"
		}

		if len(s.Content)/2 != 0 {
			c.Cursor.X = x - len(s.Content)/2 - 2
		} else {
			c.Cursor.X = x - len(s.Content)/2
		}

		// c.Cursor.X = x - len(s.Content)/2
		c.Cursor.Y = y + HEIGHT/2
	} else if s.Left != nil && s.Left.Type == RightArrow {
		//x -> x+1 && x+width+1
		//y -> y+HEIGHT/2 && y-HEIGHT/2

		c.Grid[y+HEIGHT/2][x+1] = "+"
		c.Grid[y+HEIGHT/2][x+width+2] = "+"
		c.Grid[y-HEIGHT/2-1][x+1] = "+"
		c.Grid[y-HEIGHT/2-1][x+width+2] = "+"

		for i := x + 2; i < x+width+2; i++ {
			c.Grid[y+HEIGHT/2][i] = "-"
			c.Grid[y-HEIGHT/2-1][i] = "-"
		}

		index := 0
		for i := x + 2; i < x+width+2; i++ {
			c.Grid[y-1][i] = string(s.Content[index])
			index++
		}

		for i := y - HEIGHT/2; i < y+HEIGHT/2; i++ {
			c.Grid[i][x+1] = "|"
			c.Grid[i][x+width+2] = "|"
		}

		if len(s.Content)/2 != 0 {
			c.Cursor.X = x + len(s.Content)/2 + 2
		} else {
			c.Cursor.X = x + len(s.Content)/2
		}
		c.Cursor.Y = y + HEIGHT/2

	} else {

		c.Grid[y][x-len(s.Content)/2-1] = "+"
		c.Grid[y+HEIGHT][x+len(s.Content)/2+1] = "+"
		c.Grid[y][x+len(s.Content)/2+1] = "+"
		c.Grid[y+HEIGHT][x-len(s.Content)/2-1] = "+"

		for i := x - len(s.Content)/2; i < x+len(s.Content)/2+1; i++ {
			c.Grid[y][i] = "-"
			c.Grid[y+HEIGHT][i] = "-"
		}

		index := 0
		for i := x - len(s.Content)/2; x < x+len(s.Content)/2+1 && index < len(s.Content); i++ {
			c.Grid[y+HEIGHT/2][i] = string(s.Content[index])
			index++
		}

		for i := y + 1; i < y+HEIGHT; i++ {
			c.Grid[i][x-len(s.Content)/2-1] = "|"
			c.Grid[i][x+len(s.Content)/2+1] = "|"
		}

		// c.Cursor.X = x + (width+1)/2
		c.Cursor.Y = y + HEIGHT
	}
}

// What this function does is : simple - render an arrow based upon the render direction
// takes the content length of the previous shape so that it can adjust itself to the center of
// the prev shape
func Arrow(s *Shape, c *draw.Canvas) {
	x := c.Cursor.X
	y := c.Cursor.Y
	fmt.Println("ARROW- init X ", x)
	fmt.Println("ARROW - init - Y ", y)

	if s.Type == LeftArrow {
		fmt.Println("LEFT")
		// CURSOR is at the center of the bottom of previous shape
		x = c.Cursor.X - (len(s.Right.Content) / 2) - 1
		y = c.Cursor.Y - (HEIGHT)/2

		for i := x - 1; i > x-ARROWLEN-1; i-- {
			c.Grid[y][i] = "-"
		}

		c.Grid[y][x-ARROWLEN-1] = "<"

		c.Cursor.X = x - ARROWLEN - 1
		c.Cursor.Y = y
	} else if s.Type == RightArrow {
		fmt.Println("RIGHT")
		x = c.Cursor.X + (len(s.Left.Content) / 2)
		y = c.Cursor.Y - HEIGHT/2

		for i := x + 1; i < x+ARROWLEN+1; i++ {
			c.Grid[y][i] = "-"
		}
		c.Grid[y][x+ARROWLEN+1] = ">"

		c.Cursor.X = x + ARROWLEN + 1
		c.Cursor.Y = y
	} else {
		for i := y + 1; i < y+ARROWLEN+1; i++ {
			c.Grid[i][x] = "|"
		}
		c.Grid[y+ARROWLEN+1][x] = "v"

		c.Cursor.Y = y + ARROWLEN + 2

		// Setting the center of the canvas according to the diagram
		// using the down arrow as the reference
		c.Center = c.Cursor.X
	}

	fmt.Println("ARROW- final X ", x)
	fmt.Println("ARROW - final - Y ", y)

}
