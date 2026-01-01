package main

import (
	"fmt"
	"log"

	"timelog/internal/utils"
)

func main() {
	storeFile, err := utils.GetStoreFile()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(storeFile)
}
