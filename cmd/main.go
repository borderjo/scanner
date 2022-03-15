package main

import (
	"fmt"

	"github.com/borderjo/scanner/chainscan"
	"github.com/borderjo/scanner/mapper"
)

var alex string = "0x0C498804369755445FD0Ba755bbbDcFf2d3Ff34E"
var john string = "0x653559D16713c471d3596E86C0bAb83C75b40580"

type Provider struct {
	Name         string
	Engine       *chainscan.Scanner
	Transactions map[int][]chainscan.Transaction
}

func main() {
	//testOne("0x315fda8566bd6666354babaf94828107ab7f0d17b430aef38e6a044475943f61")
	//chainscan.TestRpc()
	//mapper.InitMapper()
	//mapper.ParseEther()
	mapper.TestABI()
}

func testOne(txn string) {
	ether := chainscan.NewScanner("https://api.etherscan.io/api", "Z37GW68CSY3M4REFHN3IN9XZ758XAAYXBJ")
	tx := ether.GetTransactions(alex, 0)

	for _, t := range tx {
		//fmt.Println(i, t.Hash)
		if t.Hash == "0x315fda8566bd6666354babaf94828107ab7f0d17b430aef38e6a044475943f61" {
			fmt.Printf("%#v", t)
		}
	}

	ether.GetDetail(txn)
	fmt.Println(tx[0])
	//ether.GetTxByHash(txn)
}

func testAll() {
	bsc := chainscan.NewScanner("https://api.bscscan.com/api", "4KN3EGT11S1ZHXE9ZGFZXA9E4F5537FYIV")
	ether := chainscan.NewScanner("https://api.etherscan.io/api", "Z37GW68CSY3M4REFHN3IN9XZ758XAAYXBJ")
	scanners := make(map[string]Provider)
	scanners["Ether"] = Provider{Name: "ether", Engine: ether, Transactions: make(map[int][]chainscan.Transaction)}
	scanners["BSC"] = Provider{Name: "bsc", Engine: bsc, Transactions: make(map[int][]chainscan.Transaction)}

	for _, scanner := range scanners {
		for i := 0; i < 4; i++ {
			scanner.Transactions[i] = scanner.Engine.GetTransactions(alex, i)
		}
	}

	for _, v := range scanners {
		for i := 0; i < 4; i++ {
			fmt.Println(v.Name, i, len(v.Transactions[i]))
		}

	}
}
