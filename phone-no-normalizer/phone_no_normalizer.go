package main

import (
	"fmt"
	"regexp"
)

func phoneNoNormalizer(list []string) {
	result := make(map[string]int)
	regex, err := regexp.Compile("[^0-9]")
	if err != nil {
		fmt.Println("regex is incorrect")
		return
	}

	for _, number := range list {
		phoneNo := regex.ReplaceAllString(number, "")
		result[phoneNo]++
	}

	for number, count := range result {
		fmt.Printf("%v : %v\n", number, count)
	}
}

func main() {
	list := []string{
		"1234567890",
		"123 456 7891",
		"123-456-7890",
		"1234567892",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"(123)456-7892",
	}
	phoneNoNormalizer(list)
}
