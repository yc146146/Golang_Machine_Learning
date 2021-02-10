package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/ensemble"
	"github.com/sjwhitworth/golearn/evaluation"
	"math"
	"math/rand"
	"time"
)

func main() {

	path := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\5.监督学习\\iris.csv"
	irisData,err := base.ParseCSVToInstances(path, true)

	if err !=nil{
		fmt.Println("x", err)
	}

	rand.Seed(time.Now().UnixNano())

	tree := ensemble.NewRandomForest(10,2)
	//tree := ensemble.NewMultiLinearSVC()

	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData,tree,5)
	if err !=nil{
		fmt.Println("x", err)
	}

	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)

	stdev := math.Sqrt(variance)
	fmt.Println(mean, stdev)


}
