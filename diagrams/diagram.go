package diagrams

import (
	"fmt"
	"github.com/Ha4sh-447/diagramFromText/draw"
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
	IsLast       bool
	IsJunction   bool
	MidX         int
	MidY         int
}

type Diagram struct {
	S []Shape
}

type Store struct {
	queue []*Shape
}

func New() *Diagram {
	return &Diagram{}
}

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

func findNode(shape Shape, d []Shape) *Shape {
	for _, s := range d {
		if shape == s {
			return &s
		}
	}
	return nil
}

func (s *Store) Junction(shape *Shape, c *draw.Canvas) {

	shape.MidX = c.Cursor.X
	shape.MidY = c.Cursor.Y
	s.queue = append(s.queue, shape)
	fmt.Println("\nSHAPE ADDED TO THE QUEUE:  ", shape.MidX, shape.Content, shape.MidY)
}

func (s *Store) getJunction() *Shape {
	fmt.Println(s.queue[len(s.queue)-1])
	return s.queue[len(s.queue)-1]
}

func RenderD(shape *Shape, c *draw.Canvas, s *Store) {

	if shape.IsRendered == false {
		shape.IsRendered = true
		RenderType(*shape, c)
		if shape.IsJunction {
			s.Junction(shape, c)
		}
	}

	if shape.IsLast == true {
		fmt.Println("RenderD--")
		fmt.Println(shape.Content)
		// shape.IsRendered = true
		// RenderType(*shape, c)
		_s := s.getJunction()
		draw.Center(c, _s.MidX, _s.MidY)
		return
	}

	if shape.Left != nil && shape.IsRendered == false {
		// shape.IsRendered = true
		// RenderType(*shape.Left, c)
		RenderD(shape.Left, c, s)
	}
	if shape.Right != nil && shape.IsRendered == false {
		// shape.IsRendered = true
		RenderD(shape.Right, c, s)
	}
	if shape.Bottom != nil && shape.IsRendered == false {
		// shape.IsRendered = true
		RenderD(shape.Bottom, c, s)
	}
	if shape.Top != nil && shape.IsRendered == false {
		// shape.IsRendered = true
		RenderD(shape.Top, c, s)
	}

}

// Render Shape based upon it's type
func RenderType(shape Shape, c *draw.Canvas) {

	switch shape.Type {
	case Rectangle:
		// draw.Box(shape.Content, c)
		BoX(shape, c)
	case HRectangle:
		BoX(shape, c)
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
