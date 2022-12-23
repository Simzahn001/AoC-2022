package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	//create filescanner
	readfile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	filescanner := bufio.NewScanner(readfile)
	filescanner.Split(bufio.ScanLines)

	//parse the input

	items := [][]*big.Int{}
	operations := []string{}
	tests := []int{}
	ifTrue := []int{}
	ifFalse := []int{}

	//a counter which monkey is the current
	counter := 0
	for filescanner.Scan() {
		line := filescanner.Text()

		//get which line it is
		//monkey: go to the next monkey
		if strings.Contains(line, "Monkey") {
			counter++
			//starting items: parse and add the starting items
		} else if strings.Contains(line, "Starting items") {
			//split string to the numbers
			line = line[17:]
			line = strings.Replace(line, " ", "", -1)
			splitedNumbers := strings.Split(line, ",")

			currentMonkeyItems := []*big.Int{}
			//loop through all the starting items and convert them into ints
			for _, string := range splitedNumbers {
				number, err := strconv.Atoi(string)
				if err != nil {
					panic(err)
				}
				currentMonkeyItems = append(currentMonkeyItems, big.NewInt(int64(number)))
			}
			items = append(items, currentMonkeyItems)
		} else if strings.Contains(line, "Operation") {
			operations = append(operations, line[13:])
		} else if strings.Contains(line, "Test") {
			//get the test division number
			number, err := strconv.Atoi(line[21:])
			if err != nil {
				panic(err)
			}
			tests = append(tests, number)
		} else if strings.Contains(line, "If true") {
			//get to which monkey the item is thrown to
			number, err := strconv.Atoi(line[29:])
			if err != nil {
				panic(err)
			}
			ifTrue = append(ifTrue, number)
		} else if strings.Contains(line, "If false") {
			//get to which monkey the item is thrown to
			number, err := strconv.Atoi(line[30:])
			if err != nil {
				panic(err)
			}
			ifFalse = append(ifFalse, number)
		}
	}

	inspections := []int{}

	//simulate rounds

	for round := 0; round < 10000; round++ {

		//every monkey takes one turn per round
		for monkey, monkeyItems := range items {

			// initialize inspections if first round
			if round == 0 {
				inspections = append(inspections, 0)
			}

			//the monkey inspects every item
			for _, worryLevel := range monkeyItems {

				//increase inspections by one
				inspections[monkey]++

				//perform operation
				if operations[monkey][10:11] == "+" {
					summand, err := strconv.Atoi(operations[monkey][12:])
					if err != nil {
						panic(err)
					}
					worryLevel.Add(worryLevel, big.NewInt(int64(summand)))
				} else if operations[monkey][10:11] == "*" {

					//check if operation: "new = old * old"
					if operations[monkey][12:] == "old" {
						worryLevel.Mul(worryLevel, worryLevel)
					} else {
						factor, err := strconv.Atoi(operations[monkey][12:])
						if err != nil {
							panic(err)
						}
						worryLevel.Mul(worryLevel, big.NewInt(int64(factor)))
					}

				} else {
					panic("\"+\" or \"*\" needed!")
				}

				//run test
				worryLevelCopy := *worryLevel
				if worryLevelCopy.Mod(worryLevel, big.NewInt(int64(tests[monkey]))).BitLen() == 0 {
					items[ifTrue[monkey]] = append(items[ifTrue[monkey]], worryLevel)
				} else {
					items[ifFalse[monkey]] = append(items[ifFalse[monkey]], worryLevel)
				}
				//remove the first item from the current monkey
				items[monkey] = items[monkey][1:]
			}

		}

		//debugging
		//because counting starts at zero -> first round = 0 | 20th round = 19
		if (round+1)%1000 == 0 || round == 19 || round == 0 {
			printInspections(inspections, round)
		}

	}

	//get the most busy monkey
	firstMonkey, index := getMax(inspections)
	//delete the highest value
	inspections[index] = inspections[len(inspections)-1]
	inspections = inspections[:len(inspections)-1]
	secondMonkey, _ := getMax(inspections)

	fmt.Println(firstMonkey * secondMonkey)

}

// for debugging purposes
func printInspections(inspections []int, round int) {
	fmt.Println("")
	fmt.Println("== Round:  ", round, " ==")
	for i := 0; i < len(inspections); i++ {
		fmt.Println("Monkey ", i, ": ", inspections[i])
	}
}

// returns number and index
func getMax(array []int) (max int, index int) {

	max = 0
	index = 0

	if len(array) == 0 {
		return
	}

	for x, i := range array {
		if i > max {
			max = i
			index = x
		}
	}

	return
}
