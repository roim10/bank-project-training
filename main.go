package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

type ResultOneTransaction struct {
	fromID string
	toID   string
	amount float64
	Err    error
}

func main() {
	var wg sync.WaitGroup
	bs := NewBankStorage()
	for i := range 10 {
		acc := &Account{
			ID:      strconv.Itoa(i),
			Owner:   fmt.Sprintf("User %d", i),
			balance: 1000,
		}
		bs.AddAccount(acc)
	}
	wg.Add(100)
	result := make(chan ResultOneTransaction, 100)
	for i1 := 0; i1 < 100; i1++ {
		go func() {
			defer wg.Done()
			from := rand.Intn(10)
			to := rand.Intn(10)
			amount := float64(rand.Intn(106) - 5)
			err := bs.Transfer(strconv.Itoa(from), strconv.Itoa(to), amount)
			result <- ResultOneTransaction{
				fromID: strconv.Itoa(from),
				toID:   strconv.Itoa(to),
				amount: float64(amount),
				Err:    err,
			}
		}()
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	successCount := 0
	failCount := 0
	for res := range result {
		if res.Err != nil {
			failCount++
		} else {
			successCount++
		}
	}
	fmt.Printf("Успешно: %d, Ошибок: %d\n", successCount, failCount)
	totalBalance := bs.TotalBalance()
	fmt.Println(totalBalance)
}
