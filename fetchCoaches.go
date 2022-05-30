package main

type CoachesJSON struct {
	League struct {
		Standard []FetchedCoach `json:"standard"`
	} `json:"league"`
}

type FetchedCoach struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	IsAssistant bool   `json:"isAssistant"`
	TeamId      string `json:"teamId"`
}

func fetchCoaches() ([]FetchedCoach, error) {
	var coachesJSON CoachesJSON
	err := fetchAndUnmarshal("http://data.nba.net/prod/v1/2021/coaches.json", &coachesJSON)
	if err != nil {
		return nil, err
	}
	return coachesJSON.League.Standard, nil
}
