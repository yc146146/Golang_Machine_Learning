package main

import (
	"encoding/csv"
	"fmt"
	"github.com/berkmancenter/ridge"
	"github.com/gonum/matrix/mat64"
	"os"
	"strconv"
)

func main() {

	path := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\4.回归\\线性回归\\training.csv"

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord=4
	traindata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	//特征数据
	featureData := make([]float64,4*len(traindata))
	yData := make([]float64, len(traindata))

	var featureindex int
	var yindex int

	for idx, record := range traindata{
		if idx == 0{
			continue
		}

		for i,val := range record{
			valParsed,err := strconv.ParseFloat(val, 64)
			if err != nil {
				fmt.Println(err)
			}

			if i<3 {
				if i ==0 {
					featureData[featureindex]=1
					featureindex++
				}
				featureData[featureindex]=valParsed
				featureindex++
			}
			if i == 3{
				yData[yindex]=valParsed
				yindex++
			}
		}
	}
	features := mat64.NewDense(len(traindata), 4, featureData)
	y := mat64.NewVector(len(traindata), yData)

	//fmt.Println(features)
	//fmt.Println(y)

	r := ridge.New(features, y, 1.0)
	r.Regress()

	c1 := r.Coefficients.At(0,0)
	c2 := r.Coefficients.At(1,0)
	c3 := r.Coefficients.At(2,0)
	c4 := r.Coefficients.At(3,0)

	fmt.Printf("y=%f+%fTV+%fRadio+%fNewspaper\n",c1,c2,c3,c4)


}
