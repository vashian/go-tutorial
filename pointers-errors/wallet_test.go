package main

import "testing"



func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(20))

		assertBalance(t, wallet, Bitcoin(20))

	})

	t.Run("withdraw", func(t *testing.T) {

		wallet := Wallet{Bitcoin(20)}	

		wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))

	}) 
}


// func TestWallet(t *testing.T) {
// 	t.Run("deposit", func(t *testing.T) {

// 		wallet := Wallet{}

// 		wallet.Deposit(Bitcoin(10))

// 		got := wallet.Balance()

// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})

// 	t.Run("withdraw", func(t *testing.T) {

// 		wallet := Wallet{balance: Bitcoin(20)}

// 		wallet.Withdraw(Bitcoin(10))

// 		got := wallet.Balance()

// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})
// }
// func TestWallet(t *testing.T) {

// 	wallet := Wallet{}

// 	wallet.Deposit(Bitcoin(10))

// 	got := wallet.Balane()

// 	fmt.Printf("address of balance in test is %v \n", &wallet.balance)

// 	want := Bitcoin(10)

// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }