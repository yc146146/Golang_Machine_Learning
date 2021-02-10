package main

import (
	"encoding/csv"
	"fmt"
	"github.com/sajari/regression"
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

	var r regression.Regression
	r.SetObserved("sales")
	//x,y
	r.SetVar(0, "TV")
	for i,record := range traindata{
		if i==0{
			continue
		}
		yval, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			fmt.Println(err)
		}

		TVval, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			fmt.Println(err)
		}
		r.Train(regression.DataPoint(yval,[]float64{TVval}))

	}

	r.Run()
	fmt.Println(r.Formula)



}
