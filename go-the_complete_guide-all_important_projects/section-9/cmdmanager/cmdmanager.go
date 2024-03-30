package cmdmanager

import (
	"errors"
	"fmt"
)

type CMDManager struct {
}

func (cm CMDManager) ReadLines() ([]string, error) {
	fmt.Println("please input your prices, confirm by ENTER")
	var prices []string
	for {
		var price string
		fmt.Print("Price:")
		//if unexpected error, check here
		_, err := fmt.Scan(&price)
		if err != nil {
			return nil, errors.New("receiving input failed")
		}
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}
func (cm CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
