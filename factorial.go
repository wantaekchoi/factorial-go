package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func factorial(n int64) *big.Int {
	fact := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		fact.Mul(fact, big.NewInt(i))
	}
	return fact
}

func main() {
	var num int64
	var base int

	flag.Usage = func() {
		printUsage()
	}

	flag.Int64Var(&num, "number", 0, "Integer value to calculate factorial")
	flag.Int64Var(&num, "n", 0, "Integer value to calculate factorial")
	flag.IntVar(&base, "base", 10, "Number base for the input (default: 10)")
	flag.IntVar(&base, "b", 10, "Number base for the input (default: 10)")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		var err error
		num, err = strconv.ParseInt(args[0], base, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid input. Please enter a valid integer.")
			printUsage()
			return
		}
	} else if num == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter an integer: ")
		scanner.Scan()
		numStr := scanner.Text()

		var err error
		num, err = strconv.ParseInt(numStr, base, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid input. Please enter a valid integer.")
			printUsage()
			return
		}
	}

	if num < 0 {
		fmt.Fprintln(os.Stderr, "Negative numbers are not allowed.")
		return
	}

	result := factorial(num)
	fmt.Printf("%d! = %s\n", num, result.String())
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] <number>\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "Calculate the factorial of the given integer.")
	fmt.Fprintln(os.Stderr, "Options:")
	fmt.Fprintln(os.Stderr, "  -number, -n   Integer value to calculate factorial")
	fmt.Fprintln(os.Stderr, "  -base, -b     Number base for the input (default: 10)")
}
