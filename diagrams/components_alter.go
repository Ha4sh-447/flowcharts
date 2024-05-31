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

	// NO it won't be top left
	// it will center of the top border

	fmt.Println("BOX - x ", x)
	fmt.Println("BOX - y ", y)

	fmt.Println("BOX - UPDATED x ", x-len(s.Content)/2, x+len(s.Content)/2)

	width := len(s.Content)

	c.Grid[y][x] = "+"
	c.Grid[y+HEIGHT][x] = "+"
	c.Grid[y][x+width+1] = "+"
	c.Grid[y+HEIGHT][x+width+1] = "+"

	for i := x + 1; i < x+width+1; i++ {
		c.Grid[y][i] = "-"
		c.Grid[y+HEIGHT][i] = "-"
	}

	index := 0
	for i := x + 1; x < x+width+1 && index < len(s.Content); i++ {
		c.Grid[y+HEIGHT/2][i] = string(s.Content[index])
		index++
	}

	for i := y + 1; i < y+HEIGHT; i++ {
		c.Grid[i][x] = "|"
		c.Grid[i][x+width+1] = "|"
	}

	c.Cursor.X = x + (width+1)/2
	c.Cursor.Y = y + HEIGHT

	// UPDATE COORDs FOR NEXT RENDER
	// SETTING CURSOR TO BOTTOM RIGHT
	// c.Cursor.Y = y + HEIGHT
	/*
		if(s.RenderDir=="left"){
			c.Cursor.Y = y+HEIGHT/2
		}else if(s.RenderDir=="right"){
			c.Cursor.X = x+width+1
			c.Cursor.Y = y+HEIGHT/2
		}else{
			c.Cursor.X = x+(width+1)/2
			c.Cursor.Y = y+c.Cursor.Y
		}
	*/
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
		x = c.Cursor.X - (len(s.Right.Content) / 2) - 1
		y = c.Cursor.Y - (HEIGHT)/2

		for i := x - 1; i > x-ARROWLEN-1; i-- {
			c.Grid[y][i] = "-"
		}

		c.Grid[y][x-ARROWLEN-1] = "<"

		c.Cursor.X = x - ARROWLEN - 1
		c.Cursor.Y = y
	} else if s.Type == RightArrow {
		x = c.Cursor.X + (len(s.Left.Content) / 2)
		y = c.Cursor.Y - HEIGHT/2

		for i := x + 1; i < x+ARROWLEN+1; i++ {
			c.Grid[i][y] = "-"
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
	}

	fmt.Println("ARROW- final X ", x)
	fmt.Println("ARROW - final - Y ", y)

}
