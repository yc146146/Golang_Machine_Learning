package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func predict(score float64)float64{
	p:=1/(1+math.Exp(-13.417345*score+4.956825))
	if p>=0.5{
		return 1.0
	}
	return 0.0
}

func main() {
	path := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\5.监督学习"

	file, err := os.Open(path + "\\ltraining.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	var observed[]float64
	var predicted[]float64

	line := 1
	for {
		record,err := reader.Read()
		if err == io.EOF{
			break
		}
		if line == 1{
			line++
			continue
		}
		obs,err := strconv.ParseFloat(record[1],64)
		if err != nil {
			fmt.Println(err)
		}
		score, err := strconv.ParseFloat(record[0],64)

		if err != nil {
			fmt.Println(err)
			continue
		}

		//预测数据
		predictedVal := predict(score)
		//真实数据
		observed=append(observed,obs)
		//预测
		predicted=append(predicted, predictedVal)
		line++

	}

	var truePosNeg int
	for idx, oVal := range observed{
		if oVal == predicted[idx]{
			truePosNeg++
		}
	}

	fmt.Println("正确率",float64(truePosNeg)/float64(len(observed)))

}