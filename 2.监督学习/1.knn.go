package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	knn2 "github.com/sjwhitworth/golearn/knn"
	"math"
)

func main() {
	path := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\5.监督学习\\iris.csv"

	irisData,err := base.ParseCSVToInstances(path, true)

	//knn := knn2.NewKnnClassifier("euclidean","linear",2)
	knn := knn2.NewKnnClassifier("euclidean","kdtree",2)

	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData,knn,5)

	if err !=nil{
		fmt.Println("x", err)
	}

	fmt.Println("-----------------------")
	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	fmt.Println(mean)
	fmt.Println(stdev)



}
