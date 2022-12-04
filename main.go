package main

import (
	"fmt"
	"log"

	"github.com/deasa/advent-of-code-2022/dayone"
)


func main() {
	rations, err := dayone.ReadInputFromFile()
	if err != nil {
		log.Fatalln(err)
	}

	size := dayone.CalculateLargestRationPack(rations)
	topThree := dayone.CalculateTopThreeLargestRationPacks(rations)

	fmt.Printf("Largest pack %v\n", size)
	fmt.Println("Largest three packs: ", topThree)
}

