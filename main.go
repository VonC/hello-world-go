package main

import "github.com/gizak/termui"

func main() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	p := termui.NewPar(":PRESS q TO QUIT DEMO Hello World")
	p.Height = 3
	p.Width = 50
	p.TextFgColor = termui.ColorWhite
	p.BorderLabel = "Hello-World"
	p.BorderFg = termui.ColorCyan

	termui.Render(p)

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Loop()
}
