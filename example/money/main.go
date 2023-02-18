package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nomad-software/mongo"
)

func main() {

	m, err := mongo.MoneyGBP(1055)

	if err != nil {
		log.Fatal("Error occured creating money")
	}

	fmt.Printf("Money: %s\n", m)

	shares := m.Split(3)
	fmt.Printf("Shares: %s\n", shares)

	shares = m.Allocate(1, 2, 3)
	fmt.Printf("Allocations: %s\n", shares)

	json, _ := json.Marshal(m)
	fmt.Println(string(json))
}
