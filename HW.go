package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
C:/Users/1/GolandProjects/awesomeProject/ДЗ1/file.txt.txt
*/

func readFile(inputFile string) ([]string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	a := make([]string, 0)
	defer file.Close()
	scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		line := scaner.Text()
		a = append(a, line)
	}
	return a, nil
}

func findUniq(b []string) []string {
	cnt := len(b)
	dict := make(map[string]int)
	for _, el := range b {
		_, ok := dict[el]
		if !ok {
			dict[el] = 0
		}
		dict[el]++
		if dict[el] == 2 {
			cnt -= 2
		} else if dict[el] > 2 {
			cnt -= 1
		}
	}
	a := make([]string, cnt)
	i := 0
	for key, value := range dict {
		if value == 1 {
			a[i] = key
			i++
		}
	}
	return a
}

func up(b []string) {
	for in, el := range b {
		b[in] = strings.ToUpper(el)
	}
}

func addInfoAboutBytes(b []string) {
	for i, el := range b {
		b[i] = fmt.Sprintf("%s - %d байт", el, len(el))
	}
}
func output(output_file string, c []string) error {
	file1, err := os.Create(output_file)
	if err != nil {
		return err
	}
	defer file1.Close()
	for _, el := range c {
		fmt.Fprintln(file1, el)
	}
	return nil
}
func main() {
	var inputFile, outputFile string
	fmt.Println("Enter your input file:")
	_, f1 := fmt.Scan(&inputFile)
	if f1 != nil {
		fmt.Println("Scan error: ", f1)
		return
	}
	fmt.Println("Enter your  output file:")
	_, f2 := fmt.Scan(&outputFile)
	if f2 != nil {
		fmt.Println("Scan error: ", f2)
		return
	}
	a, err1 := readFile(inputFile)
	if err1 != nil {
		fmt.Println("Read error: ", err1)
		return
	}
	b := findUniq(a)
	up(b)
	sort.Strings(b)
	addInfoAboutBytes(b)
	err2 := output(outputFile, b)
	if err2 != nil {
		fmt.Println("Error while creating a file: ", err2)
		return
	}
}
