package main

import (
	"fmt"

	"github.com/Erik1704/acnil/pkg/acnil"
)

func main() {
	games, err := acnil.LoadGames("../../data/inventario.csv")
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	for i := 0; i < len(games); i++ {
		fmt.Printf("%#v \n", games[i])
	}
	//acnil.Game
	//fmt.Print("hello world")

}
