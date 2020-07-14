package main

import (
	"fmt"
	"github.com/trate/h2.4/pkg/card"
	"time"
)

func main() {
	const sep = "--------------------------------------------"
	t1 := &card.Transaction{
		Id:     1,
		UserId:     1,
		Amount: -735_55,
		Date:   time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
		MCC:    "5411",
		Status: "InProgress",
	}
	t2 := &card.Transaction{
		Id:     2,
		UserId:     1,
		Amount: -736_55,
		Date:   time.Date(2020, 3, 2, 0, 0, 0, 0, time.UTC),
		MCC:    "5411",
		Status: "InProgress",
	}
	t3 := &card.Transaction{
		Id:     3,
		UserId:     1,
		Amount: 2_000_00,
		Date:   time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC),
		MCC:    "0000",
		Status: "Done",
	}
	t4 := &card.Transaction{
		Id:     4,
		UserId:     1,
		Amount: 2_100_00,
		Date:   time.Date(2020, 4, 2, 0, 0, 0, 0, time.UTC),
		MCC:    "0000",
		Status: "Done",
	}
	t5 := &card.Transaction{
		Id:     5,
		UserId:     1,
		Amount: -1_203_91,
		Date:   time.Date(2020, 5, 1, 0, 0, 0, 0, time.UTC),
		MCC:    "5812",
		Status: "InProgress",
	}
	t6 := &card.Transaction{
		Id:     6,
		UserId:     1,
		Amount: -1_204_91,
		Date:   time.Date(2020, 5, 2, 0, 0, 0, 0, time.UTC),
		MCC:    "5812",
		Status: "InProgress",
	}
	transactions := []card.Transaction{*t1, *t2, *t3, *t4, *t5}

	master := &card.Card{
		Id:           1,
		Issuer:       "MasterCard",
		Balance:      65_000,
		Currency:     "RUB",
		Number:       "5177827685644009",
		Transactions: transactions,
	}

	card.AddTransaction(master, t6)
	fmt.Println("Выводим исходную структуру..")
	fmt.Println(master)


	fmt.Println(sep)
	trByCategory1 := card.TransactionsByCategory1(master.Transactions, 1)
	fmt.Println("Выводим карту транзакций, разбитую по категориям (функция 1)")
	fmt.Println(trByCategory1)

	fmt.Println(sep)
	trByCategory2 := card.TransactionsByCategory2(master.Transactions, 1)
	fmt.Println("Выводим карту транзакций, разбитую по категориям (функция 2)")
	fmt.Println(trByCategory2)

	fmt.Println(sep)
	trByCategory3 := card.TransactionsByCategory3(master.Transactions, 1)
	fmt.Println("Выводим карту транзакций, разбитую по категориям (функция 3)")
	fmt.Println(trByCategory3)


	fmt.Println(sep)
	trByCategory4 := card.TransactionsByCategory3(master.Transactions, 1)
	fmt.Println("Выводим карту транзакций, разбитую по категориям (функция 4)")
	fmt.Println(trByCategory4)
}
