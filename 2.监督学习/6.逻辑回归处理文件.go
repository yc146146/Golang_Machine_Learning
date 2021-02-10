package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\5.监督学习"

	file,err := os.Open(path+"\\loan_data.csv")


	if err !=nil{
		fmt.Println( err)
	}

	defer file.Close()

	reader := csv.NewReader(file)


	reader.FieldsPerRecord = 2

	//读取全部数据
	rawCSVData, err := reader.ReadAll()

	if err !=nil{
		fmt.Println( err)
	}

	savefile, err := os.Create(path + "\\load_clean_data.csv")

	if err !=nil{
		fmt.Println( err)
	}

	defer savefile.Close()

	w := csv.NewWriter(savefile)
	for idx, record := range rawCSVData{
		if idx == 0{
			if err := w.Write([]string{"FICO_score","class "}); err != nil{
				fmt.Println(err)
			}
			continue
		}
		outRecord := make([]string, 2)
		score, err := strconv.ParseFloat(strings.Split(record[0], "-")[0], 64)
		if err !=nil{
			fmt.Println( err)
		}

		//计算比例
		outRecord[0] = strconv.FormatFloat((score-640.0)/(830.0-640.0),'f',4,64)

		rate,err := strconv.ParseFloat(strings.TrimSuffix(record[1], "%"), 64)
		if err !=nil{
			fmt.Println( err)
		}

		if rate <= 12.0{
			outRecord[1]="1.0"

			if err := w.Write(outRecord); err != nil{
				fmt.Println(err)
			}

			continue
		}

		outRecord[1]="0.0"


	}
	w.Flush()

}
