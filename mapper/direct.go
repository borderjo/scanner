package mapper

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/onrik/ethrpc"
)

func ParseEther() {
	//client := ethrpc.New("https://mainnet.infura.io/v3/4e1fd61ab45d4cb6becd13e6550900d7")
	client := ethrpc.New("http://192.168.1.53:8546")

	bn, err := client.EthBlockNumber()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bn)
	db, err := sql.Open("mysql", "mapper:cryptomap@tcp(canter-lx:3306)/crypto")
	if err != nil {
		panic(err)
	}
	res := db.QueryRow("call getLastBlock(?)", "Ether")
	fmt.Println(res)
	var num int
	res.Scan(&num)
	//num = 46147
	//bn = 10000000
	for i := num; i <= bn; i++ {
		b, err := client.EthGetBlockByNumber(i, true)
		if err == nil {
			if len(b.Transactions) > 0 {
				//fmt.Println(b)
				for k := 0; k < len(b.Transactions); k++ {
					t := b.Transactions[k]
					saveRow(db, t.From, t.To, i, "Ether")
				}
			}
		}
	}
}

func saveRow(db *sql.DB, sender string, target string, blockNum int, chain string) {
	_, err := db.Exec("call addRow(?,?,?,?)", sender, target, blockNum, chain)
	if err != nil {
		fmt.Println(err)
	}
}
