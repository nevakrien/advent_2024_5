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

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

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

		fmt.Printf("a %d b %d\n", a,b)

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

	for {
		for{
			_,err := fmt.Fscanf(reader,"%d", &a)
			if err!=nil {
				fmt.Println("Error parsing a", err)
				os.Exit(1)
			}
			fmt.Printf("num %d \n", a)

			c,err := next_char(reader)
			if err == io.EOF {
				goto end
			}
			if err!=nil {
				fmt.Println("Error parsing char", err)
				os.Exit(1)
			}

			if c=='\n' {
				break
			}
			if c!=',' {
				fmt.Println("missing , seperator found",string(c))
				os.Exit(1)
			}
		}
		

	}

	end:
}