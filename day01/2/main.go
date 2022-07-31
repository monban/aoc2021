package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var ring []uint64
	var oldDepth uint64
	var count int
	reader := bufio.NewReader(os.Stdin)
	//count := 0

	// Prime the pump with three values
	for i := 0; i < 3; i++ {
		line, _ := reader.ReadString('\n')
		val := readLine(line)
		ring = append(ring, val)
		oldDepth += val
	}

	// Loop through remaining values
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("EOF")
			break
		}
		val := readLine(line)
		ring = append(ring, val)
		ring = ring[1:]
		var depth uint64
		for i := 0; i < 3; i++ {
			depth += ring[i]
		}
		if depth > oldDepth {
			count += 1
		}
		oldDepth = depth
	}
	fmt.Printf("%+v\n", count)
}

func readLine(line string) uint64 {
	line = strings.Trim(line, "\n ")
	if len(line) == 0 {
		return 0
	}
	val, err := strconv.ParseUint(line, 10, 32)
	if err != nil {
		fmt.Printf("error parsing to int: %s, %v\n", line, err)
		return 0
	}
	return val
}
