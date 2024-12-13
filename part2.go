package main

import (
	"bufio"
	"fmt"
	"os"
)

func fixup(byA map[int][]int, line []int) bool {

	modified := false

loop_start:
	passed := make(map[int] int)
	for i, x := range line {
		for _, y := range byA[x]{ 
			j := passed[y] - 1 
			if j >= 0{
				line[i],line[j] = line[j],line[i]
				modified = true
				goto loop_start
			}
		}
		passed[x] =i+1

	}

	return modified
}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	pairs := pair_parse(reader)
	byA := indexPairs(pairs)

	inputs := inputs_parse_chan(reader)
	
	total := 0
	for line := range inputs{
		// fmt.Println(line)
		if fixup(byA,line) {
			total+=line[len(line)/2]
		}

	}
	fmt.Println("total",total)
}