package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"os"
)

type Centroid []float64

func main() {
	path := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\6.无监督学习\\iris.csv"
	irisFile, err := os.Open(path)
	if err != nil{
		fmt.Println(err)
	}
	defer irisFile.Close()

	//数据集
	irisDF := dataframe.ReadCSV(irisFile)
	speciesNames := []string{
		"setosa","virginica","versicolor",
	}
	centroid := make(map[string]Centroid)
	for _, species := range speciesNames{
		filter := dataframe.F{
			Colname: "Species",
			Comparator: "==",
			Comparando: species,

		}
		filtered := irisDF.Filter(filter)
		summaryDF := filtered.Describe()
		var c Centroid
		for _, feature := range summaryDF.Names(){
			if feature=="column"||feature=="Species"{
				continue
			}
			c=append(c,summaryDF.Col(feature).Float()[0])
		}

		//记录map
		centroid[species]=c

	}

	for _,species := range speciesNames{
		fmt.Println(species, centroid[species])
	}


}
