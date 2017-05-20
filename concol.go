package main

import (
	"fmt"
	"io"
	"os"

	"github.com/shiena/ansicolor"
)

const rune = "  •••  "

func main() {
	w := ansicolor.NewAnsiColorWriter(os.Stdout)

	fmt.Println()
	fmt.Fprintf(w, "                 ")
	for bg := 40; bg <= 47; bg++ {
		fmt.Fprintf(w, "  %dm     ", bg)
	}
	fmt.Fprintln(w)

	for row := 30; row <= 37; row++ {
		control := fmt.Sprintf("%vm", row)
		bold := fmt.Sprintf("1;%v", row)
		controlBold := fmt.Sprintf("%vm", bold)

		printLine(w, row, control, false)
		printLine(w, row, controlBold, true)
	}

	fmt.Println()
}

func printLine(w io.Writer, fg int, control string, bold bool) {
	fmt.Fprintf(w, "%+6s ", control)

	// No BG color
	cmd := fmt.Sprintf("\x1b[%dm %v ", fg, rune)
	if bold {
		cmd = fmt.Sprintf("\x1b[1;%dm %v ", fg, rune)
	}

	fmt.Fprintf(w, cmd)

	for bg := 40; bg <= 47; bg++ {
		cmd := fmt.Sprintf("\x1b[%d;%dm %v ", fg, bg, rune)
		if bold {
			cmd = fmt.Sprintf("\x1b[1;%d;%dm %v ", fg, bg, rune)
		}

		fmt.Fprintf(w, cmd)
		fmt.Fprintf(w, "\x1b[0m ")
	}
	fmt.Fprintln(w)
}
