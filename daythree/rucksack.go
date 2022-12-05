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
type rucksackGroup [3]rucksack

func Play() {
	rucksacks, err := readInputFromFile()
	if err != nil {
		log.Fatalln(err)
	}

	rucksackGroups := splitIntoRucksackGroups(rucksacks)

	fmt.Println(calculateSumOfSharedItems(rucksacks))
	fmt.Println(calculateSumOfBadgeItemTypes(rucksackGroups))
	// for _, v := range rucksacks {
	// 	fmt.Printf("Shared item: %s\nPriority: %d\n", v.findSharedItem(), v.findSharedItem().getPriority())
	// }

	// fmt.Printf("%#v", rucksacks)
}

func splitIntoRucksackGroups(rucksacks []rucksack) []rucksackGroup {
	rG := []rucksackGroup{}
	
	r := rucksackGroup{}
	for i, v := range rucksacks {
		r[i%3] = v
		if (i+1) % 3 == 0 {
			rG = append(rG, r)
			r = rucksackGroup{}
		}
	}
	return rG
}

func (rg rucksackGroup) findCommonAcrossRucksacks() itemType {
	sacks := []string{
		rg[0].firstCompartment+rg[0].secondCompartment,
		rg[1].firstCompartment+rg[1].secondCompartment,
		rg[2].firstCompartment+rg[2].secondCompartment,
	}
	for _, v := range sacks[0] {
		if strings.Contains(sacks[1], string(v)) {
			if strings.Contains(sacks[2], string(v)) {
				return itemType(v)
			}
		}
	}
	panic("couldn't find common item type across the three rucksacks")
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

func calculateSumOfBadgeItemTypes(rucksackGroups []rucksackGroup) int {
	sum := 0
	for _, v := range rucksackGroups {
		sum += v.findCommonAcrossRucksacks().getPriority()
	}
	return sum
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