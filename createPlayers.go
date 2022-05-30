package main

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var createPlayersQuery = `
MERGE (player:Person:Player { firstName: $firstName, lastName: $lastName })
SET player += { jersey: $jersey, position: $position, nbaDebutYear: $nbaDebutYear }
RETURN player
`

func createPlayersTXWork(fb []FetchedPlayer) neo4j.TransactionWork {
	return func(transaction neo4j.Transaction) (interface{}, error) {
		for _, player := range fb {
			params := map[string]interface{}{
				"firstName":    player.FirstName,
				"lastName":     player.LastName,
				"jersey":       player.Jersey,
				"nbaDebutYear": player.NBADebutYear,
				"position":     player.Position,
			}
			result, err := transaction.Run(createPlayersQuery, params)

			if err != nil {
				transaction.Close()
				return nil, result.Err()
			}
			if result.Next() {
				fmt.Println(result.Record().Values...)
			}
		}
		return true, nil
	}
}

func createPlayers(fp []FetchedPlayer) error {
	driver, err := createDriver()
	if err != nil {
		return err
	}
	defer closeDriver(driver)
	session := createSession(driver)
	defer session.Close()
	created, err := session.WriteTransaction(createPlayersTXWork(fp))
	if err != nil {
		return err
	}
	fmt.Println("created:", created)
	return nil
}
