package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/Erik1704/acnil/pkg/acnil"
)

func main() {
	games := []acnil.Game{}
	file, err := os.Open("./data/Inventario.csv")
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	_, err = r.Read()
	for {

		record, err := r.Read()
		if err != nil {
			fmt.Print(err)
			break
		}

		game := acnil.Game{
			Id:   record[0],
			Name: record[1],
		}
		games = append(games, game)
	}
	for i := 0; i < len(games); i++ {
		fmt.Printf(" id: %s  nombre: %s \n ", games[i].Id, games[i].Name)
	}

	//acnil.Game
	//fmt.Print("hello world")
}
