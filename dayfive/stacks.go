package dayfive

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type stack []string
type instruction struct {
	fromStack, toStack, numMoves int
}

var (
	one stack   = []string{"G", "T", "R", "W"}
	two stack   = []string{"G", "C", "H", "P", "M", "S", "V", "W"}
	three stack = []string{"C", "L", "T", "S", "G", "M"}
	four stack  = []string{"J", "H", "D", "M", "W", "R", "F"}
	five stack  = []string{"P", "Q", "L", "H", "S", "W", "F", "J"}
	six stack   = []string{"P", "J", "D", "N", "F", "M", "S"}
	seven stack = []string{"Z", "B", "D", "F", "G", "C", "S", "J"}
	eight stack = []string{"R", "T", "B"}
	nine stack  = []string{"H", "N", "W", "L", "C"}
)

func Play() {
	instructions := readInputFromFile()

	tops := moveStacksAndReturnTops(instructions)

	fmt.Printf("%#v", tops)
}

func moveStacksAndReturnTops(instructions []instruction) []string {
	
}

func prepend(s []string, container string) []string {
	if len(container) != 1 {
		panic("unable to do operation")
	}

	return slices.Insert(s, 0, container)
}

func readInputFromFile() []instruction {
	b, err := os.ReadFile("dayfive/input.txt")
	if err != nil {
		panic("unable to read file")
	}

	s:= strings.Split(string(b), "\n")

	instructions := []instruction{}
	for _, line := range s {
		instructions = append(instructions, parseLine(line))
	}

	return instructions
}

func parseLine(s string) instruction {
	r, _ := regexp.Compile(`\d`)
	found := r.FindAllString(s, 3)

	num, _ :=strconv.Atoi(found[0])
	from, _ :=strconv.Atoi(found[1])
	to, _ :=strconv.Atoi(found[2])

	return instruction{
		fromStack: from,
		toStack: to,
		numMoves: num,
	}
}