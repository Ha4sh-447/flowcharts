package diagrams

import (
	"github.com/Ha4sh-447/flowchartFromText/draw"
)

type ShapeType int

// YES ONLY REACTANGLES AND ARROW FOR NOW
const (
	Rectangle ShapeType = iota
	HRectangle
	LeftArrow
	RightArrow
	DownArrow
)

type Shape struct {
	// Type: Shows the type of shape, based on the ShapeType type
	Type ShapeType

	// Content: String type
	// Stores the content to be displayed
	Content string

	// Left: Contains a reference of the shape connected to left of the current shape
	Left *Shape

	// Right: Contains a reference of the shape connected to right of the current shape
	Right *Shape

	// Top: Contains a reference of the shape connected to top of the current shape
	Top *Shape

	// Bottom: Contains a reference of the shape connected to bottom of the current shape
	Bottom *Shape

	// IsRendered: Boolean value indicating whether the shape is rendered or not
	IsRendered bool

	// IsHorizontal: If the shape is to be rendered in horizontal pipeline then this variable is set to true
	IsHorizontal bool

	// PrevLen: Horizontal space occupied by the previous shape
	PrevLen int

	// RenderDir: Similar to IsHorizontal, but also considers vertical render direction
	RenderDir string

	// IsLast: Set to true if that shape is the last one to be renderd in that particular direction
	IsLast bool

	// IsJunction: If a shape has more than 2 shapes connected to it, then it is set to true
	IsJunction bool

	// MidX: Stores the center of the shape on the X axis
	MidX int

	// MidY: Stores the center of the shape on the Y axis
	MidY int
}

// List of shape
type Diagram struct {
	S []Shape
}

// Array to store shapes which are a junction
type Store struct {
	stack []*Shape
}

// Returns a new instance of Diagram
func New() *Diagram {
	return &Diagram{}
}

// Returns a new instance of Store
func NewStore() *Store {
	return &Store{}
}

// Add Shapes to the Diagram struct
func (d *Diagram) AddShapes(shape Shape) {
	d.S = append(d.S, shape)
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

// Finds a node/shape
func findNode(shape Shape, d []Shape) *Shape {
	for _, s := range d {
		if shape == s {
			return &s
		}
	}
	return nil
}

// Adds Shape with IsJunction attribute set to true to stack
func (s *Store) Junction(shape *Shape, c *draw.Canvas) {

	shape.MidX = c.Cursor.X
	shape.MidY = c.Cursor.Y
	s.stack = append(s.stack, shape)
}

// Get shape from stack
func (s *Store) getJunction() *Shape {
	return s.stack[len(s.stack)-1]
}

// Renders the shape
func RenderD(shape *Shape, c *draw.Canvas, s *Store) {

	if shape.IsRendered == false {
		shape.IsRendered = true
		RenderType(*shape, c)
		if shape.IsJunction {
			s.Junction(shape, c)
		}
	}

	if shape.IsLast == true {
		_s := s.getJunction()
		if _s.Top != nil {
			// _s.MidX = _s.MidX + 1
		}
		draw.Center(c, _s.MidX, _s.MidY)
		return
	}

	if shape.Left != nil && shape.IsRendered == false {
		RenderD(shape.Left, c, s)
	}
	if shape.Right != nil && shape.IsRendered == false {
		RenderD(shape.Right, c, s)
	}
	if shape.Bottom != nil && shape.IsRendered == false {
		RenderD(shape.Bottom, c, s)
	}
}

// Render Shape based upon it's type
func RenderType(shape Shape, c *draw.Canvas) {

	switch shape.Type {
	case Rectangle:
		BoX(shape, c)
	case HRectangle:
		BoX(shape, c)
	case LeftArrow:
		Arrow(&shape, c)
	case RightArrow:
		Arrow(&shape, c)
	case DownArrow:
		Arrow(&shape, c)
	}

}
