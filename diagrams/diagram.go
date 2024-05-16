package diagrams

import (
	"fmt"
	"strings"
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
)

type Shape struct {
	Type       ShapeType
	Content    string
	Left       *Shape
	Right      *Shape
	Top        *Shape
	Bottom     *Shape
	IsRendered bool
}

type Diagram struct {
	Shapes []Shape
}

func New() *Diagram {
	return &Diagram{}
}

func NewShape() {

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

func DRender(shape Shape) {

	// Base condition
	if shape.Top == nil && shape.Left == nil && shape.Right == nil && shape.Bottom == nil && shape.IsRendered == false {
		shape.IsRendered = true
		RenderType(shape)
		return
	}

	// Render initial shape
	if !shape.IsRendered {
		shape.IsRendered = true
		RenderType(shape)
	}

	if shape.Right != nil && shape.IsRendered == false {
		shape.IsRendered = true
		DRender(*shape.Right)
	} else if shape.Left != nil && shape.IsRendered == false {
		shape.IsRendered = true
		DRender(*shape.Left)
	} else if shape.Bottom != nil && shape.IsRendered == false {
		shape.IsRendered = true
		DRender(*shape.Bottom)
	} else if shape.Top != nil && shape.IsRendered == false {
		shape.IsRendered = true
		DRender(*shape.Top)
	}
}

func RenderHorizontal(diagram []string) {
	// Find the maximum length of all elements in the diagram
	maxLength := 0
	for _, element := range diagram {
		if len(element) > maxLength {
			maxLength = len(element)
		}
	}

	// Print each element side by side with padding to match the maxLength
	for _, element := range diagram {
		padding := maxLength - len(element)
		fmt.Printf("%s%s", element, strings.Repeat(" ", padding))
	}
	fmt.Println() // Print newline after rendering the horizontal line
}

func (d *Diagram) RenderHorizontal() {
	// Calculate the maximum content length
	maxContentLen := 0
	for _, shape := range d.Shapes {
		if len(shape.Content) > maxContentLen {
			maxContentLen = len(shape.Content)
		}
	}

	// Render shapes
	for _, shape := range d.Shapes {
		switch shape.Type {
		case Rectangle:
			fmt.Printf("┌%s┐\n", strings.Repeat("─", maxContentLen+2))
			fmt.Printf("│ %-*s │\n", maxContentLen, shape.Content)
			fmt.Printf("└%s┘\n", strings.Repeat("─", maxContentLen+2))
		case Diamond:
			fmt.Println("    /\\")
			fmt.Println("   /  \\")
			fmt.Println("  /    \\")
			fmt.Printf(" /  %-4s  \\\n", shape.Content)
			fmt.Println("  \\    /")
			fmt.Println("   \\  /")
			fmt.Println("    \\/")

		}
	}
}

// Render Shape based upon it's type
func RenderType(shape Shape) {
	switch shape.Type {
	case Rectangle:
		Shape_rectangle(shape.Content)
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
	}

}
