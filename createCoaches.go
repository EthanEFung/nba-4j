package main

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var createCoachesQuery = `
MERGE (coach:Person:Coach { firstName: $firstName, lastName: $lastName })
SET coach.isAssistant = $isAssistant
RETURN coach
`

func createCoachesTXWork(fc []FetchedCoach) neo4j.TransactionWork {
	return func(transaction neo4j.Transaction) (interface{}, error) {
		for _, coach := range fc {
			params := map[string]interface{}{
				"firstName":   coach.FirstName,
				"lastName":    coach.LastName,
				"isAssistant": coach.IsAssistant,
			}
			result, err := transaction.Run(createCoachesQuery, params)

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

func createCoaches(fc []FetchedCoach) error {
	driver, err := createDriver()
	if err != nil {
		return err
	}
	defer closeDriver(driver)
	session := createSession(driver)
	defer session.Close()
	created, err := session.WriteTransaction(createCoachesTXWork(fc))
	if err != nil {
		return err
	}
	fmt.Println("created:", created)
	return nil
}
