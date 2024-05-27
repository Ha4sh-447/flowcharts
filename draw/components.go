package draw

import "fmt"

const (
	width    = 10
	height   = 3
	arrowLen = 4
)

func CenterX(c *Canvas, offset int) {
	c.Cursor.X = c.Cols/2 - 50 + offset
	fmt.Println("Centered ordinate x: ", c.Cursor.X)
}

func Box(content string, c *Canvas) {
	// width := len(content)
	// height := c.Cursor.Y + 3
	c.CenterX()
	size := len(content)

	c.Grid[c.Cursor.Y][c.Cursor.X] = "+"
	c.Grid[c.Cursor.Y+height][c.Cursor.X] = "+"
	c.Grid[c.Cursor.Y+height][c.Cursor.X+size+1] = "+"
	c.Grid[c.Cursor.Y][c.Cursor.X+size+1] = "+"

	for i := c.Cursor.X + 1; i < c.Cursor.X+size+1; i++ {
		c.Grid[c.Cursor.Y][i] = "-"
		c.Grid[c.Cursor.Y+height][i] = "-"
	}

	for s := 0; s < len(content); s++ {
		x := c.Cursor.X
		c.Grid[c.Cursor.Y+height/2][x+s+1] = string(content[s])
	}

	for y := 1; y < height; y++ {
		c.Grid[c.Cursor.Y+y][c.Cursor.X] = "|"
		c.Grid[c.Cursor.Y+y][c.Cursor.X+size+1] = "|"
	}

	c.Cursor.Y += height
	c.Cursor.X += (size) / 2
	fmt.Println("First box render post points: ", c.Cursor.X, c.Cursor.Y)
}

func BoxHorizontal(content string, c *Canvas) {
	size := len(content)

	fmt.Println("h Box - x: ", c.Cursor.X)
	fmt.Println("h Box - y: ", c.Cursor.Y)
	x := c.Cursor.X + 1
	y := c.Cursor.Y + 1

	c.Grid[y+height/2][x] = "+"
	c.Grid[y+height/2][x+size+1] = "+"
	c.Grid[y-height/2-1][x] = "+"
	c.Grid[y-height/2-1][x+size+1] = "+"

	for i := x + 1; i < x+size+1; i++ {
		c.Grid[y+height/2][i] = "-"
		c.Grid[y-height/2-1][i] = "-"
	}

	for s := 0; s < len(content); s++ {
		c.Grid[y-1][x+s+1] = string(content[s])
	}

	for i := y - height/2; i < y+height/2; i++ {
		c.Grid[i][x] = "|"
		c.Grid[i][x+size+1] = "|"
	}
	c.Cursor.Y = y + height/2
	offset := x + size/2
	fmt.Println(offset)

	CenterX(c, offset)
}

func DownArrow(content string, c *Canvas) {
	size := len(content)
	// 2 spaces -> content-> space-> arrow head

	for y := range arrowLen - 1 {
		c.Grid[c.Cursor.Y+1+y][c.Cursor.X] = "|"
	}

	c.Grid[c.Cursor.Y+arrowLen][c.Cursor.X] = "v"
	c.Cursor.Y += arrowLen + 1

	for s := range size {
		c.Grid[c.Cursor.Y-2][c.Cursor.X+1+s] = string(content[s])
	}
}

func RightArrow(content string, c *Canvas) {
	//take the pointer
	// go x+width/2 in right on x-axis (col)
	// go y-height/2 upwards on the y-axis (row)

	x := c.Cursor.X + (width / 2) + 1
	y := c.Cursor.Y - (height / 2) - 1
	// fmt.Println(c.Cursor.Y)
	// fmt.Println(y)

	for i := x; i < x+arrowLen; i++ {
		c.Grid[y][i+1] = "-"
	}

	c.Grid[y][x+arrowLen+1] = ">"

	c.Cursor.Y = y
	c.Cursor.X = x + arrowLen + 1
}

func LeftArrow(content string, c *Canvas) {
	x := c.Cursor.X - (width / 2) - 1
	y := c.Cursor.Y - (height / 2) - 1

	for i := x; i > x-arrowLen; i-- {
		c.Grid[y][i] = "-"
	}

	c.Grid[y][x-arrowLen] = "<"

	c.Cursor.Y = y
	c.Cursor.X = x - arrowLen
}
