package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func div(numanator, denomator int) (result float64, err error) {
	if denomator == 0 {
		err = errors.New("tumi shuno diye vag korteco")
	} else {
		result = float64(numanator) / float64(denomator)
	}
	return
}
func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID: %d", a, b, os.Getuid())
	}
	return nil
}

func main() {

	r, e := div(1, 0)
	fmt.Println(r, " and ", e.Error())
	err := formattedError(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	i, err := strconv.Atoi("-123")
	if err == nil {
		fmt.Println("Int value is", i)
	}
	i, err = strconv.Atoi("Y123")
	if err != nil {
		fmt.Println(err)
	}
}
