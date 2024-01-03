package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalancerFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}
	return balance, nil

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)

}

func main() {

	var accountBalance, err = getBalancerFromFile()
	//fmt.Println(err)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println("error message:", err)
		fmt.Println("--------")
		//panic("cant continue, sorry") or return
	}

	fmt.Println("Welcome to go bank")

	for {

		fmt.Println("what do you want to do ?")
		fmt.Println("1. check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. withdraw money")
		fmt.Println("4. exit")

		var choice int
		fmt.Print("your choice:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("your balance is", accountBalance)
		case 2:
			fmt.Println("how much do you want to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("invalid amount, must be greater than 0")
				//return // stop the execution , no other code its execute
				continue
			}

			accountBalance += depositAmount
			fmt.Println("balance updated ! new amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Println("how much do you want to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("invalid amount, must be greater than 0")
				continue // stop the execution , no other code its execute
			}

			if withdrawAmount > accountBalance {
				fmt.Println("invalid amount, you cant withdraw more than you have")
				continue

			}

			accountBalance -= withdrawAmount
			fmt.Println("balance updated ! new amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("goodbye ! ")
			fmt.Println("thanks for choosing our bank")
			return
			//break

		}
	}

}
