package main

import "github.com/gizak/termui"
import "math"
func main() {

	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	sinps := (func() []float64 {
		n := 220
		ps := make([]float64, n)
		for i := range ps {
			ps[i] = 1 + math.Sin(float64(i)/5)
		}
		return ps
	})()

	p := termui.NewPar(":PRESS q or Esc TO QUIT DEMO Hello World")
	p.Height = 3
	p.Width = 50
	p.TextFgColor = termui.ColorWhite
	p.BorderLabel = "Hello-World"
	p.BorderFg = termui.ColorCyan

	lc1 := termui.NewLineChart()
	lc1.BorderLabel = "dot-mode Line Chart"
	lc1.Mode = "dot"
	lc1.Data = sinps
	lc1.Width = 26
	lc1.Height = 12
	lc1.X = 51
	lc1.DotStyle = '+'
	lc1.AxesColor = termui.ColorWhite
	lc1.LineColor = termui.ColorYellow | termui.AttrBold

	g0 := termui.NewGauge()
	g0.Percent = 40
	g0.Width = 50
	g0.Height = 3
	g0.Y = 3
	g0.BorderLabel = "Slim Gauge"
	g0.BarColor = termui.ColorRed
	g0.BorderFg = termui.ColorWhite
	g0.BorderLabelFg = termui.ColorCyan

	termui.Render(p, g0, lc1)

	termui.Handle("/sys", func(e termui.Event) {
		k, ok := e.Data.(termui.EvtKbd)
		if ok && (k.KeyStr == "q" || k.KeyStr == "<escape>") {
			termui.StopLoop()
		}
	})
	termui.Loop()
}
