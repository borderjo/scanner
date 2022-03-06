package mapper

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitMapper() {
	fmt.Println("start map")
	//	db, err := sql.Open("mysql", "root:crypto@tcp(ubuntu2004.wsl:3306)/sys")
	db, err := sql.Open("mysql", "mapper:cryptomap@tcp(ubuntu2004.wsl:3306)/crypto")

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("pong")

	defer db.Close()
}

func writeRow(source string, address string, blockIndex uint64) {

}

func createDb() {

	db, err := sql.Open("mysql", "root:crypto@tcp(ubuntu2004.wsl:3306)/sys")

	if err != nil {
		panic(err)
	}
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+"crypto")
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return
	}
	log.Printf("rows affected %d\n", no)
}
