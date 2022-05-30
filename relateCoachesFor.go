package main

import (
	"errors"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var relateCoachesForQuery = `
MATCH (c:Coach { firstName: $firstName, lastName: $lastName })
MATCH (t:Team { city: $city, fullName: $fullName })
MERGE (c)-[cf:COACHES_FOR]->(t)
RETURN c, cf, t
`

func relateCoachesForTXWork(tMap map[string]FetchedTeam, fc []FetchedCoach) neo4j.TransactionWork {
	return func(tx neo4j.Transaction) (interface{}, error) {
		for _, coach := range fc {
			team, teamExists := tMap[coach.TeamId]
			if teamExists == false {
				return nil, errors.New("Your logic is off")
			}

			params := map[string]interface{}{
				"firstName": coach.FirstName,
				"lastName":  coach.LastName,
				"city":      team.City,
				"fullName":  team.FullName,
			}
			res, err := tx.Run(relateCoachesForQuery, params)
			if err != nil {
				return nil, err
			}
			if res.Next() {
				fmt.Println(res.Record().Values...)
			}
		}

		return true, nil
	}
}
func relateCoachesFor(tMap map[string]FetchedTeam, fc []FetchedCoach) error {
	driver, err := createDriver()
	if err != nil {
		return err
	}
	defer closeDriver(driver)
	session := createSession(driver)
	defer session.Close()
	created, err := session.WriteTransaction(relateCoachesForTXWork(tMap, fc))
	if err != nil {
		return err
	}
	fmt.Println("created:", created)
	return nil
}
