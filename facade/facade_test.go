package facade_test

import (
	"design-pattern/facade"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Helper()
	walletFacade := facade.NewWalletFacade("abc", 1234)

	err := walletFacade.AddMoneyToWallet("cde", 1234, 10)
	assert.EqualError(t, err, "Account Name is incorrect")

	err = walletFacade.AddMoneyToWallet("abc", 234, 10)
	assert.EqualError(t, err, "Security Code is incorrect")
}

func TestDel(t *testing.T) {
	walletFacade := facade.NewWalletFacade("abc", 1234)

	err := walletFacade.DeductMoneyFromWallet("cde", 1234, 10)
	assert.EqualError(t, err, "Account Name is incorrect")

	err = walletFacade.DeductMoneyFromWallet("abc", 234, 10)
	assert.EqualError(t, err, "Security Code is incorrect")

	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 50)
	assert.EqualError(t, err, "Balance is not sufficient")
}

func Example() {
	fmt.Println()
	walletFacade := facade.NewWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	// Output:
	// Starting create account
	// Account created
	//
	// Starting add money to wallet
	// Account Verified
	// SecurityCode Verified
	// Wallet balance added successfully
	// Sending wallet credit notification
	// Make ledger entry for accountId abc with txnType credit for amount 10
	//
	// Starting debit money from wallet
	// Account Verified
	// SecurityCode Verified
	// Wallet balance is Sufficient
	// Sending wallet debit notification
	// Make ledger entry for accountId abc with txnType debit for amount 5
}
