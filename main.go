package main

import (
	"bufio"
	"fmt"
	"os"
)

func interpretBrainfuck(source string) error {
	memory := make([]byte, 30000)
	ptr := 0
	stack := []int{}

	for pc := 0; pc < len(source); pc++ {
		switch source[pc] {
		case '>':
			ptr++
			if ptr >= len(memory) {
				return fmt.Errorf("memory pointer out of bounds")
			}
		case '<':
			ptr--
			if ptr < 0 {
				return fmt.Errorf("memory pointer out of bounds")
			}
		case '+':
			memory[ptr]++
		case '-':
			memory[ptr]--
		case '.':
			fmt.Printf("%c", memory[ptr])
		case ',':
			input := bufio.NewReader(os.Stdin)
			char, err := input.ReadByte()
			if err != nil {
				return fmt.Errorf("error reading input: %v", err)
			}
			memory[ptr] = char
		case '[':
			if memory[ptr] == 0 {
				loop := 1
				for loop > 0 {
					pc++
					if pc >= len(source) {
						return fmt.Errorf("unmatched '[' at position %d", pc)
					}
					switch source[pc] {
					case '[':
						loop++
					case ']':
						loop--
					}
				}
			} else {
				if len(stack) == 0 || stack[len(stack)-1] != pc {
					stack = append(stack, pc)
				}
			}
		case ']':
			if len(stack) == 0 {
				return fmt.Errorf("unmatched ']' at position %d", pc)
			}
			if memory[ptr] != 0 {
				pc = stack[len(stack)-1] - 1
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) > 0 {
		return fmt.Errorf("unmatched '[' at positions: %v", stack)
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input.bf>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	source, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		os.Exit(1)
	}

	err = interpretBrainfuck(string(source))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error interpreting Brainfuck source: %v\n", err)
		os.Exit(1)
	}
}
