package main

import (
	"Data_model/data_struct/Init"
	"fmt"
	"os"
)

func main() {
	err := Init.InitDb()
	if err != nil {
		fmt.Println("db connect  filed!")
		os.Exit(0)
	}
	//Init.InsertData()
	id := 6
	Init.QueryRowDemo(id)
	//Init.Delete(id)
}
