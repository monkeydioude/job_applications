package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// entities

// TestCase contains the amount of element and the list
// of elements that will be squared then displayed
type TestCase struct {
	X   uint
	Yns []int64
}

// InputData carries over data used by the program
type InputData struct {
	N     uint
	Tests []TestCase
	// Rather store the scanner than make a new instance before each read.
	// This work better this way and covers more situations, like data piping
	// e.g.: go run main.go < file1.txt
	Scanner *bufio.Scanner
}

/* utils */

// readInputLine reads a single line from stdin and returns it
func readInputLine(scanner *bufio.Scanner) (string, error) {
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", nil
}

// readLineAndParseInt is a wrapper of readInputLine.
// ParseInt on readInputLine's result
func readLineAndParseInt(scanner *bufio.Scanner) (int64, error) {
	line, err := readInputLine(scanner)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(line, 10, 64)
}

// parseYns is a recursive function that will ParseInt
// up to 'x' element from 'str_list', and append them
// to 'yns'
func parseYns(str_list []string, yns *[]int64, x, it uint) {
	// loop stopper
	if yns == nil || int(it) >= len(str_list) || it >= x {
		return
	}
	res, err := strconv.ParseInt(str_list[it], 10, 64)
	if err != nil {
		return
	}
	(*yns)[it] = res
	parseYns(str_list, yns, x, it+1)
}

func computeTestCase(yns []int64, tot int64, x, it uint) int64 {
	if it >= x || it < 0 {
		return tot
	}
	if yns[it] > 0 {
		tot = tot + (yns[it] * yns[it])
	}
	return computeTestCase(yns, tot, x, it+1)
}

/* services */

// readN parse and store the N parameter from stdin
func readN() (*InputData, error) {
	scanner := bufio.NewScanner(os.Stdin)
	test_nb, err := readLineAndParseInt(scanner)
	if err != nil {
		return nil, err
	}
	if test_nb < 1 || test_nb > 100 {
		return nil, errors.New("N: invalid range (1 <= N <= 100)")
	}
	return &InputData{
		N:       uint(test_nb),
		Tests:   make([]TestCase, test_nb),
		Scanner: scanner,
	}, nil
}

// readXAndInitYns will parse the X parameter of each testCase
// and create a Yns of X amount of element
func readXAndInitYns(scanner *bufio.Scanner) (TestCase, error) {
	x, err := readLineAndParseInt(scanner)
	if err != nil {
		return TestCase{}, err
	}
	if x < 1 || x > 100 {
		return TestCase{}, errors.New("X: invalid range (0 < X <= 100)")
	}
	testCase := TestCase{
		X:   uint(x),
		Yns: make([]int64, x),
	}
	return testCase, nil
}

// readXAndYns is a recursive function that handles the parsing
// of both X and Yns
func readXAndYns(inputData *InputData, it uint) {
	// loop stopper
	if inputData == nil || it >= inputData.N || it < 0 {
		return
	}
	testCase, err := readXAndInitYns(inputData.Scanner)
	if err == nil {
		readYns(&testCase, inputData.Scanner)
		// push in already allocated array
		inputData.Tests[it] = testCase
	}
	readXAndYns(inputData, it+1)
}

// readYns will parse raw Yns from stdin and append them in
// Yns, a x sized, []int64 list
func readYns(testCase *TestCase, scanner *bufio.Scanner) error {
	if testCase == nil {
		return errors.New("nil testCase")
	}
	line, err := readInputLine(scanner)
	if err != nil {
		return err
	}
	parseYns(
		strings.Split(line, " "),
		&testCase.Yns,
		testCase.X,
		0,
	)
	return nil
}

// displayResult will write on stdou the result of each testcase
func displayResult(testCases []TestCase, n, it uint) {
	if it >= n || it < 0 {
		return
	}
	fmt.Println(computeTestCase(testCases[it].Yns, 0, testCases[it].X, 0))
	displayResult(testCases, n, it+1)
}

func main() {
	inputData, err := readN()
	if err != nil {
		// from the challenge's description:
		// Note: There should be no output until all the input has been received.
		// Note 3: Take input from standard input, and output to standard output.
		//
		// Since the program is auto-tested, I didn't want to confuse the tester
		// with error messages. So I commentxed it.
		// fmt.Print64ln(err)
		return
	}
	if inputData == nil || inputData.N == 0 {
		return
	}
	readXAndYns(inputData, 0)
	displayResult(inputData.Tests, inputData.N, 0)
}
