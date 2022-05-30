package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
}

func main() {
  teams, err := fetchTeams()
	if err != nil {
		log.Fatal(err)
	}
	teamMap := make(map[string]FetchedTeam)
	for _, team := range teams {
		teamMap[team.TeamID] = team
	}
	// err = createTeams(teams)
	// if err != nil {
	//   log.Fatal(err)
	// }
	// fp, err := fetchPlayers()
	// if err != nil {
	//   log.Fatal(err)
	// }
	// err = createPlayers(fp)
	// if err != nil {
	//   log.Fatal(err)
	// }
	
	if err != nil {
		log.Fatal(err)
	}
	// relatePlaysFor(teamMap, fp)
	result, err := fetchCoaches()
	if err != nil {
		log.Fatal(err)
	}
	// err = createCoaches(result)
	// if err != nil {
	//   log.Fatal(err)
	// }
	err = relateCoachesFor(teamMap, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fin")
}
