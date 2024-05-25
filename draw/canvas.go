package draw

import "fmt"

// Canvas grid will be the draw area
type Canvas struct {
	Rows   int
	Cols   int
	Grid   [][]string
	Cursor Point
}

type Point struct {
	X int
	Y int
}

// Create a new canvas
func NewCanvas(r, c int) *Canvas {
	g := make([][]string, r)
	for i := range g {
		g[i] = make([]string, c)
	}
	for i := range g {
		for j := range g[i] {
			g[i][j] = "."
		}
	}

	p := Point{
		X: c/2 - 10,
		Y: 0,
	}

	return &Canvas{
		Rows:   r,
		Cols:   c,
		Grid:   g,
		Cursor: p,
	}
}

func (c *Canvas) CenterX() {
	c.Cursor.X = c.Cols/2 - 10
}

func (c *Canvas) Render() {
	for i := range c.Grid {
		for j := range c.Grid[0] {
			fmt.Printf("%s", c.Grid[i][j])
		}
		fmt.Println()
	}
}

// When drawing on canvas
// Draw the first shape in center
func (c *Canvas) Draw() {

}
