package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bitcounts := [12]uint{}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.Trim(line, "\n")
		for i, k := range line {
			if k == '1' {
				bitcounts[i] += 1
			}
		}
	}
	fmt.Printf("%+v\n", bitcounts)
	var gamma uint
	for _, k := range bitcounts {
		gamma = gamma << 1
		if k > 500 {
			gamma += 1
		}
	}
	epsilon := (^gamma) & 0xFFF
	fmt.Printf("%8s : %012b\n%8s : %012b\n%8s : %d\n", "gamma", gamma, "epsilon", epsilon, "answer", gamma*epsilon)
}
