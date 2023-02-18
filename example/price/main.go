package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nomad-software/mongo"
)

func main() {

	m, err := mongo.PriceGBP(1055, 17.5)

	if err != nil {
		log.Fatal("Error occured creating price")
	}

	fmt.Printf("Price: %s\n", m)

	json, _ := json.Marshal(m)
	fmt.Println(string(json))
}
