package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var bitcounts []uint // Count of each bit
	var bitLength int    // Total width of incoming bitstream
	var bitMask uint     // Mask representing the bitstream width
	var records uint     // Number of records processed
	var gamma uint       // Output variable
	var epsilon uint     // Output variable

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		line = strings.Trim(line, "\n")
		record, err := strconv.ParseUint(line, 2, 16)
		if err != nil {
			continue
		}
		records++

		if bitLength == 0 {
			fmt.Println("Setup...")
			if err != nil {
				fmt.Printf("Error reading first line: %s\n", err.Error())
				return
			}
			bitLength = len(line)
			bitcounts = make([]uint, len(line), len(line))
			for i := 0; i < bitLength; i++ {
				bitMask = bitMask | (1 << i)
			}
			fmt.Printf("bitLength: %d\nbitMask: %032b\nbitcounts: %+v\n", bitLength, bitMask, bitcounts)
			fmt.Println("Processing...")
		}
		fmt.Printf("Record %04d: %016b\n", records-1, record)
		for i := 0; i < bitLength; i++ {
			if record&(1<<i) != 0 {
				bitcounts[bitLength-i-1] += 1
			}
		}
	}
	fmt.Printf("bitcounts: %+v\n", bitcounts)
	for _, k := range bitcounts {
		gamma = gamma << 1
		if k >= (records >> 1) {
			gamma += 1
		}
	}
	epsilon = (^gamma) & bitMask
	fmt.Printf("%8s : %d\n%8s : %016b\n%8s : %016b\n%8s : %d\n", "records", records, "gamma", gamma, "epsilon", epsilon, "answer1", gamma*epsilon)
}
