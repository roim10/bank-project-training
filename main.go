package main

import "fmt"

func main() {
	acc1 := Account{
		ID:      "1",
		Owner:   "982349034",
		balance: 100,
	}
	i, err := acc1.Deposit(100)
	fmt.Println(i)
	fmt.Println(err)
	i, err = acc1.Withdraw(20)
	fmt.Println(i)
	fmt.Println(err)
	i = acc1.GetBalance()
	fmt.Println(i)
}
