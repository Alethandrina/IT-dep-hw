package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
C:/Users/1/GolandProjects/awesomeProject/ДЗ1/file.txt.txt
C:/Users/1/GolandProjects/awesomeProject/ДЗ1/5.txt
*/
func main() {
	var input_file, output_file string
	fmt.Print("Enter your input file: ")
	fmt.Scan(&input_file)
	fmt.Print("Enter your output file: ")
	fmt.Scan(&output_file)
	file, err := os.Open(input_file)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	a := []string{}
	b := []string{}
	for s.Scan() {
		a = append(a, s.Text())
	}
	for idx, el := range a {
		t := true
		for j := 0; j < len(a); j++ {
			if j == idx {
				continue
			}
			if el == a[j] {
				t = false
				break
			}
		}
		if t {
			b = append(b, el)
		}
	}
	for in, el := range b {
		b[in] = strings.ToUpper(el)
	}
	for i := 0; i < len(b)-1; i++ {
		imin := i
		for j := i + 1; j < len(b); j++ {
			if b[j] < b[imin] {
				imin = j
			}
		}
		b[i], b[imin] = b[imin], b[i]
	}
	file1, err := os.Create(output_file)
	if err != nil {
		fmt.Println("Нельзя создать файл", err)
		os.Exit(1)
	}
	defer file1.Close()
	for _, el := range b {
		file1.WriteString(el + " - " + strconv.Itoa(len(el)) + " байт" + "\n")
	}
}
