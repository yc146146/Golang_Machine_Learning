package main

import (
	"encoding/csv"
	"fmt"
	"github.com/mash/gokmeans"
	"io"
	"os"
	"strconv"
)

func main() {
	path := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\4.回归\\线性回归\\Advertising.csv"

	ADfile,err := os.Open(path)
	if err != nil{
		fmt.Println(err)
	}

	defer ADfile.Close()

	r := csv.NewReader(ADfile)
	r.FieldsPerRecord = 3

	//节点
	var data []gokmeans.Node

	for  {
		record, err := r.Read()
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println(err)
		}

		if record[0]=="Driver_ID"{
			continue
		}

		var point []float64
		for i:=1;i<3;i++{
			val, err := strconv.ParseFloat(record[i],64)
			if err != nil{
				fmt.Println(err)
			}
			point = append(point, val)

			data = append(data,gokmeans.Node{point[0], point[1]})
		}
		success,centroids := gokmeans.Train(data, 2,50)
		if !success{
			fmt.Println("失败")
		}

		//输出核心
		for _,centroid := range centroids{
			fmt.Println(centroid)

			for _, observation := range data {
				index := gokmeans.Nearest(observation, centroids)
				fmt.Println(observation, "belongs in cluster", index+1, ".")
			}
		}

	}



}
