package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var depth int
	var position int

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.Trim(line, "\n")
		if len(line) == 0 {
			break
		}
		cmdAry := strings.Split(line, " ")
		command := cmdAry[0]
		fmt.Printf("cmdAry: %+v\n", cmdAry)
		distance, _ := strconv.ParseUint(cmdAry[1], 10, 32)
		switch command {
		case "forward":
			position += int(distance)
		case "down":
			depth += int(distance)
		case "up":
			depth -= int(distance)
		}
	}
	fmt.Printf("Position: %d\nDepth: %d\nAnswer: %d\n", position, depth, position*depth)
}
