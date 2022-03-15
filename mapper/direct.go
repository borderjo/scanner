package mapper

import (
	"database/sql"
	"fmt"
	"log"

	//	"github.com/ethereum/go-ethereum/accounts/abi"
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

func TestABI() {

	// myContractAbi := "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rune\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"TransferAllowance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"TransferOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"struct THORChain_Router.Coin[]\",\"name\":\"coins\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"VaultTransfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RUNE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address payable\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address payable\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"depositWithExpiry\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address payable\",\"name\":\"asgard\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"struct THORChain_Router.Coin[]\",\"name\":\"coins\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"returnVaultAssets\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"transferAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address payable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"transferOut\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"vaultAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
	// //example of transaction input data
	// txInput := "0xa5142faa00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000003"

	// // load contract ABI
	// abi, err := abi.JSON(strings.NewReader(myContractAbi))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(abi)
	// fmt.Println(txInput)
	// method, err := abi.MethodById(decodedSig)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(method)

	// // decode txInput Payload
	// decodedData, err := hex.DecodeString(txInput[10:])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(decodedData)
}
