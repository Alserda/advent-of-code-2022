package main

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input_day string

// Inventory is a mapping of elf's carrying a total calories.
type Inventory map[int]int

func main() {
	fmt.Println(day01(input_day))
	fmt.Println(day02(input_day))
}

func day01(input string) string {
	inv := makeInventory(input)
	elf, calories, err := inv.highest()

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("elf %d carries the most calories: %d", elf, calories)
}

func day02(input string) string {
	inv := makeInventory(input)

	elfs := 3
	sum := 0
	for _, calories := range inv.order()[0:elfs] {
		sum = sum + calories
	}

	return fmt.Sprintf("the top %d elfs carry %d calories in total", elfs, sum)
}

// makeInventory reads input and assigns the amount of calories per elf
func makeInventory(input string) Inventory {
	total := make(Inventory)

	elf := 1
	for _, v := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		if v == "" {
			elf++
			continue
		}

		amount, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("failed to convert %s to int", v)
		}

		calories, ok := total[elf]

		if !ok {
			total[elf] = amount
		} else {
			total[elf] = calories + amount
		}
	}

	return total
}

// highest returns the elf with the most amount of calories.
// an error is returned if the inventory is empty.
func (inv Inventory) highest() (int, int, error) {
	var highest *int

	for elf, calories := range inv {
		i := elf

		if highest == nil || calories > inv[*highest] {
			highest = &i
		}
	}

	if highest == nil {
		return 0, 0, fmt.Errorf("could not determine a highest value")
	}

	return *highest, inv[*highest], nil
}

// order returns the calories from highest to lowest.
func (inv Inventory) order() []int {
	t := []int{}
	for _, calories := range inv {
		t = append(t, calories)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(t)))

	return t
}
