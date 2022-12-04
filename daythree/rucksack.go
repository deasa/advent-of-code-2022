package daythree

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const itemTypeOptions = "abcdefghijklmnopqrstuvwxyz"
var itemTypeOptionsUpper = strings.ToUpper(itemTypeOptions)

type itemType string
type rucksack struct {
	firstCompartment, secondCompartment string
}

func (r *rucksack) findSharedItem() itemType {
	for _, v := range r.firstCompartment {
		if strings.Contains(r.secondCompartment, string(v)) {
			return itemType(v)
		}
	}
	panic("couldn't find shared item")
}

func (it itemType) getPriority() int {
	i := strings.Index(itemTypeOptions, string(it))
	add26 := false
	if i == -1 {
		add26 = true
		i = strings.Index(itemTypeOptionsUpper, string(it))
		if i == -1 {
			panic("invalid priority string")
		}
	}

	if add26 {
		return i + 1 + 26
	}
	return i + 1
}

func Play() {
	rucksacks, err := readInputFromFile()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(calculateSumOfSharedItems(rucksacks))
	// for _, v := range rucksacks {
	// 	fmt.Printf("Shared item: %s\nPriority: %d\n", v.findSharedItem(), v.findSharedItem().getPriority())
	// }

	// fmt.Printf("%#v", rucksacks)
}

func calculateSumOfSharedItems(rucksacks []rucksack) int {
	sum := 0
	for _, v := range rucksacks {
		sum += v.findSharedItem().getPriority()
	}
	return sum
}

func readInputFromFile() ([]rucksack, error) {
	b, err := os.ReadFile("daythree/input.txt")
	if err != nil {
		return nil, err
	}
	s := strings.Split(string(b), "\n")

	var rL []rucksack
	for _, v := range s {
		length := len(v)
		if length % 2 != 0 {
			panic("odd number of items can't be evenly split")
		}
		rL = append(rL, rucksack{
			firstCompartment: v[:length/2],
			secondCompartment: v[length/2:],
		})
	}
	return rL, nil
}