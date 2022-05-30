package main

type PlayersJSON struct {
	League struct {
		Standard []FetchedPlayer `json:"standard"`
	} `json:"league"`
}

type FetchedPlayer struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Jersey       string `json:"jersey"`
	Position     string `json:"pos"`
	NBADebutYear string `json:"nbaDebutYear"`
	Teams        []struct {
		ID          string `json:"teamId"`
		SeasonStart string `json:"seasonStart"`
		SeasonEnd   string `json:"seasonEnd"`
	} `json:"teams"`
	Draft struct {
		TeamID      string `json:"teamId"`
		SeasonYear  string `json:"seasonYear"`
		RoundNumber string `json:"roundNum"`
		PickNumber  string `json:"pickNum"`
	} `json:"draft"`
}

func fetchPlayers() ([]FetchedPlayer, error) {
	var playersJSON PlayersJSON
	err := fetchAndUnmarshal("http://data.nba.net/prod/v1/2020/players.json", &playersJSON)
	if err != nil {
		return nil, err
	}
	return playersJSON.League.Standard, nil
}
