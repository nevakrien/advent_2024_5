all: part1 part2

part1:
	go build part1.go common.go 
part2:
	go build part2.go common.go 

clean:
	rm part1 part2

.PHONY: all clean