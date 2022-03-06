package chainscan

import (
	"fmt"
	"log"

	"github.com/onrik/ethrpc"
)

func TestRpc() {
	client := ethrpc.New("https://mainnet.infura.io/v3/4e1fd61ab45d4cb6becd13e6550900d7")
	version, err := client.Web3ClientVersion()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(version)

	accts, err := client.EthAccounts()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(accts)

	// b, err := client.EthGetBlockByNumber(12345, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(b)

	bn, err := client.EthBlockNumber()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bn)

	// for i := 46147; i <= bn; i++ {
	// 	b, err := client.EthGetBlockByNumber(i, true)
	// 	if err == nil {
	// 		if len(b.Transactions) > 0 {
	// 			fmt.Println(b)
	// 			for k := 0; k < len(b.Transactions); k++ {
	// 				t := b.Transactions[k]

	// 			}
	// 		}
	// 	}
	//}

	// bal, err := client.EthGetBalance("0x0C498804369755445FD0Ba755bbbDcFf2d3Ff34E", "12345")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(bal)

}
