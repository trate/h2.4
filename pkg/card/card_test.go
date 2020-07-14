package card

import (
	"reflect"
	"testing"
	"time"
)

func makeTransactions() []Transaction {
	const users = 100
	const transactionsPerUser = 6
	const transactionAmount = 1_00
	transactions := make([]Transaction, users*transactionsPerUser)
	for index := range transactions {
		switch index % 100 {
		case 0:
			transactions[index] = Transaction{
				Id:     100,
				UserId: 1,
				Amount: transactionAmount,
				Date:   time.Date(2020, 5, 5, 12, 0, 0, 0, time.UTC),
				MCC:    "5411",
				Status: "InProgress",
			} // Например, каждая 100-ая транзакция в банке от нашего юзера в категории такой-то
		case 2:
			transactions[index] = Transaction{
				Id:     120,
				UserId: 1,
				Amount: transactionAmount,
				Date:   time.Date(2020, 6, 5, 12, 0, 0, 0, time.UTC),
				MCC:    "5812",
				Status: "Done",
			} // Например, каждая 120-ая транзакция в банке от нашего юзера в категории такой-то
		default:
			transactions[index] = Transaction{
				Id:     300,
				UserId: 5,
				Amount: transactionAmount,
				Date:   time.Date(2020, 7, 5, 12, 0, 0, 0, time.UTC),
				MCC:    "5811",
				Status: "Done",
			} // Транзакции других юзеров, нужны для "общей" массы
		}
	}
	return transactions
}

func TestTransactionsByCategory1(t *testing.T) {
	type args struct {
		transactions []Transaction
		userid       int64
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		{name: "Тестируем функцию 1", args: args{
			transactions: makeTransactions(),
			userid:       1,
		}, want: map[string]int64{"5411": 600, "5812": 600}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransactionsByCategory1(tt.args.transactions, tt.args.userid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionsByCategory1() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTransactionsByCategory2(t *testing.T) {
	type args struct {
		transactions []Transaction
		userid       int64
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		{name: "Тестируем функцию 2", args: args{
			transactions: makeTransactions(),
			userid:       1,
		}, want: map[string]int64{"5411": 600, "5812": 600}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransactionsByCategory2(tt.args.transactions, tt.args.userid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionsByCategory2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionsByCategory3(t *testing.T) {
	type args struct {
		transactions []Transaction
		userid       int64
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		{name: "Тестируем функцию 3", args: args{
			transactions: makeTransactions(),
			userid:       1,
		}, want: map[string]int64{"5411": 600, "5812": 600}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransactionsByCategory3(tt.args.transactions, tt.args.userid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionsByCategory3() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTransactionsByCategory4(t *testing.T) {
	type args struct {
		transactions []Transaction
		userid       int64
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		{name: "Тестируем функцию 4", args: args{
			transactions: makeTransactions(),
			userid:       1,
		}, want: map[string]int64{"5411": 600, "5812": 600}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransactionsByCategory4(tt.args.transactions, tt.args.userid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionsByCategory4() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Сокращение GoLand: bench + Tab
func BenchmarkTransactionsByCategory1(b *testing.B) {
	transactions := makeTransactions()
	want := map[string]int64{
		"5411": 600,
		"5812": 600,
	}
	b.ResetTimer() // сбрасываем таймер, т.к. сама генерация транзакций достаточно ресурсоёмка
	for i := 0; i < b.N; i++ {
		result := TransactionsByCategory1(transactions,1)
		b.StopTimer() // останавливаем таймер, чтобы время сравнения не учитывалось
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer() // продолжаем работу таймера
	}
}
func BenchmarkTransactionsByCategory2(b *testing.B) {
	transactions := makeTransactions()
	want := map[string]int64{
		"5411": 600,
		"5812": 600,
	}
	b.ResetTimer() // сбрасываем таймер, т.к. сама генерация транзакций достаточно ресурсоёмка
	for i := 0; i < b.N; i++ {
		result := TransactionsByCategory2(transactions,1)
		b.StopTimer() // останавливаем таймер, чтобы время сравнения не учитывалось
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer() // продолжаем работу таймера
	}
}
func BenchmarkTransactionsByCategory3(b *testing.B) {
	transactions := makeTransactions()
	want := map[string]int64{
		"5411": 600,
		"5812": 600,
	}
	b.ResetTimer() // сбрасываем таймер, т.к. сама генерация транзакций достаточно ресурсоёмка
	for i := 0; i < b.N; i++ {
		result := TransactionsByCategory3(transactions,1)
		b.StopTimer() // останавливаем таймер, чтобы время сравнения не учитывалось
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer() // продолжаем работу таймера
	}
}
func BenchmarkTransactionsByCategory4(b *testing.B) {
	transactions := makeTransactions()
	want := map[string]int64{
		"5411": 600,
		"5812": 600,
	}
	b.ResetTimer() // сбрасываем таймер, т.к. сама генерация транзакций достаточно ресурсоёмка
	for i := 0; i < b.N; i++ {
		result := TransactionsByCategory4(transactions,1)
		b.StopTimer() // останавливаем таймер, чтобы время сравнения не учитывалось
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer() // продолжаем работу таймера
	}
}
