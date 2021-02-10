package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/gonum/floats"
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

	//
	cluster := make(map[string]dataframe.DataFrame)


	for _, species := range speciesNames{
		filter := dataframe.F{
			Colname: "Species",
			Comparator: "==",
			Comparando: species,

		}
		filtered := irisDF.Filter(filter)
		//帅选的数据
		cluster[species]=filtered
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

	labels := irisDF.Col("Species").Records()
	floatColumns:=[]string{
		"Sepal.Length","Sepal.Width","Petal.Length","Petal.Width",
	}

	var silhoutette float64
	for idx, label := range labels{
		var a float64
		for i:=0;i<cluster[label].Nrow();i++{
			cur := dfFloatRow(irisDF, floatColumns,idx)
			other := dfFloatRow(cluster[label],floatColumns,i)
			a += floats.Distance(cur,other,2)/float64(cluster[label].Nrow())
		}
		var otherCluster string
		var distanceToCluster float64

		for _,species := range speciesNames{
			if species==label{
				continue
			}
			distanceTothisCluster := floats.Distance(centroid[label],centroid[species],2)
			if distanceTothisCluster==0.0 || distanceTothisCluster < distanceToCluster{
				otherCluster=species
				distanceToCluster=distanceTothisCluster
			}


		}

		var b float64
		for i:=0;i<cluster[otherCluster].Nrow();i++{
			cur :=dfFloatRow(irisDF,floatColumns,idx)
			other := dfFloatRow(cluster[otherCluster],floatColumns,i)
			b+=floats.Distance(cur,other,2)/float64(cluster[otherCluster].Nrow())
		}

		if a > b{
			silhoutette += ((b-a)/a)/float64((len(labels)))
		}
		silhoutette += ((b-a)/a)/float64((len(labels)))
	}
	fmt.Println(silhoutette)




}

//取出行的数据
func dfFloatRow(df dataframe.DataFrame, names []string, idx int)[]float64{
	var row[]float64
	for _,name := range names{
		row=append(row, df.Col(name).Float()[idx])
	}
	return row
}
