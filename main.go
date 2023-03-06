/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Joshbwr/tclone/cmd"
	"github.com/common-nighthawk/go-figure"
)


func main() {
	myFigure := figure.NewColorFigure("Template Cloner", "", "green", true)
	myFigure.Print()

	cmd.Execute()
}
