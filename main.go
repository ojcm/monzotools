package monzoutils

import (
	"errors"
	"fmt"

	monzo "github.com/tjvr/go-monzo"
)

// GetAccountsFromAccessToken fetches all accounts for the given access
// token.
func GetAccountsFromAccessToken(accessToken string) []*monzo.Account {
	cl := monzo.Client{
		BaseURL:     "https://api.monzo.com",
		AccessToken: accessToken,
	}
	accounts, _ := cl.Accounts("uk_retail")
	return accounts
}

// GetFirstAccountFromAccessToken gets the first returned account for the
// given access token.
func GetFirstAccountFromAccessToken(accessToken string) monzo.Account {
	return *GetAccountsFromAccessToken(accessToken)[0]
}

// GetFirstAccountIDFromAccessToken gets the ID of the first returned
// account for the given access token.
func GetFirstAccountIDFromAccessToken(accessToken string) string {
	return GetFirstAccountFromAccessToken(accessToken).ID
}

// GetClient returns a client for interracting with accounts associated
// with the given access token.
func GetClient(accessToken string) monzo.Client {
	return monzo.Client{
		BaseURL:     "https://api.monzo.com",
		AccessToken: accessToken,
	}
}

func ProcessDeposits(cl monzo.Client, deposits []*monzo.DepositRequest) error {
	for _, deposit := range deposits {
		_, err := cl.Deposit(deposit)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetActivePots(cl monzo.Client) ([]*monzo.Pot, error) {
	allPots, err := cl.Pots()
	if err != nil {
		return nil, err
	}
	var activePots []*monzo.Pot
	for _, pot := range allPots {
		if !pot.Deleted {
			activePots = append(activePots, pot)
		}
	}
	if len(activePots) == 0 {
		return nil, errors.New("No active pots found")
	}
	return activePots, nil
}

func GetAccountWithID(cl monzo.Client, accountID string) (*monzo.Account, error) {

	accounts, err := cl.Accounts("uk_retail")
	if err != nil {
		return nil, err
	}

	var sourceAccount *monzo.Account
	for _, account := range accounts {
		if account.ID == accountID {
			sourceAccount = account
			break
		}
	}
	if sourceAccount == nil {
		return nil, errors.New("Account associated with trigger transaction not found")
	}
	return sourceAccount, nil
}

// FormatPenceToGbp takes an integer number of pence and returns a string
// representing that value in GBP.
func FormatPenceToGbp(pennies int64) string {
	return fmt.Sprintf("£%v", float64(pennies)/100)
}
