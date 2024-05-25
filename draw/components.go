package draw

func Box(content string, c *Canvas) {
	width := len(content)
	height := c.Cursor.Y + 3

	c.Grid[c.Cursor.Y][c.Cursor.X] = "+"
	c.Grid[c.Cursor.Y+height][c.Cursor.X] = "+"
	c.Grid[c.Cursor.Y+height][c.Cursor.X+width+2] = "+"
	c.Grid[c.Cursor.Y][c.Cursor.X+width+2] = "+"

	for i := c.Cursor.X + 1; i <= c.Cursor.X+width+1; i++ {
		c.Grid[c.Cursor.Y][i] = "-"
		c.Grid[height][i] = "-"
	}

	for s := 0; s < width; s++ {
		x := c.Cursor.X
		c.Grid[height/2][x+s+2] = string(content[s])
	}

	for y := c.Cursor.Y + 1; y < height; y++ {
		c.Grid[y][c.Cursor.X] = "|"
		c.Grid[y][c.Cursor.X+width+2] = "|"
	}

	c.Cursor.Y += height
	c.Cursor.X += (width + 2) / 2
}

func DownArrow(content string, c *Canvas) {
	size := len(content)
	arrowLen := 4
	// 2 spaces -> content-> space-> arrow head

	for y := range arrowLen - 1 {
		c.Grid[c.Cursor.Y+1+y][c.Cursor.X] = "|"
	}

	c.Grid[c.Cursor.Y+arrowLen][c.Cursor.X] = "v"
	c.Cursor.Y += arrowLen

	for s := range size {
		c.Grid[c.Cursor.Y-2][c.Cursor.X+1+s] = string(content[s])
	}
}
