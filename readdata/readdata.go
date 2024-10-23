package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./data/sample_small.csv")

	if err != nil {
		log.Fatal("Read error", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Read error!")
	}

	for _, eachrecord := range records {
		fmt.Println(eachrecord)
	}

}
