package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)


func ReadFromFile(fileName string) (float64, error) {

	data, err := os.ReadFile(fileName)


	if(err != nil){
		return 1000,  errors.New("no such file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if(err != nil){
		return 1000, errors.New("no such file")
	}
	return balance, nil
}


func WriteBalanceToFile(balance float64, fileName string) {

	balanceText := fmt.Sprint(balance)

	os.WriteFile(fileName, []byte(balanceText), 0644)
}