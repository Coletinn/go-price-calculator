package cmd

import "fmt"

type CMD struct {

}

func (cmd CMD) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMD) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMD {
	return CMD{}
}