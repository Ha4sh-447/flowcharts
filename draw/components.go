package draw

import "fmt"

const (
	width    = 7
	height   = 3
	arrowLen = 4
)

func CenterX(c *Canvas, offset int) {
	c.Cursor.X = (c.Cols / 2) - 10 + offset
	fmt.Println("Centered ordinate x: ", c.Cursor.X)
}

func Box(content string, c *Canvas) {
	// width := len(content)
	// height := c.Cursor.Y + 3
	c.CenterX()
	fmt.Println("V-BOX ", c.Cursor.X)
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

	fmt.Println("h Box - x: ", c.Cursor.X+1)
	fmt.Println("h Box - y: ", c.Cursor.Y+1)
	x := c.Cursor.X + 1
	y := c.Cursor.Y + 1
	updated_x := x + size + 1

	c.Grid[y+height/2][x] = "+"
	c.Grid[y+height/2][updated_x] = "+"
	c.Grid[y-height/2-1][x] = "+"
	c.Grid[y-height/2-1][updated_x] = "+"

	for i := x + 1; i < updated_x; i++ {
		c.Grid[y+height/2][i] = "-"
		c.Grid[y-height/2-1][i] = "-"
	}

	for s := 0; s < len(content); s++ {
		c.Grid[y-1][x+s+1] = string(content[s])
	}

	for i := y - height/2; i < y+height/2; i++ {
		c.Grid[i][x] = "|"
		c.Grid[i][updated_x] = "|"
	}
	c.Cursor.Y = y + height/2
	c.Cursor.X += size / 2
	offset := x + size
	fmt.Println("UP--> ", c.Cursor.Y)
	fmt.Println(offset)

	// CenterX(c, offset)
}

func BoxHorizontalLeft(content string, c *Canvas) {
	size := len(content)

	fmt.Println("h l Box - x: ", c.Cursor.X-1)
	fmt.Println("h l Box - y: ", c.Cursor.Y+1)
	x := c.Cursor.X - 1
	y := c.Cursor.Y + 1
	updated_x := x - size - 1

	fmt.Println("UPDATED X FOR LEFT RENDER DIR BOX: ", updated_x)

	c.Grid[y+height/2][x] = "+"
	c.Grid[y+height/2][updated_x] = "+"
	c.Grid[y-height/2-1][x] = "+"
	c.Grid[y-height/2-1][updated_x] = "+"

	for i := x - 1; i > updated_x; i-- {
		c.Grid[y+height/2][i] = "-"
		c.Grid[y-height/2-1][i] = "-"
	}

	for s := len(content) - 1; s >= 0; s-- {
		index := len(content) - s
		fmt.Println("UPDATED COORDS FOR CONTENT RENDER: ", y-1, updated_x+index)
		c.Grid[y-1][x-index] = string(content[s])
	}

	for i := y - height/2; i < y+height/2; i++ {
		c.Grid[i][x] = "|"
		c.Grid[i][updated_x] = "|"
	}

	c.Cursor.Y = y + height/2
	c.Cursor.X = x - size/2 - 1
	// offset := x - size
	// fmt.Println("UP--> ", c.Cursor.Y)
	fmt.Println(c.Cursor.X)

}

func DownArrow(content string, c *Canvas, pLen int) {
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
	CenterX(c, pLen)
}

func RightArrow(content string, c *Canvas) {
	//take the pointer
	// go x+width/2 in right on x-axis (col)
	// go y-height/2 upwards on the y-axis (row)

	x := c.Cursor.X + (width / 2) + 1
	y := c.Cursor.Y - (height / 2) - 1
	// fmt.Println("Right Arrow - ", c.Cursor.X)
	fmt.Println("Right Arrow updated ordinate - ", x)

	for i := x; i < x+arrowLen; i++ {
		c.Grid[y][i+1] = "-"
	}

	c.Grid[y][x+arrowLen+1] = ">"

	c.Cursor.Y = y
	c.Cursor.X = x + arrowLen + 1
	// fmt.Println(c.Cursor.X)
}

func LeftArrow(content string, c *Canvas) {
	x := c.Cursor.X - (width / 2)
	y := c.Cursor.Y - (height / 2) - 1

	fmt.Println("Left Arrow - ", c.Cursor.X)
	fmt.Println("Left Arrow updated ordinate - ", x)

	for i := x; i > x-arrowLen; i-- {
		c.Grid[y][i] = "-"
	}

	c.Grid[y][x-arrowLen] = "<"

	c.Cursor.Y = y
	c.Cursor.X = x - arrowLen
	fmt.Println(c.Cursor.X)
}
