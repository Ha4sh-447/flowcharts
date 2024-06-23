package main

import (
	"github.com/Ha4sh-447/flowcharts/diagrams"
	"github.com/Ha4sh-447/flowcharts/draw"
)

func main() {
	// Large example
	example1()

	// Small example
	example2()
}

func example1() {

	//Create canvas of desired size
	canvas := draw.NewCanvas(30, 100)

	// Describe your shapes/nodes with type and content
	diagram := diagrams.New()

	subNode3 := diagrams.Shape{
		Type:    diagrams.HRectangle,
		Content: "SubZero",
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

		IsJunction: true,
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

		IsLast: true,
	}

	hArrow7 := diagrams.Shape{
		Type:    diagrams.LeftArrow,
		Content: "",
	}

	hArrow8 := diagrams.Shape{
		Type: diagrams.RightArrow,
	}

	subNode8 := diagrams.Shape{
		Type:    diagrams.HRectangle,
		Content: "SubRight",

		IsJunction: true,
		// IsLast:  true,
	}

	hArrow9 := diagrams.Shape{
		Type: diagrams.RightArrow,
	}

	subNode9 := diagrams.Shape{
		Type:    diagrams.HRectangle,
		Content: "SubRight2",

		IsLast: true,
	}

	subNode10 := diagrams.Shape{
		Type:    diagrams.HRectangle,
		Content: "SubDownnn",

		IsLast: true,
	}

	hArrow10 := diagrams.Shape{
		Type: diagrams.DownArrow,
	}

	hArrow11 := diagrams.Shape{
		Type: diagrams.DownArrow,
	}

	subNode11 := diagrams.Shape{
		Type:    diagrams.HRectangle,
		Content: "SubDownnn2",

		IsLast: true,
	}

	// Make connections
	// if head of arrow points to left side of box
	// then the left of box should have a reference to head of the arrow
	// like this
	/*
		diagrams.AddToRight(&hArrow8, &subNode8) // Adds subNode8 to right of hArrow8
	*/

	// Check diagram.go for more details on how Shape works

	diagrams.AddToBottom(&subNode3, &hArrow4)
	diagrams.AddToLeft(&subNode4, &hArrow5)
	diagrams.AddToRight(&subNode5, &hArrow5)
	diagrams.AddToRight(&hArrow6, &subNode5)
	diagrams.AddToRight(&subNode6, &hArrow6)
	diagrams.AddToLeft(&hArrow7, &subNode7)
	diagrams.AddToRight(&hArrow7, &subNode6)
	diagrams.AddToTop(&subNode5, &hArrow5)
	diagrams.AddToRight(&subNode5, &hArrow8)
	diagrams.AddToRight(&hArrow8, &subNode8)
	diagrams.AddToRight(&subNode8, &hArrow9)
	diagrams.AddToRight(&hArrow9, &subNode9)
	diagrams.AddToBottom(&subNode8, &hArrow10)
	diagrams.AddToBottom(&hArrow10, &subNode10)
	diagrams.AddToBottom(&subNode5, &hArrow11)
	diagrams.AddToBottom(&hArrow11, &subNode11)

	// Add shapes to the Diagram
	// The order in which the shapes are addes is importent (since findNode function isn't being used)
	// Since hArrow4 comes after subNode3 it is added after subNode3 and so on

	diagram.AddShapes(subNode3)
	diagram.AddShapes(hArrow4)
	diagram.AddShapes(subNode4)
	diagram.AddShapes(hArrow5)
	diagram.AddShapes(subNode5)
	diagram.AddShapes(hArrow11)
	diagram.AddShapes(subNode11)
	diagram.AddShapes(hArrow6)
	diagram.AddShapes(subNode6)
	diagram.AddShapes(hArrow7)
	diagram.AddShapes(subNode7)
	diagram.AddShapes(hArrow8)
	diagram.AddShapes(subNode8)
	diagram.AddShapes(hArrow10)
	diagram.AddShapes(subNode10)
	diagram.AddShapes(hArrow9)
	diagram.AddShapes(subNode9)

	s := diagrams.NewStore()

	// Renders Shape on canvas using loop
	for _, shape := range diagram.S {
		diagrams.RenderD(&shape, canvas, s)
	}

	// Display canvas in terminal
	canvas.Render()
}

func example2() {
	canvas := draw.NewCanvas(30, 50)

	diagram := diagrams.New()

	A := diagrams.Shape{
		Content: " A ",
		Type:    diagrams.Rectangle,
	}

	Ar1 := diagrams.Shape{
		Content: "",
		Type:    diagrams.DownArrow,
	}

	B := diagrams.Shape{
		Content:    " B ",
		Type:       diagrams.Rectangle,
		IsJunction: true,
		IsLast:     true,
	}

	Ar2 := diagrams.Shape{
		Content: "",
		Type:    diagrams.LeftArrow,
	}

	C := diagrams.Shape{
		Content: " C ",
		Type:    diagrams.HRectangle,
		IsLast:  true,
	}

	Ar3 := diagrams.Shape{
		Content: "",
		Type:    diagrams.RightArrow,
	}

	D := diagrams.Shape{
		Content: " D ",
		Type:    diagrams.HRectangle,
		IsLast:  true,
	}

	diagrams.AddToBottom(&A, &Ar1)

	diagrams.AddToBottom(&Ar1, &B)

	diagrams.AddToLeft(&B, &Ar2)

	diagrams.AddToRight(&B, &Ar3)

	diagrams.AddToRight(&C, &Ar2)

	diagrams.AddToRight(&Ar3, &D)

	diagram.AddShapes(A)
	diagram.AddShapes(Ar1)
	diagram.AddShapes(B)
	diagram.AddShapes(Ar2)
	diagram.AddShapes(C)
	diagram.AddShapes(Ar3)
	diagram.AddShapes(D)

	s := diagrams.NewStore()

	for _, shape := range diagram.S {
		diagrams.RenderD(&shape, canvas, s)
	}
	canvas.Render()
}
