package main

import (
	"bufio"
	"fmt"
	"os"
)



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

	// fmt.Println("A:",byA,"B:",byB)

	inputs := inputs_parse_chan(reader)
	
	total := 0
	for line := range inputs{
		// fmt.Println(line)
		passed := make(map[int]bool)
		for _, x := range line {
			for _, y := range byA[x]{ 
				if passed[y]{
					goto bad
				}
			}
			passed[x] =true

		}
		total+=line[len(line)/2]
	bad:
	}
	fmt.Println("total",total)
}