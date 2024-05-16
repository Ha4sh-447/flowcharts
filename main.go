package main

import (
	"fmt"

	"github.com/Ha4sh-447/diagramFromText/diagrams"
	// "github.com/Ha4sh-447/diagramFromText/flowchart"
)

func main() {
	diagramRenderer()
}

func diagramRenderer() {
	diagram := diagrams.New()

	// Creating Shapes
	parent := diagrams.Shape{
		Type:    diagrams.Rectangle,
		Content: "Start",
	}

	vArrow1 := diagrams.Shape{
		Type:    diagrams.VArrow,
		Content: "1",
	}

	subNode1 := diagrams.Shape{
		Type:    diagrams.Rectangle,
		Content: "Sub1",
	}

	vArrow2 := diagrams.Shape{
		Type:    diagrams.VArrow,
		Content: "2",
	}

	subNode2 := diagrams.Shape{
		Type:    diagrams.Rectangle,
		Content: "Sub2",
	}

	diagrams.AddToBottom(&subNode1, &vArrow1)
	diagrams.AddToBottom(&parent, &vArrow2)
	// diagrams.AddToRight(&parent, &hArrow3)
	// diagrams.AddToRight(&hArrow3, &subNode3)

	// Same fking problem!!!
	// can render linear diagrams but
	// What for non linear diagrams with nodes branching from sides
	diagram.AddShapes(subNode1)
	diagram.AddShapes(vArrow1)
	diagram.AddShapes(parent)
	diagram.AddShapes(vArrow2)
	diagram.AddShapes(subNode2)

	for _, shape := range diagram.Shapes {
		fmt.Println(shape)
	}

	// diagram.Render()
	for _, shape := range diagram.Shapes {
		// fmt.Println(shape.Type)
		diagrams.DRender(shape)
	}

	diagrams.Horizontal_shape_rectangle("foobar")
	// fmt.Println()
	diagrams.HorizontalArrow("foobar")

	// diagram.RenderHorizontal()
}
