package main

import (
	"encoding/csv"
	"fmt"
	"github.com/sajari/regression"
	"math"
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
	r.SetVar(1, "radio")
	r.SetVar(2, "newspaper")
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

		Rval, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Println(err)
		}
		Nval, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			fmt.Println(err)
		}

		r.Train(regression.DataPoint(yval,[]float64{TVval,Rval,Nval}))

	}

	r.Run()
	fmt.Println(r.Formula)


	patht := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\4.回归\\线性回归\\test.csv"

	ftest, err := os.Open(patht)
	if err != nil {
		fmt.Println(err)
	}
	defer ftest.Close()

	reader = csv.NewReader(ftest)
	reader.FieldsPerRecord=4
	testdata, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	//误差
	var mAE float64
	for i,record := range testdata{
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

		Rval, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Println(err)
		}

		Nval, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			fmt.Println(err)
		}

		//预测的
		ypredicted, err := r.Predict([]float64{TVval,Rval,Nval})

		mAE += math.Abs(yval-ypredicted)/float64(len(testdata))


	}

	fmt.Println(mAE)

}
