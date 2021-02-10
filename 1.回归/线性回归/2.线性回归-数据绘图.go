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
	for _,colName := range ADDF.Names(){
		plotvals := make(plotter.Values,ADDF.Nrow())
		for i,floatval := range ADDF.Col(colName).Float(){
			plotvals[i]=floatval
		}

		p,err := plot.New()
		if err != nil{
			fmt.Println(err)
		}
		p.Title.Text = fmt.Sprintf("图的展示字段%s",colName)

		h,err :=plotter.NewHist(plotvals,16)
		if err != nil{
			fmt.Println(err)
		}

		h.Normalize(1)
		p.Add(h)
		if err := p.Save(4*vg.Inch, 4*vg.Inch,colName+"_hisg.png");err!=nil{
			fmt.Println(err)
		}

	}





}
