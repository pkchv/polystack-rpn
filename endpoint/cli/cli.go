package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var stream = os.Stdin
var scanner = bufio.NewScanner(os.Stdin)

func init() {
	scanner.Split(bufio.ScanLines)
}

func getIntFromStdin() (int, error) {
	scanner.Scan()
	input := scanner.Text()
	err := scanner.Err()

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(input)
}

func getStringFromStdin() (string, error) {
	scanner.Scan()
	return scanner.Text(), scanner.Err()
}

func parseInt(options *ParseOptions) int {
	inputValue, err := getIntFromStdin()

	for err != nil {

		if options.LogErrors {
			fmt.Fprintln(os.Stderr, err)
		}

		inputValue, err = getIntFromStdin()
	}

	return inputValue
}

func parseString(options *ParseOptions) string {
	inputValue, err := getStringFromStdin()

	for err != nil {

		if options.LogErrors {
			fmt.Fprintln(os.Stderr, err)
		}

		inputValue, err = getStringFromStdin()
	}

	return inputValue
}

func ParseInput(options *ParseOptions) []string {
	numberOfComputations := parseInt(options)

	for numberOfComputations > options.BufferSize {

		if options.LogErrors {
			fmt.Fprintf(os.Stderr, "Number of expressions cannot exceed %d\n", options.BufferSize)
		}

		numberOfComputations = parseInt(options)
	}

	computations := make([]string, 0)

	for i := 0; i < numberOfComputations; i++ {
		computation := parseString(options)
		computations = append(computations, computation)
	}

	fmt.Println()

	return computations
}
