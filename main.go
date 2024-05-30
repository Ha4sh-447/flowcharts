package main

import (
	"fmt"

	"github.com/Ha4sh-447/diagramFromText/diagrams"
	// "github.com/Ha4sh-447/diagramFromText/flowchart"
	"github.com/Ha4sh-447/diagramFromText/draw"
)

func main() {
	// diagramRenderer()

	canvas := draw.NewCanvas(20, 100)
	diagramRenderer(canvas)
	canvas.Render()

}

func canvaseDiag() {
	canvas := draw.NewCanvas(30, 100)
	draw.Box("foobar", canvas)
	draw.RightArrow("", canvas)
	draw.BoxHorizontal("foobar", canvas)
	// draw.Box("foobar", canvas)
	draw.RightArrow("", canvas)
	draw.BoxHorizontal("foobar", canvas)
	draw.RightArrow("", canvas)
	draw.BoxHorizontal("foobar", canvas)
	// draw.DownArrow("Yes", canvas)
	draw.Box("foobar", canvas)
	draw.LeftArrow("", canvas)
	// canvas.Render()
}

func diagramRenderer(c *draw.Canvas) {
	diagram := diagrams.New()

	// Creating Shapes
	parent := diagrams.Shape{
		Type:    diagrams.Rectangle,
		Content: "Start",
	}

	vArrow1 := diagrams.Shape{
		Type:    diagrams.DownArrow,
		Content: "1",
	}

	subNode1 := diagrams.Shape{
		Type:    diagrams.Rectangle,
		Content: "Sub111",
	}

	vArrow2 := diagrams.Shape{
		Type:    diagrams.DownArrow,
		Content: "2",
		PrevLen: 5 / 2,
	}

	subNode2 := diagrams.Shape{
		Type:    diagrams.Rectangle,
		Content: "Sub223",
	}

	subNode3 := diagrams.Shape{
		Type:    diagrams.HRectangle,
		Content: "Sub3",
	}

	hArrow3 := diagrams.Shape{
		Type:    diagrams.RightArrow,
		Content: "",
	}

	hArrow4 := diagrams.Shape{
		Type:    diagrams.RightArrow,
		Content: "",
	}

	subNode4 := diagrams.Shape{
		Type:    diagrams.HRectangle,
		Content: "Sub4",
	}

	subNode5 := diagrams.Shape{
		Type:      diagrams.HRectangle,
		Content:   "Sub5",
		RenderDir: "left",
	}

	hArrow5 := diagrams.Shape{
		Type:    diagrams.LeftArrow,
		Content: "",
	}

	hArrow6 := diagrams.Shape{
		Type:    diagrams.LeftArrow,
		Content: "",
	}

	subNode6 := diagrams.Shape{
		Type:      diagrams.HRectangle,
		Content:   "Sub6",
		RenderDir: "left",
	}

	fmt.Println(hArrow4, subNode5)

	diagrams.AddToBottom(&subNode1, &vArrow1)
	diagrams.AddToBottom(&parent, &vArrow2)
	// diagrams.AddToRight(&parent, &hArrow3)
	// diagrams.AddToRight(&hArrow3, &subNode3)

	diagrams.AddToRight(&parent, &hArrow3)
	diagrams.AddToRight(&subNode3, &hArrow4)
	diagrams.AddToRight(&hArrow4, &subNode4)
	diagrams.AddToRight(&hArrow3, &subNode3)

	diagrams.AddToLeft(&parent, &hArrow5)
	diagrams.AddToLeft(&hArrow5, &subNode5)
	diagrams.AddToLeft(&hArrow6, &subNode6)
	diagrams.AddToLeft(&subNode5, &hArrow6)
	// Same fking problem!!!
	// can render linear diagrams but
	// What for non linear diagrams with nodes branching from sides
	diagram.AddShapes(subNode1)
	diagram.AddShapes(vArrow1)
	diagram.AddShapes(parent)
	diagram.AddShapes(hArrow5)
	diagram.AddShapes(subNode5)
	diagram.AddShapes(hArrow6)
	diagram.AddShapes(subNode6)
	diagram.AddShapes(hArrow3)
	diagram.AddShapes(subNode3)
	diagram.AddShapes(hArrow4)
	diagram.AddShapes(subNode4)
	diagram.AddShapes(vArrow2)
	diagram.AddShapes(subNode2)

	fmt.Println("LOGS FROM MAIN.GO-diagramRenderer")
	for _, shape := range diagram.Shapes {
		fmt.Printf("%d ---> %s\n", shape.Type, shape.Content)
	}

	// diagram.Render()
	for _, shape := range diagram.Shapes {
		// fmt.Println(shape.Type)
		diagrams.DRender(shape, c)
	}

	// diagrams.Horizontal_shape_rectangle("foobar")
	// fmt.Println()
	// diagrams.HorizontalArrow("foobar")
	// fmt.Println()

	// diagram.RenderHorizontal()
}
