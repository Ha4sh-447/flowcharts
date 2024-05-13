package flowchart

import (
	"fmt"
	"strings"
)

// Direction represents the direction of the flowchart
type Direction int

const (
	Vertical Direction = iota
	Horizontal
)

// ShapeType represents the type of shape in the flowchart
type ShapeType int

const (
	Rectangle ShapeType = iota
	Diamond
	Arrow
)

// Shape represents a shape in the flowchart
type Shape struct {
	Type    ShapeType
	Content string
}

// Flowchart represents a collection of shapes
type Flowchart struct {
	Shapes    []Shape
	Direction Direction
}

// NewFlowchart initializes a new empty flowchart
func NewFlowchart(direction Direction) *Flowchart {
	return &Flowchart{
		Direction: direction,
	}
}

// AddShape adds a shape to the flowchart
func (f *Flowchart) AddShape(shape Shape) {
	f.Shapes = append(f.Shapes, shape)
}

// AddArrow adds an arrow to the flowchart
func (f *Flowchart) AddArrow(content string) {
	f.Shapes = append(f.Shapes, Shape{Type: Arrow, Content: content})
}

// Render renders the flowchart in ASCII format
func (f *Flowchart) Render() {
	switch f.Direction {
	case Vertical:
		f.renderVertical()
	case Horizontal:
		f.renderHorizontal()
	}
}

func (f *Flowchart) renderVertical() {
	for _, shape := range f.Shapes {
		switch shape.Type {
		case Rectangle:
			fmt.Println("┌──────────────┐")
			fmt.Printf("│ %-12s │\n", shape.Content)
			fmt.Println("└──────────────┘")
		case Diamond:
			fmt.Println("    /\\")
			fmt.Println("   /  \\")
			fmt.Println("  /    \\")
			fmt.Printf(" /  %-4s  \\\n", shape.Content)
			fmt.Println("  \\    /")
			fmt.Println("   \\  /")
			fmt.Println("    \\/")
		case Arrow:
			fmt.Printf("     ↑\n")
			fmt.Printf("     │\n")
			fmt.Printf("     │\n")
			fmt.Printf("     │  %s\n", shape.Content)
			fmt.Printf("     │\n")
			fmt.Printf("     │\n")
			fmt.Printf("     ↓\n")
		}
	}
}

func (f *Flowchart) renderHorizontal() {
	// Calculate the maximum content length
	maxContentLen := 0
	for _, shape := range f.Shapes {
		if len(shape.Content) > maxContentLen {
			maxContentLen = len(shape.Content)
		}
	}

	// Render shapes
	for _, shape := range f.Shapes {
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
		case Arrow:
			fmt.Printf("     ↑  %s\n", shape.Content)
			fmt.Printf("     │\n")
			fmt.Printf("     │\n")
			fmt.Printf("     │\n")
			fmt.Printf("     ↓\n")
		}
	}
}
