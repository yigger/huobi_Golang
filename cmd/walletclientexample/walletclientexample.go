package walletclientexample

import (
	"github.com/yigger/huobi_golang/config"
	"github.com/yigger/huobi_golang/logging/applogger"
	"github.com/yigger/huobi_golang/pkg/client"
	"github.com/yigger/huobi_golang/pkg/getrequest"
	"github.com/yigger/huobi_golang/pkg/postrequest"
)

func RunAllExamples() {
	getDepositAddress()
	getSubUserDepositAddress()
	getWithdrawQuota()
	createWithdraw()
	cancelWithdraw()
	queryDepositWithdraw()
	querySubUserDepositHistory()
}

func getDepositAddress() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "usdt"
	resp, err := client.GetDepositAddress(currency)
	if err != nil {
		applogger.Error("Get deposit address error: %s", err)
	} else {
		applogger.Info("Get deposit address, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("DepositAddress: %+v", result)
		}
	}
}

func getSubUserDepositAddress() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "usdt"
	resp, err := client.GetSubUserDepositAddress(config.SubUid, currency)
	if err != nil {
		applogger.Error("Get sub user deposit address error: %s", err)
	} else {
		applogger.Info("Get sub user deposit address, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("DepositAddress: %+v", result)
		}
	}
}

func getWithdrawQuota() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "usdt"
	resp, err := client.GetWithdrawQuota(currency)
	if err != nil {
		applogger.Error("Get withdraw quota error: %s", err)
	} else {
		applogger.Info("Currency: %s, Chain: %s, MaxWithdrawAmt: %s", resp.Currency, resp.Chains[0].Chain, resp.Chains[0].MaxWithdrawAmt)
	}
}

func createWithdraw() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	createWithdrawRequest := postrequest.CreateWithdrawRequest{
		Address:  "xxxx",
		Amount:   "1.0",
		Currency: "usdt",
		Fee:      "1.0"}
	resp, err := client.CreateWithdraw(createWithdrawRequest)
	if err != nil {
		applogger.Error("Create withdraw request error: %s", err)
	} else {
		applogger.Info("Create withdraw request successfully: id=%d", resp)
	}
}

func cancelWithdraw() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.CancelWithdraw(12345)
	if err != nil {
		applogger.Error("Cancel withdraw error: %s", err)
	} else {
		applogger.Info("Cancel withdraw successfully: id=%d", resp)
	}
}

func queryDepositWithdraw() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	depositType := "deposit"
	queryDepositWithdrawOptionalRequest := getrequest.QueryDepositWithdrawOptionalRequest{Currency: "usdt"}
	resp, err := client.QueryDepositWithdraw(depositType, queryDepositWithdrawOptionalRequest)
	if err != nil {
		applogger.Error("Query deposit-withdraw history error: %s", err)
	} else {
		applogger.Info("Query deposit-withdraw history, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("resp: %+v", result)
		}
	}
}

func querySubUserDepositHistory() {
	client := new(client.WalletClient).Init(config.AccessKey, config.SecretKey, config.Host)
	optionalRequest := getrequest.QuerySubUserDepositHistoryOptionalRequest{Currency: "usdt"}
	resp, err := client.QuerySubUserDepositHistory(config.SubUid, optionalRequest)
	if err != nil {
		applogger.Error("Query sub user deposit history error: %s", err)
	} else {
		applogger.Info("Query sub user deposit history, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("resp: %+v", result)
		}
	}
}
