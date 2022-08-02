package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var gamma uint   // Output variable
	var epsilon uint // Output variable
	//var oxygen uint      // Output variable
	//var co2 uint // Output variable

	// Read input
	records, bitLength := readInput(bufio.NewReader(os.Stdin))

	// Set bitMask
	var bitMask uint // Mask representing the bitstream width
	for i := uint(0); i < bitLength; i++ {
		bitMask = bitMask | (1 << i)
	}
	fmt.Printf("bitMask set to %017b\n", bitMask)
	bitcounts := countBits(records, bitLength)

	// Calculate gamma
	for i, k := range bitcounts {
		if k >= (uint(len(records)) >> 1) {
			gamma = gamma | (1 << i)
		}
	}

	// Calculate epsilon
	epsilon = (^gamma) & bitMask

	/* o2Rating, err := filterDown(records, gamma, bitLength)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("o2Rating: %d\n", o2Rating)

	co2Rating, err := filterDown(records, epsilon, bitLength)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("co2Rating: %d\n", co2Rating) */

	// Display answers
	fmt.Printf("%8s : %d\n%8s : %016b (%d)\n%8s : %016b (%d)\n%8s : %d\n", "records", len(records), "gamma", gamma, gamma, "epsilon", epsilon, epsilon, "answer1", gamma*epsilon)
}

func filter(in []uint, pred func(uint) bool) []uint {
	var out []uint
	for _, e := range in {
		if pred(e) {
			out = append(out, e)
		}
	}
	return out
}

func filterDown(in []uint, comparison uint, bitLength uint) (uint, error) {
	out := make([]uint, len(in), len(in))
	copy(out, in)

	var i uint
	for i = uint(bitLength - 1); i >= 0; i-- {
		mask := uint(1 << i)
		fmt.Printf("\nBefore checking bit %d, %d items exist in the list\n", i, len(out))
		out = filter(out, func(e uint) bool {
			fmt.Printf("i: %d\ne: %05b\nm: %05b\nc: %05b\nl: %05b\nr: %05b\n\n", i, e, mask, comparison, e&mask, comparison&mask)
			return e&mask == comparison&mask
		})
		fmt.Printf("After checking bit %d, %d items remain in the list\n", i, len(out))
		if len(out) == 1 {
			return out[0], nil
		} else if len(out) == 0 {
			break
		}
	}
	return 0, errors.New("couldn't find a proper thing to return")
}

func countBits(records []uint, bitLength uint) []uint {
	bitcounts := make([]uint, bitLength, bitLength)
	for _, v := range records {
		for i := bitLength - 1; i > 0; i-- {
			if v&(1<<i) != 0 {
				bitcounts[i]++
			}
			if i > bitLength {
				panic(fmt.Errorf("i is like, huge: %d", i))
			}
		}
	}
	return bitcounts
}

func printSlice(i []uint) string {
	out := "[ "
	for _, j := range i {
		out += fmt.Sprintf("%05b ", j)
	}
	out += "]"
	return out
}

func readInput(reader *bufio.Reader) ([]uint, uint) {
	var records []uint
	var bitLength uint

	// Read first line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Errorf("Error reading first line: %s\n", err.Error()))
	}

	// Set the bitLength accounting for the \n at the end
	bitLength = uint(len(line) - 1)
	firstRecord, err := parseLine(line)
	if err != nil {
		panic(fmt.Errorf("Error reading first line: %v", err))
	}

	// Insert the first record into the slice
	records = append(records, firstRecord)

	// Loop through the remaining input and load into memory
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			break // end read on EOF
		}
		record, err := parseLine(line)
		if err != nil {
			// If a line can't be parsed for some reason, just skip it
			continue
		}
		records = append(records, record)
	}
	return records, bitLength
}

func parseLine(line string) (uint, error) {
	line = strings.Trim(line, "\n")
	r, err := strconv.ParseUint(line, 2, 16)
	record := uint(r)
	if err != nil {
		return 0, err
	}
	return record, nil
}
