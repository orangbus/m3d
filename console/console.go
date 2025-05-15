package console

import (
	"fmt"
	"github.com/mgutz/ansi"
)

func Success(msg string) {
	color := ansi.Color(msg, "green+h")
	fmt.Println(color)
}

func Error(msg string) {
	color := ansi.Color(msg, "red+h")
	fmt.Println(color)
}

func Info(msg string) {
	color := ansi.Color(msg, "info+h")
	fmt.Println(color)
}
