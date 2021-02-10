package main

import (
	"bufio"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"os"
)

func main() {

	path := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\4.回归\\线性回归\\load_clean_data.csv"

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	diabetesDF := dataframe.ReadCSV(file)
	trainNum := (4*diabetesDF.Nrow())/5
	testNum := (diabetesDF.Nrow())/5
	if trainNum+testNum < diabetesDF.Nrow(){
		trainNum++
	}
	trainIdx:=make([]int, trainNum)
	testIdx:=make([]int, testNum)

	for i:=0; i<trainNum;i++{
		trainIdx[i]=i
	}

	for i:=0; i<testNum;i++{
		testIdx[i]=trainNum+i
	}

	trainDF := diabetesDF.Subset(trainIdx)
	testDF := diabetesDF.Subset(testIdx)


	setMap := map[int]dataframe.DataFrame{
		0:trainDF,
		1:testDF,
	}

	pathdir := "./"
	for idx, setName := range[]string{pathdir+"ltraining.csv",pathdir+"ltest.csv"}{
		filex, err := os.Create(setName)
		if err != nil {
			fmt.Println(err)
		}

		w:=bufio.NewWriter(filex)
		if err := setMap[idx].WriteCSV(w);err !=nil{
			fmt.Println(err)
		}

	}


}
