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
func reading() (map[string]int, bool) {
	var input_file string
	fmt.Print("Enter your input file: ")
	fmt.Scan(&input_file)
	file, err := os.Open(input_file)
	if err != nil {
		fmt.Println("Can not find the file", err)
		dict := make(map[string]int)
		return dict, false
	} else {
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
		return dict, true
	}
}

func find_uniq(dict map[string]int) []string {
	a := make([]string, 0, 0)
	for key, value := range dict {
		if value == 1 {
			a = append(a, key)
		}
	}
	return a
}

func up_sort(b []string) []string {
	for in, el := range b {
		b[in] = strings.ToUpper(el)
	}
	sort.Strings(b)
	return b
}
func output(b []string) {
	var output_file string
	fmt.Print("Enter your output file: ")
	fmt.Scan(&output_file)

	file1, err := os.Create(output_file)
	if err != nil {
		fmt.Println("Can not create a file", err)
		return
	}
	defer file1.Close()
	for _, el := range b {
		file1.WriteString(el + " - " + strconv.Itoa(len(el)) + " байт" + "\n")
	}
}
func main() {
	a, flag := reading()
	if flag {
		b := find_uniq(a)
		b = up_sort(b)
		output(b)
	}
}
