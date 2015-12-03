package main

import "github.com/gizak/termui"

func main() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	p := termui.NewPar(":PRESS q or Esc TO QUIT DEMO Hello World")
	p.Height = 3
	p.Width = 50
	p.TextFgColor = termui.ColorWhite
	p.BorderLabel = "Hello-World"
	p.BorderFg = termui.ColorCyan

	termui.Render(p)

	termui.Handle("/sys", func(e termui.Event) {
		k, ok := e.Data.(termui.EvtKbd)
		if ok && (k.KeyStr == "q" || k.KeyStr == "<escape>") {
			termui.StopLoop()
		}
	})
	termui.Loop()
}
