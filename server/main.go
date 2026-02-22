package main

import (
	"fmt"
	"log"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/db"
	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/models"
	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/repos"
)

func main() {
	fmt.Println("Trying to Connect with DB")

	conn, err := db.OpenDB("./mmeMain.db")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("connected to mmeMain successfully")

	acc := &models.Account{
		AccID:       "0000pot",
		AccType:     "BUSINESS",
		AccName:     "potbeater",
		AccAddress:  "pala,kottayam",
		AccArchived: false,
	}

	res, err := repos.InsertIntoAccounts(conn, acc)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Rows affected:", res)
}
