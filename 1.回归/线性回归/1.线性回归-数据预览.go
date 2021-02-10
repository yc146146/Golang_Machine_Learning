package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"os"
)

func main() {
	path := "G:\\Go_workspace\\src\\yinchen.com\\12.机器学习\\4.回归\\线性回归\\Advertising.csv"

	ADfile,err := os.Open(path)
	if err != nil{
		fmt.Println(err)
	}

	defer ADfile.Close()
	ADDF := dataframe.ReadCSV(ADfile)
	fmt.Println(ADDF.Describe())
	fmt.Println(ADDF)


}
