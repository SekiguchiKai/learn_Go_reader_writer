package main

import (
	"os"
	"github.com/SekiguchiKai/learn_Go_reader_writer/tools"
	"fmt"
	"golang.org/x/text/encoding/japanese"
)

func main(){
	file := new(os.File)
	if err := tools.ReadFile("sample", file);  err != nil {
		fmt.Println(err.Error())
	}

	decoder := japanese.ShiftJIS.NewDecoder()
	contents, err := tools.ReadCSV(file, decoder)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(contents)
}
