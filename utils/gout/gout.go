package gout

import "github.com/fatih/color"

func GreenSPrintf(format string, a ...interface{}) {
	color.Green(format, a)
}

func RedSPrintf(format string, a ...interface{}) {
	color.Red(format, a)
}

func GreenPrintln(format string) {
	color.Green(format)
}

func RedPrintln(format string) {
	color.Red(format)
}
