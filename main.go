package monzotools

import (
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
