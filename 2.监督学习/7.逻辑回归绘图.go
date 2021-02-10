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
	path := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\5.监督学习"

	file,err := os.Open(path+"\\load_clean_data.csv")


	if err !=nil{
		fmt.Println( err)
	}

	defer file.Close()

	LoadDF := dataframe.ReadCSV(file)
	fmt.Println(LoadDF.Describe())

	for _,colName := range LoadDF.Names(){
		//处理x,y之间的关系
		plotvals := make(plotter.Values,LoadDF.Nrow())
		for i,floatval := range LoadDF.Col(colName).Float(){
			plotvals[i]=floatval

		}

		p,err := plot.New()
		if err != nil{
			fmt.Println(err)
		}

		p.X.Label.Text=colName
		p.Y.Label.Text="y"
		p.Title.Text=fmt.Sprintf("histogram of %s", colName)

		h,err := plotter.NewHist(plotvals, 16)
		if err != nil{
			fmt.Println(err)
		}
		h.Normalize(1)

		p.Add(h)
		if err := p.Save(4*vg.Inch, 4*vg.Inch,colName+"_scatter.png");err!=nil{
			fmt.Println(err)
		}

	}

}
