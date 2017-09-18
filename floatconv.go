package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var ErrInvalidNumber = errors.New("the value is not a number")

func main() {
	if len(os.Args) < 2 {
		log.Fatal("One arg is required for this program")
	}

	msg, err := checkNumber(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(msg)
}

func checkNumber(value string) (string, error) {
	fValue, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return "", ErrInvalidNumber
	}

	return fmt.Sprintf("The value <%f> is a valid number", fValue), nil
}
