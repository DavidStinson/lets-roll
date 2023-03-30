package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	dice := flag.Int("d", 6, "The type of dice to roll. Format: X where X is an integer. Defaults to 6.")
	numRoll := flag.Int("n", 1, "The number of die to roll. Defaults to 1")
	sum := flag.Bool("s", false, "Get the sum of all dice rolls")
	advantage := flag.Bool("adv", false, "Roll the dice with advantage")
	disadvantage := flag.Bool("dis", false, "Roll the dice with disadvantage")
	flag.Parse()

	if *advantage && *disadvantage {
		log.Fatalf("cannot roll with both advantage and disadvantange")
	} else {
		fmt.Printf("You chose a d%d\n", *dice)
		rolls := rollDice(dice, numRoll)
		printDice(rolls)
		if *sum {
			sum := sumDice(rolls)
			fmt.Printf("The sum of the dice was %d\n", sum)
		}
		if *advantage {
			roll := rollWithAdvantage(rolls)
			fmt.Printf("Roll with advantage was %d\n", roll)
		}
		if *disadvantage {
			roll := rollWithDisadvantage(rolls)
			fmt.Printf("Roll with disadvantage was %d\n", roll)
		}
	}
}

func rollDice(dice *int, times *int) []int {
	var rolls []int
	for i := 0; i < *times; i++ {
		rolls = append(rolls, rand.Intn(*dice)+1)
	}
	return rolls
}

func printDice(rolls []int) {
	for i, dice := range rolls {
		fmt.Printf("Roll %d was %d\n", i+1, dice)
	}
}

func sumDice(rolls []int) int {
	sum := 0
	for _, dice := range rolls {
		sum += dice
	}
	return sum
}

func rollWithAdvantage(rolls []int) int {
	sort.Ints(rolls)
	return rolls[len(rolls)-1]
}

func rollWithDisadvantage(rolls []int) int {
	sort.Ints(rolls)
	return rolls[0]
}
