package diagrams

import (
	// "fmt"
	"github.com/Ha4sh-447/diagramFromText/draw"
	// "strings"
)

type ShapeType int

const (
	Rectangle ShapeType = iota
	Diamond
	HArrow
	VArrow
	HLine
	VLine
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

	// Base condition
	if shape.Top == nil && shape.Left == nil && shape.Right == nil && shape.Bottom == nil && shape.IsRendered == false {
		shape.IsRendered = true
		RenderType(shape, c)
		// draw.CenterX(c)
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
		// draw.CenterX(c)
	} else if shape.Left != nil && shape.IsRendered == false {
		shape.IsRendered = true
		DRender(*shape.Left, c)
		// draw.CenterX(c)
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
	switch shape.Type {
	case Rectangle:
		draw.Box(shape.Content, c)
	case Diamond:
		Shape_daimond(shape.Content)
	case HArrow:
		HorizontalArrow(shape.Content)
	case HLine:
		HorizontalLine(shape.Content)
	case VArrow:
		VerticalArrow(shape.Content)
	case VLine:
		VerticalLine(shape.Content)
	case HRectangle:
		Horizontal_shape_rectangle(shape.Content)
	case LeftArrow:
		draw.LeftArrow(shape.Content, c)
	case RightArrow:
		draw.RightArrow(shape.Content, c)
	case DownArrow:
		draw.DownArrow(shape.Content, c)
	}

}
