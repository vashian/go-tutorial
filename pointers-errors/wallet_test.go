package main

import "testing"

func TestWallet(t *testing.T) {
	
	t.Run("deposit", func(t *testing.T) {
		
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		
		wallet := Wallet{Bitcoin(20)}
		
		wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		
		wallet := Wallet{Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})
}


func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {

	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
// func TestWallet(t *testing.T) {

// 	// assertError := func(t testing.TB, err error) {

// 	// 	t.Helper()

// 	// 	if err == nil {

// 	// 		t.Error("wanted an error bud didn't get one")	
	
// 	// 	}
// 	// }

// 	assertError := func(t testing.TB, got error, want string) {

// 		t.Helper()

// 		if got == nil {
// 			t.Fatal("didn't get an error but wanted one")
// 		}

// 		if got.Error() != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	}
// 	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		
// 		t.Helper()

// 		got := wallet.Balance()

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	}

// 	t.Run("deposit", func(t *testing.T) {

// 		wallet := Wallet{}

// 		wallet.Deposit(Bitcoin(20))

// 		assertBalance(t, wallet, Bitcoin(20))

// 	})

// 	t.Run("withdraw", func(t *testing.T) {

// 		wallet := Wallet{Bitcoin(20)}	

// 		wallet.Withdraw(Bitcoin(10))

// 		assertBalance(t, wallet, Bitcoin(10))

// 	}) 

// 	t.Run("withdraw insufficient funds", func(t *testing.T) {
		
// 		startingBalance := Bitcoin(20)

// 		wallet := Wallet{startingBalance}

// 		err := wallet.Withdraw(Bitcoin(100))

// 		assertError(t, err,"cannot withdraw, insufficient funds")
// 		assertBalance(t, wallet, startingBalance)

// 	})
// }	


// // func TestWallet(t *testing.T) {
// // 	t.Run("deposit", func(t *testing.T) {

// // 		wallet := Wallet{}

// // 		wallet.Deposit(Bitcoin(10))

// // 		got := wallet.Balance()

// // 		want := Bitcoin(10)

// // 		if got != want {
// // 			t.Errorf("got %s want %s", got, want)
// // 		}
// // 	})

// // 	t.Run("withdraw", func(t *testing.T) {

// // 		wallet := Wallet{balance: Bitcoin(20)}

// // 		wallet.Withdraw(Bitcoin(10))

// // 		got := wallet.Balance()

// // 		want := Bitcoin(10)

// // 		if got != want {
// // 			t.Errorf("got %s want %s", got, want)
// // 		}
// // 	})
// // }
// // func TestWallet(t *testing.T) {

// // 	wallet := Wallet{}

// // 	wallet.Deposit(Bitcoin(10))

// // 	got := wallet.Balane()

// // 	fmt.Printf("address of balance in test is %v \n", &wallet.balance)

// // 	want := Bitcoin(10)

// // 	if got != want {
// // 		t.Errorf("got %d want %d", got, want)
// // 	}
// // }