package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./data/Inventario.csv")
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err != nil {
			fmt.Print(err)
			break
		}
		fmt.Println("la linea leida es:", record)
	}
	//fmt.Print("la linea leida es: ", record)
	//acnil.Game
	//fmt.Print("hello world")
}
