package main

import (
	"errors"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var relatePlaysForQuery = `
MATCH (p:Player { firstName: $firstName, lastName: $lastName })
MATCH (t:Team { city: $city, fullName: $fullName })
MERGE (p)-[pf:PLAYS_FOR { seasonStart: $seasonStart }]->(t)
SET pf.seasonEnd = $seasonEnd
RETURN p, t
`

func relatePlayesForTXWork(tMap map[string]FetchedTeam, fp []FetchedPlayer) neo4j.TransactionWork {
	return func(tx neo4j.Transaction) (interface{}, error) {
		for _, player := range fp {
			for _, association := range player.Teams {
				team, teamExists := tMap[association.ID]
				if teamExists == false {
					return nil, errors.New("Your logic is off")
				}

				params := map[string]interface{}{
					"firstName":   player.FirstName,
					"lastName":    player.LastName,
					"city":        team.City,
					"fullName":    team.FullName,
					"seasonStart": association.SeasonStart,
					"seasonEnd":   association.SeasonEnd,
				}
				res, err := tx.Run(relatePlaysForQuery, params)
				if err != nil {
					return nil, err
				}
				if res.Next() {
					records := res.Record().Values
					fmt.Println(records[0], records[1])
				}
			}
		}

		return true, nil
	}
}
func relatePlaysFor(tMap map[string]FetchedTeam, fp []FetchedPlayer) error {
	driver, err := createDriver()
	if err != nil {
		return err
	}
	defer closeDriver(driver)
	session := createSession(driver)
	defer session.Close()
	created, err := session.WriteTransaction(relatePlayesForTXWork(tMap, fp))
	if err != nil {
		return err
	}
	fmt.Println("created:", created)
	return nil
}
