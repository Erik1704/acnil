package acnil

import (
	"encoding/csv"
	"os"
	"time"
)

type Game struct {
	Id         string
	Name       string
	Locacion   string
	Holder     string
	Comment    string
	TakeDate   time.Time
	ReturnDate time.Time
}

func LoadGames(filePatch string) ([]Game, error) {
	games := []Game{}
	file, err := os.Open(filePatch)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r := csv.NewReader(file)
	_, err = r.Read()
	if err != nil {
		return nil, err
	}
	for {
		record, err := r.Read()
		if err != nil {
			if err.Error() == "EOF" {
				return games, nil
			} else {
				return nil, err
			}

		}

		game, err := NewGame(record)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}

}
func NewGame(record []string) (Game, error) {
	//con una lista de string que nos devuelva un juegardo

	var takeDate time.Time
	var err error

	if record[5] == "" {
		takeDate = time.Time{}
	} else {
		takeDate, err = time.Parse("2/1/2006", record[5])
		if err != nil {
			takeDate = time.Time{}
		}
	}

	game := Game{
		Id:       record[0],
		Name:     record[1],
		Locacion: record[2],
		Holder:   record[3],
		Comment:  record[4],
		TakeDate: takeDate,
	}
	return game, nil
}
