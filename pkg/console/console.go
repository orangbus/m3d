package console

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mgutz/ansi"
	"log"
	"strings"
)

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
)

func Success(msg string) {
	color := ansi.Color(msg, "green+h")
	log.Println(color)
}

func Error(name string, err error) {
	msg := fmt.Sprintf(" %s:%s \n", red(name), yellow(err.Error()))
	log.Println(msg)
}

func Info(msg string) {
	color := ansi.Color(msg, "info+h")
	fmt.Println(color)
}

func Alert(title string, msg interface{}) {
	result := fmt.Sprintf("%s:%v", green(title), msg)
	log.Println(result)
	log.Println(strings.Repeat("-", 60))
}

func Hr(s ...string) {
	line := "="
	if len(s) > 0 {
		line = s[0]
	}
	log.Println(strings.Repeat(line, 100))
}
