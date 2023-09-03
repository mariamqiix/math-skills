package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
	"math"
	"sort"
	"strconv"
)

func main() {
	if !ifItsNo("numbers.txt") {
		fmt.Print("the file contains a text")
		return
	}
	fmt.Println(ReadFile("numbers.txt"))
	x := ReadFile("numbers.txt")
	fmt.Print("the Average : ")
	fmt.Println(math.Round(ava(x)))
	fmt.Print("the Median : ")
	fmt.Println(math.Round(Median(x)))
	fmt.Print("the Variance : ")
	fmt.Println(math.Round(Variance(x)))
	fmt.Print("the Standard Deviation : ")
	fmt.Println(math.Round(StandardDeviation(x)))
}

func ReadFile(fileName string) []int {
	ReadFile, _ := os.Open(fileName)
	FileScanner := bufio.NewScanner(ReadFile)
	var numbers []int
	for FileScanner.Scan() {
		if s, err := strconv.Atoi(FileScanner.Text()); err == nil {
			numbers = append(numbers, s)
		}
	}
	ReadFile.Close()
	return numbers
}

func ifItsNo(text string) bool {
	ReadFile, _ := os.Open(text)
	FileScanner := bufio.NewScanner(ReadFile)
	for FileScanner.Scan() {
		x := FileScanner.Text()
		for _, c := range x {
			if c < '0' || c > '9' {
				return false
			}
		}
	}
	ReadFile.Close()
	return true
}

func ava(i []int) float64 {
	z := 0.0
	for x := 0; x < len(i); x++ {
		z += float64(i[x])
	}
	return (z / float64(len(i)))
}

func Median(i []int) float64 {
	sort.Ints(i)
	Median := 0.0
	if len(i)%2 == 1 {
		Median = float64(i[(len(i) / 2)])
	} else {
		Median = (float64(i[(len(i)/2)]) + float64(i[(len(i)/2)-1]))/2.0
	}
	return Median
}

func Variance(i []int) float64 {
	DeltaX := Median(i)
	TheX := 0.0
	for x := 0; x < len(i); x++ {
		z := float64(i[x]) - float64(DeltaX)
		TheX += (z * z)
	}
	TheX /= (float64(len(i)))
	return TheX
}

func StandardDeviation(i []int) float64 {
	x := Variance(i)
	return (math.Sqrt(x))
}
