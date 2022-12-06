package dayfour

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sectionList []int
type elfCleaningAssignmentPairs struct {
	firstElf sectionList
	secondElf sectionList
}

func Play() {
	assignmentPairs := readInputFromFile()

	total := countFullyContainedAssignments(assignmentPairs)
	fmt.Printf("%#v\n", total)
}

func (p *elfCleaningAssignmentPairs) hasFullyContainedAssignment() bool {
	fStart, fEnd := p.firstElf[0], p.firstElf[len(p.firstElf)-1]
	sStart, sEnd := p.secondElf[0], p.secondElf[len(p.secondElf)-1]

	if fStart >= sStart && fEnd <= sEnd {
		return true
	}
	if sStart >= fStart && sEnd <= fEnd {
		return true
	}
	return false
}

func countFullyContainedAssignments(assignmentPairs []elfCleaningAssignmentPairs) int {
	sum := 0

	for _, pair := range assignmentPairs {
		if pair.hasFullyContainedAssignment() {
			sum++
		}
	}
	return sum
}

func readInputFromFile() []elfCleaningAssignmentPairs {
	b, err := os.ReadFile("dayfour/input.txt")
	if err != nil {
		panic("couldn't read file")
	}

	s := strings.Split(string(b), "\n")
	var pairs []elfCleaningAssignmentPairs
	for _, v := range s {
		pairSplit := strings.Split(v, ",") 
		pairs = append(pairs, elfCleaningAssignmentPairs{
			firstElf: sequentialInt(getStartAndFinish(pairSplit[0])),
			secondElf: sequentialInt(getStartAndFinish(pairSplit[1])),
		})
	}
	return pairs
}

func getStartAndFinish(sectionAssignment string) (int, int) {
	firstS, secondS, _ := strings.Cut(sectionAssignment, "-")

	start, _ := strconv.Atoi(firstS)
	end, _ := strconv.Atoi(secondS)

	return start, end
}

func sequentialInt(start, end int) []int {
	var size int
	// check if end > start
	if end > start {
			size = end - start
			size = size + 1 // because counting starts from 0
			intSlice := make([]int, size)

			for i := 0; i < len(intSlice); i++ {
					intSlice[i] = start + i
			}
			return intSlice

	} else {
			size = start - end
			size = size + 1 // plus 1 because counting starts from 0
			intSlice := make([]int, size)

			for i := 0; i < len(intSlice); i++ {
					intSlice[i] = start - i
			}
			return intSlice
	}
}