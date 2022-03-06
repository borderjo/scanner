package chainscan

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type TxDetail struct {
}

func (es *Scanner) GetDetail(txhash string) {
	url := "?module=proxy"
	url = url + "&action=eth_getTransactionByHash"
	url = url + "&txhash=" + txhash
	url = url + "&apiKey=" + es.apiKey

	resp, _ := http.Get(es.baseUrl + url)

	var bld strings.Builder
	io.Copy(&bld, resp.Body)
	fmt.Println(bld.String())

	// to seems to be a token type

	// var pkg transactionPkg
	// decoder := json.NewDecoder(resp.Body)
	// if err := decoder.Decode(&pkg); err != nil {
	// 	fmt.Println(err)
	// }

	// transactions := make([]Transaction, len(pkg.Result))
	// for i, v := range pkg.Result {
	// 	transactions[i] = *v.GetTx()
	// }
	//return transactions
}

func (es *Scanner) GetTxByHash(txhash string) {
	url := "?module=account&action=txlistinternal&txhash=" + txhash
	url = url + "&apiKey=" + es.apiKey

	resp, _ := http.Get(es.baseUrl + url)

	var bld strings.Builder
	io.Copy(&bld, resp.Body)
	fmt.Println(bld.String())

	// to seems to be a token type

	// var pkg transactionPkg
	// decoder := json.NewDecoder(resp.Body)
	// if err := decoder.Decode(&pkg); err != nil {
	// 	fmt.Println(err)
	// }

	// transactions := make([]Transaction, len(pkg.Result))
	// for i, v := range pkg.Result {
	// 	transactions[i] = *v.GetTx()
	// }
	//return transactions
}
