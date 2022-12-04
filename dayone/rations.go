package dayone

import (
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type RationListStrings []string
type rationSums []int

func (rls *RationListStrings) Sum() int {
	sum := 0
	for _, v := range *rls {
		if i, err := strconv.Atoi(v); err == nil {
			sum += i
		}
	}
	return sum
}

func CalculateLargestRationPack(rations []RationListStrings) int {
	largest := 0
	for _, v := range rations {
		sum := v.Sum()
		largest = int(math.Max(float64(largest), float64(sum)))
	}
	return largest
}

func CalculateTopThreeLargestRationPacks(rations []RationListStrings) []int {
	var allRations []int
	for _, v := range rations {
		allRations = append(allRations, v.Sum())
	}

	sort.Slice(allRations, func(i, j int) bool {
		return allRations[i] < allRations[j]
	})

	return allRations[len(allRations)-3:]
}

func ReadInputFromFile() ([]RationListStrings, error) {
	b, err := os.ReadFile("dayone/input.txt")
	if err != nil {
		return nil, err
	}
	s := strings.Split(string(b), "\n")
	var rL []RationListStrings
	startIndex := 0
	for i, str := range s {
		if str == "" {
			rL = append(rL, s[startIndex:i])
			startIndex = i+1
		}
	}
	return rL, nil
}