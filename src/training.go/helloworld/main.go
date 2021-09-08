package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line
	if strings.Contains(res, old) || strings.Contains(res, oldLower) {
		found = true
		occ += strings.Count(res, old)
		occ += strings.Count(res, oldLower)
		res = strings.Replace(res, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)
	}
	return found, res, occ
}

func FindAndReplace(srcFile, old, new string) (occ int, lines []int, err error) {
	file, err := os.Open(srcFile)
	if err != nil {
		return occ, lines, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lineIdx := 1
	for scanner.Scan() {
		found, res, o := ProcessLine(scanner.Text(), old, new)
		if found {
			occ += o
			lines = append(lines, lineIdx)
		}
		fmt.Println(res)
		lineIdx++
	}
	return occ, lines, nil
}

func main() {
	old := "Go"
	new := "Python"
	occ, lines, err := FindAndReplace("wikigo.txt", old, new)
	if err != nil {
		fmt.Printf("Error while executing find replace: %v\n", err)
	}
	fmt.Println("=== Summary ===")
	defer fmt.Println("=== END Summary ===")
	fmt.Printf("Number of occurences of %v: %v\n", old, occ)
	fmt.Printf("Number of lines: %d\n", len(lines))

	len := len(lines)
	fmt.Print("Lines [ ")
	for i, l := range lines {
		fmt.Print(l)
		if i < len-1 {
			fmt.Print(" - ")
		}
	}
	fmt.Println(" ]")
}
