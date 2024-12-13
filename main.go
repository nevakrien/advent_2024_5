package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
)

func next_char(r *bufio.Reader) (byte,error){
	c := byte(' ')
	var err error

	for ;c==' '; {
		c,err = r.ReadByte()
		if err!=nil {
			return c,err
		}
	}

	return c,nil
}

func pair_parse(reader *bufio.Reader) [][2]int {
	ans := make([][2]int,0)

	var a,b int
	for {
		_,err := fmt.Fscanf(reader,"%d", &a)
		if err!=nil {
			fmt.Println("Error parsing a", err)
			os.Exit(1)
		}
		c,err := next_char(reader)
		if err!=nil {
			fmt.Println("Error parsing char", err)
			os.Exit(1)
		}

		_,err = fmt.Fscanf(reader,"%d", &b)
		if err!=nil {
			fmt.Println("Error parsing b", err)
			os.Exit(1)
		}

		if c!='|' {
			break
		}

		// fmt.Printf("a %d b %d\n", a,b)
		ans = append(ans,[2]int{a,b})

		c,err = next_char(reader)
		if err!=nil {
			fmt.Println("Error parsing char", err)
			os.Exit(1)
		}

		if c!='\n' {
			fmt.Println("missing newline seperator")
			os.Exit(1)
		}

		c,err = next_char(reader)
		if err!=nil {
			fmt.Println("Error parsing char", err)
			os.Exit(1)
		}

		if c=='\n' {
			break
		} else {
			reader.UnreadByte()
		}
	}

	return ans
}

func inputs_parse_chan(reader *bufio.Reader) <-chan []int {
	// Create a channel to stream parsed lines
	ch := make(chan []int)

	go func() {
		defer close(ch)
		for {
			cur := make([]int, 0)
			var a int

			for {
				// Read integers
				_, err := fmt.Fscanf(reader, "%d", &a)
				if err != nil {
					if err == io.EOF {
						if len(cur) > 0 {
							ch <- cur
						}
						return
					}
					fmt.Println("Error parsing integer:", err)
					os.Exit(1)
				}
				cur = append(cur, a)

				// Read the next character
				c, err := next_char(reader)
				if err == io.EOF {
					ch <- cur
					return
				}
				if err != nil {
					fmt.Println("Error parsing character:", err)
					os.Exit(1)
				}

				// Handle line breaks and separators
				if c == '\n' {
					break
				}
				if c != ',' {
					fmt.Println("Missing ',' separator, found:", string(c))
					os.Exit(1)
				}
			}
			// Yield the current line
			ch <- cur
		}
	}()

	return ch
}

// Helper to index pairs by a and b
func indexPairs(pairs [][2]int) (map[int][]int, map[int][]int) {
    byA := make(map[int][]int)
    byB := make(map[int][]int)

    for _, pair := range pairs {
        a, b := pair[0], pair[1]

        // Index by a
        byA[a] = append(byA[a], b)

        // Index by b
        byB[b] = append(byB[b], a)
    }

    return byA, byB
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
	byA, _ := indexPairs(pairs)

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