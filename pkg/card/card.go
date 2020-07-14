package card

import (
	"sync"
	"time"
)

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Transactions []Transaction
}

type Transaction struct {
	Id     int64
	UserId int64
	Amount int64
	Date   time.Time
	MCC    string
	Status string
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, *transaction)
}

func TransactionsByCategory1(transactions []Transaction, userid int64) map[string]int64 {
	result := make(map[string]int64)
	for _, v := range transactions {
		if v.UserId == userid {
			result[v.MCC] += v.Amount
		}
	}
	return result
}

func TransactionsByCategory2(transactions []Transaction, userid int64) map[string]int64 {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	result := make(map[string]int64)

	// choose the partition interval
	x := 3
	if len(transactions)%2 == 0 {
		x = 2
	}

	n := 0
	for n < len(transactions) { // TODO здесь ваши условия разделения
		wg.Add(1)
		part := transactions[n : n+x]
		n += x

		go func() {
			m := TransactionsByCategory1(part, userid)

			mu.Lock()
			for k, v := range m {
				result[k] += v
			}
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	return result
}

func TransactionsByCategory3(transactions []Transaction, userid int64) map[string]int64 {
	result := make(map[string]int64)
	ch := make(chan map[string]int64)

	// choose the partition interval
	x := 3
	if len(transactions)%2 == 0 {
		x = 2
	}

	n := 0
	partsCount := 0
	for n < len(transactions) {
		part := transactions[n : n+x]
		n += x
		partsCount++

		go func(ch chan<- map[string]int64) {
			ch <- TransactionsByCategory1(part, userid)
		}(ch)
	}

	finished := 0
	for  {
		for k, v := range <-ch {
			result[k] += v
		}
		finished++
		if finished == partsCount {
			break
		}
	}
	return result
}

func TransactionsByCategory4(transactions []Transaction, userid int64) map[string]int64 {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	result := make(map[string]int64)

	// choose the partition interval
	x := 3
	if len(transactions)%2 == 0 {
		x = 2
	}

	n := 0
	for n < len(transactions) {
		part := transactions[n : n+x]
		n += x
		wg.Add(1)
		go func() {
			for _, t := range part {
				if t.UserId == userid {
					mu.Lock()
					result[t.MCC] += t.Amount
					mu.Unlock()
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return result

}
