package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/trees"
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
	//tree := trees.NewID3DecisionTree(0.7)
	tree := trees.NewID3DecisionTree(0.8)
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData,tree,5)
	if err !=nil{
		fmt.Println("x", err)
	}

	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)

	stdev := math.Sqrt(variance)
	fmt.Println(mean, stdev)


}
