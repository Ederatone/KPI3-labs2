package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	lab2 "github.com/Ederatone/KPI3-labs2"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute(prefix string) error {
	infix, err := lab2.PrefixToInfix(prefix)
	if err != nil {
		return err
	}
	_, err = ch.Output.Write([]byte(infix))
	return err
}

func main() {
	var expression string
	var inputFile string
	var outputFile string

	flag.StringVar(&expression, "e", "", "Expression to compute")
	flag.StringVar(&inputFile, "f", "", "File with expression")
	flag.StringVar(&outputFile, "o", "", "File to output result (optional)")
	flag.Parse()

	if expression != "" && inputFile != "" {
		fmt.Fprintln(os.Stderr, "Error: Both -e and -f flags provided. Use only one.")
		os.Exit(1)
	}

	var inputReader io.Reader
	if inputFile != "" {
		file, err := os.Open(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		inputReader = file
	} else if expression != "" {
		inputReader = strings.NewReader(expression)
	} else {
		fmt.Fprintln(os.Stderr, "Error: No expression provided. Use -e or -f flag.")
		os.Exit(1)
	}

	var outputWriter io.Writer = os.Stdout
	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		outputWriter = file
	}

	inputBytes, err := ioutil.ReadAll(inputReader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %s\n", err)
		os.Exit(1)
	}
	prefixExpression := string(inputBytes)

	handler := ComputeHandler{
		Input:  inputReader,
		Output: outputWriter,
	}

	err = handler.Compute(strings.TrimSpace(prefixExpression))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Syntax Error: %s\n", err)
		os.Exit(1)
	}
}
