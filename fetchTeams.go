package main

type TeamsJSON struct {
	League struct {
		Standard []FetchedTeam `json:"standard"`
	} `json:"league"`
}

type FetchedTeam struct {
	FullName       string `json:"fullName"`
	City           string `json:"city"`
	TeamShortName  string `json:"teamShortName"`
	IsNBAFranchise bool   `json:"isNBAFranchise"`
	ConferenceName string `json:"confName"`
	Tricode        string `json:"tricode"`
	DivisionName   string `json:"divName"`
	IsAllStar      bool   `json:"isAllStar"`
	Nickname       string `json:"nickname"`
	URLName        string `json:"urlName"`
	TeamID         string `json:"teamId"`
}

func fetchTeams() ([]FetchedTeam, error) {
	var teams TeamsJSON
	err := fetchAndUnmarshal("http://data.nba.net/prod/v2/2021/teams.json", &teams)
	if err != nil {
		return nil, err
	}
	return teams.League.Standard, nil
}
