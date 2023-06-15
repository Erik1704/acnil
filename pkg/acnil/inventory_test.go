package acnil_test

import (
	"testing"
	"time"

	"github.com/Erik1704/acnil/pkg/acnil"
)

func TestNewGame(t *testing.T) {
	record := []string{
		"1936",                           //     id
		"1936: la Guerra Civil Española", //  juego
		"Centro",                         //   local
		"",                               //              prestado a
		"",                               //       comentarios
		"21/03/2020",                     //          fecha prestamo
		"20/01/1900",                     //            fecha devolucion
		"50,00 €", "Brigada Wargames", "318323", "6,4", "1,8", "14", "2", "6", "240", "2019", "Extensive use of text - massive conversion needed to be playable", "", "", ""}
	game, err := acnil.NewGame(record)
	if err != nil {
		t.Logf("ha petao hermano, %s", err)
		t.FailNow()
	}
	if game.Id != "1936" {
		t.Logf(",expected %s,actual %s", "1936", game.Id)
		t.FailNow()
	}
	expected := time.Date(2020, 3, 21, 0, 0, 0, 0, time.UTC)
	if game.TakeDate != expected {
		t.Logf(",expected %s,actual %s", expected, game.TakeDate)
		t.FailNow()
	}
}
