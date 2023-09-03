package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
	fileName := os.Args[1] 
	x := ReadFile(fileName)
	fmt.Print("the Average : ")
	fmt.Println(math.Round(ava(x)))
	fmt.Print("the Median : ")
	fmt.Println(math.Round(Median(x)))
	fmt.Print("the Variance : ")
	fmt.Print(int(math.Round(Variance(x))))
	fmt.Println()
	fmt.Print("the Standard Deviation : ")
	fmt.Println(math.Round(StandardDeviation(x)))
} else {
	fmt.Println("you should write \ngo run . FileName")
}
}

func ReadFile(fileName string) []int {
	ReadFile, error := os.Open(fileName)
	if error != (nil) {
		printError(error)
	}
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
	DeltaX := ava(i)
	TheX := 0.0
	for x := 0; x < len(i); x++ {
		z := (float64(i[x]) - float64(DeltaX))
		TheX += (z * z)
	}
	Variance := TheX/(float64(len(i)))
	return Variance
}

func StandardDeviation(i []int) float64 {
	x := Variance(i)
	return (math.Sqrt(x))
}

func printError(err error) {
	fmt.Println("ERROR: " + err.Error())
	os.Exit(1)
}