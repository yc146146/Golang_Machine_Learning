package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"os"
)

func Predict(Radio float64)float64{
	return 9.3178 + Radio*0.1945
}

func main() {
	path := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\4.回归\\线性回归\\Advertising.csv"

	ADfile,err := os.Open(path)
	if err != nil{
		fmt.Println(err)
	}

	defer ADfile.Close()

	ADDF := dataframe.ReadCSV(ADfile)
	yVals := ADDF.Col("sales").Float()
	pts := make(plotter.XYs,ADDF.Nrow())
	//预览数据
	ptsPred := make(plotter.XYs,ADDF.Nrow())

	for i, floatval := range ADDF.Col("radio").Float(){
		pts[i].X = floatval
		pts[i].Y = yVals[i]
		ptsPred[i].X = floatval
		ptsPred[i].Y = Predict(floatval)
	}

	p,err := plot.New()
	if err != nil{
		fmt.Println(err)
	}

	p.X.Label.Text="Radio"
	p.Y.Label.Text="Sales"
	p.Add(plotter.NewGrid())

	s,err := plotter.NewScatter(pts)
	if err != nil{
		fmt.Println(err)
	}

	s.GlyphStyle.Radius = vg.Points(3)

	l,err := plotter.NewLine(ptsPred)
	if err != nil{
		fmt.Println(err)
	}

	l.LineStyle.Width=vg.Points(1)
	l.LineStyle.Dashes=[]vg.Length{vg.Points(5),vg.Points(5)}
	p.Add(s,l)
	if err := p.Save(4*vg.Inch, 4*vg.Inch,"last.png");err!=nil{
		fmt.Println(err)
	}


}