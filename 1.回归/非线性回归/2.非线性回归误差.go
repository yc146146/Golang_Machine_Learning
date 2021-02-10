package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

func predict(tv, radio, newspaper float64)float64{
	return 3.038296+0.046537*tv+0.177006*radio+0.001088*newspaper
}


func main() {
	path := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\4.回归\\线性回归\\test.csv"

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord=4
	testdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	var mAE float64

	for idx, record := range testdata{
		if idx == 0{
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

		ypredict := predict(TVval,Rval,Nval)
		mAE += math.Abs(yval-ypredict)/float64(len(testdata))


	}

	fmt.Println(mAE)

}
