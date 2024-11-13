package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
C:/Users/1/GolandProjects/awesomeProject/ДЗ1/file.txt.txt
*/

func readFile(inputFile string) []string {
	file, _ := os.Open(inputFile)
	a := make([]string, 0)
	defer file.Close()
	scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		line := scaner.Text()
		a = append(a, line)
	}
	fmt.Println(a)
	return a
}

func findUniq(b []string) []string {
	a := make([]string, 0)
	dict := make(map[string]int)
	for i := 0; i < len(b); i++ {
		kol, ok := dict[b[i]]
		if ok {
			dict[b[i]] = kol + 1
		} else {
			dict[b[i]] = 1
		}
	}
	for key, value := range dict {
		if value == 1 {
			a = append(a, key)
		}
	}
	return a
}

func up(b []string) {
	for in, el := range b {
		b[in] = strings.ToUpper(el)
	}
}

func addInfo(b []string) {
	for i := 0; i < len(b); i++ {
		b[i] = b[i] + " - " + strconv.Itoa(len(b[i])) + " байт"
	}
}
func output(output_file string, c []string) {
	file1, _ := os.Create(output_file)
	defer file1.Close()
	for _, el := range c {
		fmt.Fprintln(file1, el)
	}
}
func main() {
	var inputFile, outputFile string
	fmt.Println("Enter your input file:")
	_, f1 := fmt.Scan(&inputFile)
	if f1 != nil {
		fmt.Println("Scan error: ", f1)
		return
	}
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("There's a problem with opening file: ", err)
		return
	}
	file.Close()
	fmt.Println("Enter your  output file:")
	_, f2 := fmt.Scan(&outputFile)
	if f2 != nil {
		fmt.Println("Scan error: ", f2)
		return
	}
	file1, err1 := os.Create(outputFile)
	if err1 != nil {
		fmt.Println("There's a problem with creating file: ", err1)
		return
	}
	file1.Close()
	a := readFile(inputFile)
	b := findUniq(a)
	up(b)
	sort.Strings(b)
	addInfo(b)
	output(outputFile, b)
}
