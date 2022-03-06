package chainscan

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Transaction struct {
	BlockNumber       uint64
	TimeStamp         time.Time
	Hash              string
	Nonce             uint64
	BlockHash         string
	Index             uint64
	From              string
	To                string
	Value             uint64
	Gas               uint64
	GasPrice          uint64
	IsError           bool
	TxReceiptStatus   uint64
	Input             string
	ContractAddress   string
	CumulativeGasUsed uint64
	GasUsed           uint64
	Confirmations     uint64
}

func (t *Transaction) String() string {
	e, _ := json.Marshal(t)
	return string(e)
}

type transactionPkg struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Result  []transactionJson `json:"result"`
}

type transactionJson struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	Index             string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxReceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
}

func (t *transactionJson) GetTx() *Transaction {
	tx := Transaction{Hash: t.Hash, BlockHash: t.BlockHash, From: t.From, To: t.To, Input: t.Input, ContractAddress: t.ContractAddress}

	if ui, err := strconv.ParseUint(t.BlockNumber, 10, 64); err == nil {
		tx.BlockNumber = ui
	}
	if ui, err := strconv.ParseUint(t.Nonce, 10, 64); err == nil {
		tx.Nonce = ui
	}
	if ui, err := strconv.ParseUint(t.Index, 10, 64); err == nil {
		tx.Index = ui
	}
	if ui, err := strconv.ParseUint(t.Value, 10, 64); err == nil {
		tx.Value = ui
	}
	if ui, err := strconv.ParseUint(t.Gas, 10, 64); err == nil {
		tx.Gas = ui
	}
	if ui, err := strconv.ParseUint(t.GasPrice, 10, 64); err == nil {
		tx.GasPrice = ui
	}
	if ui, err := strconv.ParseUint(t.TxReceiptStatus, 10, 64); err == nil {
		tx.TxReceiptStatus = ui
	}
	if ui, err := strconv.ParseUint(t.CumulativeGasUsed, 10, 64); err == nil {
		tx.CumulativeGasUsed = ui
	}
	if ui, err := strconv.ParseUint(t.GasUsed, 10, 64); err == nil {
		tx.GasUsed = ui
	}
	if ui, err := strconv.ParseUint(t.Confirmations, 10, 64); err == nil {
		tx.Confirmations = ui
	}
	tx.IsError = t.IsError != "0"

	i, err := strconv.ParseInt(t.TimeStamp, 10, 64)
	if err != nil {
		panic(err)
	}
	tx.TimeStamp = time.Unix(i, 0)
	//fmt.Println(tx.TimeStamp)
	return &tx
}

func (es *Scanner) GetTransactions(addr string, Type int) []Transaction {
	var action string
	switch Type {
	case 0:
		action = "txlist"
	case 1:
		action = "txlistinternal"
	case 2:
		action = "tokentx"
	case 3:
		action = "tokennfttx"
	}

	return es.getTxTypes(addr, action)
}

// func (es *Provider) GetAltTransactions(addr string) []Transaction {
// 	return es.getTxTypes(addr, "tokentx")
// }

func (es *Scanner) getTxTypes(addr string, action string) []Transaction {
	url := "?module=account"
	url = url + "&action=" + action
	url = url + "&address=" + addr
	url = url + "&tag=latest"
	url = url + "&startblock=0"
	url = url + "&endblock=99999999"
	url = url + "&page=1"
	url = url + "&offset=1000"
	url = url + "&sort=asc"
	url = url + "&apiKey=" + es.apiKey
	resp, _ := http.Get(es.baseUrl + url)

	var pkg transactionPkg
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&pkg); err != nil {
		fmt.Println(err)
	}

	transactions := make([]Transaction, len(pkg.Result))
	for i, v := range pkg.Result {
		transactions[i] = *v.GetTx()
	}
	return transactions
}
