package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) (string, error) {
	fmt.Print(prompt)
	s, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(s), nil
}

func readInt(prompt string) (int, error) {
	s, err := readLine(prompt)
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid integer: %w", err)
	}
	return i, nil
}
