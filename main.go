package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"timelog/internal/store"
)

func main() {
	store, err := store.ReadFile()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", *store)

	if err := json.NewEncoder(os.Stdout).Encode(store); err != nil {
		log.Fatal(err)
	}
}
