package diagrams

import (
	// "fmt"
	"fmt"

	"github.com/Ha4sh-447/diagramFromText/draw"
	// "strings"
)

type ShapeType int

const (
	Rectangle ShapeType = iota
	Diamond
	HRectangle
	LeftArrow
	RightArrow
	DownArrow
)

type Shape struct {
	Type         ShapeType
	Content      string
	Left         *Shape
	Right        *Shape
	Top          *Shape
	Bottom       *Shape
	IsRendered   bool
	IsHorizontal bool
	PrevLen      int
	RenderDir    string
}

type Diagram struct {
	Shapes []Shape
}

func New() *Diagram {
	return &Diagram{}
}

// Add Shapes to the Diagram struct
func (d *Diagram) AddShapes(shape Shape) {
	d.Shapes = append(d.Shapes, shape)
}

// Add Shape to the right of the node
func AddToRight(shape *Shape, subShape *Shape) Shape {
	// Adds shape to right of the parent shape
	shape.Right = subShape
	return *shape
}

// Add Shape to the left of the node
func AddToLeft(shape *Shape, subShape *Shape) Shape {
	// Adds shape to right of the parent shape
	shape.Left = subShape
	return *shape
}

// Add Shape to the bottom of the node
func AddToBottom(shape *Shape, subShape *Shape) Shape {
	// Adds shape to right of the parent shape
	shape.Bottom = subShape
	return *shape
}

// Add Shape to the top of the node
func AddToTop(shape *Shape, subShape *Shape) Shape {
	// Adds shape to right of the parent shape
	shape.Top = subShape
	return *shape
}

func DRender(shape Shape, c *draw.Canvas) {

	fmt.Println()
	// Base condition
	if shape.Top == nil && shape.Left == nil && shape.Right == nil && shape.Bottom == nil && shape.IsRendered == false {
		shape.IsRendered = true
		// fmt.Println("DRENDER- ", shape.Content, shape.Right)
		RenderType(shape, c)
		draw.CenterX(c, len(shape.Content)/2)
		return
	}

	// Render initial shape
	if !shape.IsRendered {
		shape.IsRendered = true
		RenderType(shape, c)
	}

	if shape.Right != nil && shape.IsRendered == false {
		shape.IsRendered = true
		DRender(*shape.Right, c)
		draw.CenterX(c, (len(shape.Content)+2)/2)
	} else if shape.Left != nil && shape.IsRendered == false {
		shape.IsRendered = true
		shape.Left.RenderDir = "left"
		DRender(*shape.Left, c)
		// draw.CenterX(c, len(shape.Content)/2)
	} else if shape.Bottom != nil && shape.IsRendered == false {
		shape.IsRendered = true
		DRender(*shape.Bottom, c)
	} else if shape.Top != nil && shape.IsRendered == false {
		shape.IsRendered = true
		DRender(*shape.Top, c)
	}
}

// Render Shape based upon it's type
func RenderType(shape Shape, c *draw.Canvas) {
	if shape.Type == HRectangle && shape.RenderDir == "left" {
		draw.BoxHorizontalLeft(shape.Content, c)
		return
	}
	switch shape.Type {
	case Rectangle:
		// draw.Box(shape.Content, c)
		BoX(shape, c)
	case Diamond:
		Shape_daimond(shape.Content)
	case HRectangle:
		draw.BoxHorizontal(shape.Content, c)
	case LeftArrow:
		// draw.LeftArrow(shape.Content, c)
		Arrow(&shape, c)
	case RightArrow:
		// draw.RightArrow(shape.Content, c)
		Arrow(&shape, c)
	case DownArrow:
		// draw.DownArrow(shape.Content, c, shape.PrevLen)
		Arrow(&shape, c)
	}

}
