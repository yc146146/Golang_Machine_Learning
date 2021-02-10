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
	yvals := ADDF.Col("sales").Float()

	for _,colName := range ADDF.Names(){
		//处理x,y之间的关系
		plotvals := make(plotter.XYs,ADDF.Nrow())
		for i,floatval := range ADDF.Col(colName).Float(){
			plotvals[i].X=floatval
			plotvals[i].Y=yvals[i]
		}

		p,err := plot.New()
		if err != nil{
			fmt.Println(err)
		}

		p.X.Label.Text=colName
		p.Y.Label.Text="y"
		p.Add(plotter.NewGrid())

		s,err := plotter.NewScatter(plotvals)

		if err != nil{
			fmt.Println(err)
		}


		s.GlyphStyle.Radius=vg.Points(3)

		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch,colName+"_scatter.png");err!=nil{
			fmt.Println(err)
		}

	}





}
