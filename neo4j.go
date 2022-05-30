package main

import (
	"fmt"
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func createDriver() (neo4j.Driver, error) {
	uri := os.Getenv("NEO4J_URI")
	pw := os.Getenv("NEO4J_PASSWORD")
	username := os.Getenv("NEO4J_USERNAME")

	return neo4j.NewDriver(uri, neo4j.BasicAuth(username, pw, ""))
}

func closeDriver(driver neo4j.Driver) error {
	return driver.Close()
}

func createSession(driver neo4j.Driver) neo4j.Session {
	return driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
}

func greet(driver neo4j.Driver) {
	session := createSession(driver)
	greeting, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
			map[string]interface{}{"message": "hello, neo4j!"})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}
		return nil, result.Err()
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(greeting.(string))
}
