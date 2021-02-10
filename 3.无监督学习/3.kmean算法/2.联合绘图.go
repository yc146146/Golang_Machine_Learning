package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"os"
)

func main() {
	path := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\4.回归\\线性回归\\Advertising.csv"

	ADfile,err := os.Open(path)
	if err != nil{
		fmt.Println(err)
	}

	defer ADfile.Close()
	ADDF := dataframe.ReadCSV(ADfile)
	yvals := ADDF.Col("Distance_Feature").Float()
	plotvals := make(plotter.XYs,ADDF.Nrow())

	p,err := plot.New()
	if err != nil{
		fmt.Println(err)
	}

	p.X.Label.Text="Distance_Feature"
	p.Y.Label.Text="Distance"
	p.Add(plotter.NewGrid())

	s,err := plotter.NewScatter(plotvals)

	if err != nil{
		fmt.Println(err)
	}


	s.GlyphStyle.Radius=vg.Points(3)

	p.Add(s)

	if err := p.Save(4*vg.Inch, 4*vg.Inch,"last_scatter.png");err!=nil{
		fmt.Println(err)
	}





}
