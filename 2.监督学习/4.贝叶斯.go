package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/filters"
	"github.com/sjwhitworth/golearn/naive"
)

func convertToBinary(src base.FixedDataGrid)base.FixedDataGrid{
	b := filters.NewBinaryConvertFilter()
	attrs := base.NonClassAttributes(src)
	for _, a := range attrs{
		b.AddAttribute(a)
	}
	b.Train()
	ret := base.NewLazilyFilteredInstances(src, b)
	return ret
}


func main() {
	path := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\5.监督学习\\"

	trainData,err := base.ParseCSVToInstances(path+"training.csv", true)

	if err !=nil{
		fmt.Println("x", err)
	}

	nb:=naive.NewBernoulliNBClassifier()
	//训练
	nb.Fit(convertToBinary(trainData))


	testData,err := base.ParseCSVToInstances(path+"test.csv", true)

	if err !=nil{
		fmt.Println("x", err)
	}

	predicted,_ := nb.Predict(convertToBinary(testData))

	cm, err := evaluation.GetConfusionMatrix(testData,predicted)

	if err !=nil{
		fmt.Println("x", err)
	}
	fmt.Println(evaluation.GetAccuracy(cm))



}
