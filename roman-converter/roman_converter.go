package main

import "fmt"

type romanSymbol struct {
	symbol string
	number int
}

func romanConverter(number int) string {
	result := ""

	if number < 1 || number > 100 {
		fmt.Println("the number must be between 1 and 100")
		return result
	}

	romans := []romanSymbol{
		romanSymbol{"C", 100},
		romanSymbol{"XC", 90},
		romanSymbol{"L", 50},
		romanSymbol{"XL", 40},
		romanSymbol{"X", 10},
		romanSymbol{"IX", 9},
		romanSymbol{"V", 5},
		romanSymbol{"IV", 4},
		romanSymbol{"I", 1},
	}

	for i := 0; i < len(romans) || number > 0; i++ {
		if number < romans[i].number {
			continue
		}

		result += romans[i].symbol
		number -= romans[i].number
		i--
	}

	return result
}

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%v ", romanConverter(i))
		if i%10 == 0 {
			fmt.Println()
		}
	}
}
