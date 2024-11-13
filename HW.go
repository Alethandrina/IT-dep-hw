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

func input(s string) (string, bool) {
	var name string
	fmt.Println(s)
	_, scanerr := fmt.Scan(&name)
	if scanerr != nil {
		fmt.Println("Scan error: ", scanerr)
		return "", false
	}
	return name, true
}
func openFile(name string) bool {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("There's a problem with opening file: ", err)
		return false
	}
	file.Close()
	return true
}

func createFile(name string) bool {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Can not create a file", err)
		return false
	}
	file.Close()
	return true
}
func getDict(inputFile string) map[string]int {
	file, _ := os.Open(inputFile)
	defer file.Close()
	scaner := bufio.NewScanner(file)
	dict := make(map[string]int)
	for scaner.Scan() {
		line := scaner.Text()
		kol, ok := dict[line]
		if ok {
			dict[line] = kol + 1
		} else {
			dict[line] = 1
		}
	}
	return dict
}

func findUniq(dict map[string]int) []string {
	a := make([]string, 0)
	for key, value := range dict {
		if value == 1 {
			a = append(a, key)
		}
	}
	return a
}

func up(b []string) []string {
	for in, el := range b {
		b[in] = strings.ToUpper(el)
	}
	return b
}

func modify(b []string) []string {
	c := b[:]
	for i := 0; i < len(c); i++ {
		c[i] = c[i] + " - " + strconv.Itoa(len(b[i])) + " байт"
	}
	return c
}
func output(output_file string, c []string) {
	file1, _ := os.Create(output_file)
	defer file1.Close()
	for _, el := range c {
		fmt.Fprintf(file1, el)
		fmt.Fprintf(file1, "\n")
	}
}
func main() {
	inputFile, f1 := input("Enter your input file: ")
	if !f1 {
		return
	}
	f1 = openFile(inputFile)
	if !f1 {
		return
	}
	outputFile, f2 := input("Enter your output file: ")
	if !f2 {
		return
	}
	f2 = createFile(outputFile)
	if !f2 {
		return
	}
	a := getDict(inputFile)
	b := findUniq(a)
	b = up(b)
	sort.Strings(b)
	c := modify(b)
	output(outputFile, c)
}
