package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadTextFileLinesAsString(path string) []string {
	f, err := os.Open(path)

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	strings := make([]string, 0)

	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return strings
}

func ReadTextFileLinesAsInt(path string) []int {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	numbers := make([]int, 0)

	for scanner.Scan() {
		v, errAtoi := strconv.Atoi(scanner.Text())
		if errAtoi != nil {
			log.Fatalln(errAtoi)
		}

		numbers = append(numbers, v)
	}

	return numbers
}

func ReadTextFile(path string) string {
	content, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
