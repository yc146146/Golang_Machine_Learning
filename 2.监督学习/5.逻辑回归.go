package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"math"
)

func logsitc(x float64)float64{
	return 1/(1+math.Exp(-x))
}

func main1() {
	fmt.Println(logsitc(1.0))
}

func main() {


	p,err := plot.New()
	if err != nil{
		fmt.Println(err)
	}
	p.Title.Text = "逻辑回归图"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "f(x)"

	myp := plotter.NewFunction(func(f float64) float64 {
		return logsitc(f)
	})

	myp.Color=color.RGBA{B:255,A:255}
	p.Add(myp)
	p.X.Min=-10
	p.X.Max=10
	p.Y.Min=-0.1
	p.Y.Max=1.1

	if err := p.Save(4*vg.Inch, 4*vg.Inch,"logic.png");err!=nil{
		fmt.Println(err)
	}


}
