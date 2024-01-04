package tools

import (
	"log"
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "ALEX_SK",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "JASON_SK",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "MARIE_SK",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
	},
	"marie": {
		Coins:    3000,
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// simulate DB call
	logger := log.Default()
	logger.Println("Making DB call...")
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	logger.Println("Finished DB call!")
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
