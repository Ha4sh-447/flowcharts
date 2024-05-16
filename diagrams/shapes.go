package diagrams

import (
	"fmt"
)

const H_s_rectangle_len = len("╭――――――――――――――――――╮")
const H_a_len = len("--%s-->")
const s_rectangle_len = len("┌──────────────┐")

func Shape_rectangle(content string) {
	fmt.Println("┌──────────────┐")
	fmt.Printf("│ %-12s │\n", content)
	fmt.Println("└──────────────┘")
}

func Horizontal_shape_rectangle(content string) {
	fmt.Printf(
		`╭――――――――――――――――――╮
│%-12s      │
╰――――――――――――――――――╯`, content)
}

func Horizontal_shape_rectangle2(content string) {
	fmt.Println("┌──────────────┐")
	fmt.Printf("│ %-12s │\n", content)
	fmt.Println("└──────────────┘")
}

func Shape_daimond(content string) {
	fmt.Println("    /\\")
	fmt.Println("   /  \\")
	fmt.Println("  /    \\")
	fmt.Printf(" /  %-4s  \\\n", content)
	fmt.Println("  \\    /")
	fmt.Println("   \\  /")
	fmt.Println("    \\/")

}

func VerticalArrow(content string) {
	fmt.Printf("     │\n")
	fmt.Printf("     %s\n", content)
	fmt.Printf("     │\n")
	fmt.Printf("     ↓\n")
}

func HorizontalArrow(content string) {
	// fmt.Printf("\n")
	fmt.Printf(`
--%s-->`, content)
	// fmt.Printf("\n")
}

func HorizontalLine(content string) {
	fmt.Printf("\n")
	fmt.Printf("-----%s-----", content)
	fmt.Printf("\n")
}

func VerticalLine(content string) {
	fmt.Printf("     │\n")
	fmt.Printf("     │\n")
	fmt.Printf("     %s\n", content)
	fmt.Printf("     │\n")
	fmt.Printf("     │\n")
}
