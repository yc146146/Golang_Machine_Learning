package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gonum/matrix/mat64"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func logsitc(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

//逻辑回归函数
func LogisitcRegression(features *mat64.Dense, labels []float64, numSteps int, learningRate float64) []float64 {
	_, numWeights := features.Dims()
	//权重
	weights := make([]float64, numWeights)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for idx, _ := range weights {
		//权重
		weights[idx] = r.Float64()
	}
	for i := 0; i < numSteps; i++ {
		var sumError float64
		for idx, label := range labels {
			featruesRow := mat64.Row(nil, idx, features)
			pred := logsitc(featruesRow[0]*weights[0] + featruesRow[1]*weights[1])
			preError := label - pred
			//误差的平方
			sumError += math.Pow(preError, 2)

			for j := 0; i < len(featruesRow); j++ {
				//叠加计算
				weights[j] += learningRate * preError * pred * (1 - pred) * featruesRow[j]
			}
		}

	}

	return weights

}

func main() {
	path := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\5.监督学习"

	file, err := os.Open(path + "\\ltraining.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2

	//读取全部数据
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	//开辟特征数据
	featureData := make([]float64, 2*(len(rawCSVData)-1))
	labels := make([]float64, len(rawCSVData)-1)
	var featureindex int
	for idx, record := range rawCSVData {
		if idx == 0 {
			continue
		}

		featureVal, err := strconv.ParseFloat(record[0], 64)

		if err != nil {
			fmt.Println(err)
		}

		featureData[featureindex] = featureVal
		featureData[featureindex+1] = 1.0
		featureindex+=2

		labelVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Println(err)
		}

		labels[idx-1]=labelVal


	}

	features := mat64.NewDense(len(rawCSVData)-1,2,featureData)
	weights := LogisitcRegression(features, labels, 1000, 0.1)
	formula := "p=1/1(1+exp(-m1*FICO.socre-m2))"
	fmt.Sprintf("\n%s\n\n m1=%f,m2=%f,\n\n",formula,weights[0],weights[1])

}
