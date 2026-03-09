package main

import (
	"fmt"
	"log"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/db"
	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/models"
	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/repos"
)

var tempAcc1 = models.Account{
	AccID:        "temp-1",
	AccType:      "CONSUMER",
	AccName:      "Temporary Account 1",
	AccContact:   "0000000000",
	AccAddress:   "Nowhere",
	AccArchived:  false,
	AccCreatedAt: "2026-03-08",
}

var tempAcc2 = models.Account{
	AccID:        "temp-2",
	AccType:      "CONSUMER",
	AccName:      "Temporary Account 2",
	AccContact:   "1111111111",
	AccAddress:   "Somewhere",
	AccArchived:  false,
	AccCreatedAt: "2026-03-08",
}

func main() {
	fmt.Println("Trying to Connect with DB")

	conn, err := db.OpenDB("./mmeMain.db")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("connected to mmeMain successfully")
	_, err = repos.InsertIntoAccounts(conn, &tempAcc1)
	if err != nil {
		fmt.Println(err)
	}
	_, err = repos.InsertIntoAccounts(conn, &tempAcc2)
	if err != nil {
		fmt.Println(err)
	}

	res, err := repos.CreateTransactionRow(conn, tempAcc1.AccID, tempAcc2.AccID, 100, "SALE", "null", nil)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("transcation successfull id", res)
	}
}
