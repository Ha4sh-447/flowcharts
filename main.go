package main

import (
	"fmt"

	"github.com/Ha4sh-447/diagramFromText/diagrams"
	// "github.com/Ha4sh-447/diagramFromText/flowchart"
	"github.com/Ha4sh-447/diagramFromText/draw"
)

func main() {
	// diagramRenderer()
	subNode3 := diagrams.Shape{
		Type:    diagrams.HRectangle,
		Content: "SubZeroMygawwwwwwdddd3",
	}

	hArrow4 := diagrams.Shape{
		Type:    diagrams.DownArrow,
		Content: "",
	}

	subNode4 := diagrams.Shape{
		Type:    diagrams.HRectangle,
		Content: "SubOne4",
	}

	hArrow5 := diagrams.Shape{
		Type:    diagrams.DownArrow,
		Content: "",
	}

	subNode5 := diagrams.Shape{
		Type:    diagrams.HRectangle,
		Content: "SubOne5",
	}

	hArrow6 := diagrams.Shape{
		Type:    diagrams.LeftArrow,
		Content: "",
	}

	subNode6 := diagrams.Shape{
		Type:    diagrams.HRectangle,
		Content: "SubOne666",
	}

	subNode7 := diagrams.Shape{
		Type:    diagrams.HRectangle,
		Content: "Sub",
	}

	hArrow7 := diagrams.Shape{
		Type:    diagrams.LeftArrow,
		Content: "",
	}

	hArrow8 := diagrams.Shape{
		Type:    diagrams.RightArrow,
		Content: "",
	}

	subNode8 := diagrams.Shape{
		Type:    diagrams.HRectangle,
		Content: "SubRight",
	}

	fmt.Println(hArrow8, subNode8)
	canvas := draw.NewCanvas(30, 100)

	diagrams.AddToLeft(&subNode4, &hArrow5)
	diagrams.AddToRight(&hArrow5, &subNode4)
	diagrams.AddToRight(&subNode5, &hArrow5)
	diagrams.AddToLeft(&subNode5, &hArrow6)
	diagrams.AddToRight(&hArrow6, &subNode5)
	diagrams.AddToLeft(&hArrow6, &subNode6)
	diagrams.AddToRight(&subNode6, &hArrow6)
	diagrams.AddToLeft(&subNode6, &hArrow7)
	diagrams.AddToLeft(&hArrow7, &subNode7)
	diagrams.AddToRight(&subNode7, &hArrow7)
	diagrams.AddToRight(&hArrow7, &subNode6)

	diagrams.AddToRight(&subNode5, &hArrow8)
	diagrams.AddToLeft(&hArrow8, &subNode5)

	diagrams.AddToRight(&hArrow8, &subNode8)
	diagrams.AddToLeft(&subNode8, &hArrow8)

	diagrams.AddToRight(&subNode8, &hArrow8)
	diagrams.AddToLeft(&hArrow8, &subNode8)

	diagrams.BoX(subNode3, canvas)
	diagrams.Arrow(&hArrow4, canvas)
	diagrams.BoX(subNode4, canvas)
	diagrams.Arrow(&hArrow5, canvas)
	diagrams.BoX(subNode5, canvas)
	diagrams.Arrow(&hArrow8, canvas)
	diagrams.BoX(subNode8, canvas)
	diagrams.Arrow(&hArrow8, canvas)
	diagrams.BoX(subNode8, canvas)

	/*
		diagrams.Arrow(&hArrow6, canvas)
		diagrams.BoX(subNode6, canvas)
		diagrams.Arrow(&hArrow7, canvas)
		diagrams.BoX(subNode7, canvas)
	*/
	// diagramRenderer(canvas)
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
		Content: "Subbbbb4",
	}

	subNode5 := diagrams.Shape{
		Type:      diagrams.HRectangle,
		Content:   "Subvvv5",
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

	// diagrams.Arrow(&hArrow6, c)

	fmt.Println(hArrow4, subNode5)

	diagrams.AddToBottom(&subNode1, &vArrow1)
	diagrams.AddToTop(&vArrow1, &subNode1)
	diagrams.AddToBottom(&parent, &vArrow2)
	diagrams.AddToTop(&vArrow2, &parent)
	// diagrams.AddToRight(&parent, &hArrow3)
	// diagrams.AddToRight(&hArrow3, &subNode3)

	diagrams.AddToRight(&parent, &hArrow3)
	diagrams.AddToLeft(&hArrow3, &parent)
	diagrams.AddToRight(&subNode3, &hArrow4)
	diagrams.AddToLeft(&hArrow4, &subNode3)
	diagrams.AddToRight(&hArrow4, &subNode4)
	diagrams.AddToLeft(&subNode4, &hArrow4)
	diagrams.AddToRight(&hArrow3, &subNode3)
	diagrams.AddToLeft(&subNode3, &hArrow3)

	diagrams.AddToLeft(&parent, &hArrow5)
	diagrams.AddToRight(&hArrow5, &parent)
	diagrams.AddToLeft(&hArrow5, &subNode5)
	diagrams.AddToRight(&subNode5, &hArrow5)
	diagrams.AddToLeft(&hArrow6, &subNode6)
	diagrams.AddToRight(&subNode6, &hArrow6)
	diagrams.AddToLeft(&subNode5, &hArrow6)
	diagrams.AddToRight(&hArrow6, &subNode5)
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

	// diagram.Render()
	for _, shape := range diagram.Shapes {
		// fmt.Println(shape)
		diagrams.DRender(shape, c)
	}

	// diagrams.Horizontal_shape_rectangle("foobar")
	// fmt.Println()
	// diagrams.HorizontalArrow("foobar")
	// fmt.Println()

	// diagram.RenderHorizontal()
}
