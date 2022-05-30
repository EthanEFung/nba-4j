package main

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var createTeamsQuery = `
MERGE (team:Team { fullName: $fullName, shortName: $shortName, city: $city, tricode: $tricode })
`

func createTeamsTXWork(ft []FetchedTeam) neo4j.TransactionWork {
	return func(tx neo4j.Transaction) (interface{}, error) {
		for _, team := range ft {
			params := map[string]interface{}{
				"fullName":  team.FullName,
				"shortName": team.TeamShortName,
				"city":      team.City,
				"tricode":   team.Tricode,
			}
			result, err := tx.Run(createTeamsQuery, params)
			if err != nil {
				return nil, result.Err()
			}
			if result.Next() {
				fmt.Println(result.Record().Values...)
			}
		}
		return true, nil
	}
}

func createTeams(ft []FetchedTeam) error {
	driver, err := createDriver()
	if err != nil {
		return err
	}
	defer closeDriver(driver)
	session := createSession(driver)
	defer session.Close()
	created, err := session.WriteTransaction(createTeamsTXWork(ft))
	if err != nil {
		return err
	}
	fmt.Println("created:", created)
	return nil
}
