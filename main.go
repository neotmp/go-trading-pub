package main

import (
	"fmt"

	"github.com/neotmp/go-trading/broker"
	"github.com/neotmp/go-trading/database"
)

func main() {

	database.Connect()

	b := broker.Broker{}
	brs, err := b.ListBrokers()
	if err != nil {
		fmt.Println(err, "Problem with listing brokers")
	}
	fmt.Println(brs, "Brokers")

	// o := broker.Order{
	// 	Type:       1,
	// 	Volume:     0.03,
	// 	Pair:       "EURGBP",
	// 	Price:      0.84459,
	// 	Status:     1,
	// 	AccountId:  1,
	// 	BrokerId:   1,
	// 	Margin:     0,
	// 	PairId:     19,
	// 	Commission: 0,
	// 	SpreadPips: 2.5,
	// 	Timestamp:  time.Now(),
	// }

	roboForex := brs[0]

	// for _, v := range roboForex.Orders {
	// 	fmt.Println(v.Margin, "Roboforex Orders Margin")
	// }

	acc, err := roboForex.CalculateAccountEquity(roboForex.Accounts[0])
	if err != nil {
		fmt.Println(err, "Problem w/ Equity")
	}

	// acc, err := roboForex.FindAccount(roboForex.Accounts, roboForex.Accounts[0].Id)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println(acc.Accounts[0].Equity, "accs Equity")

	//fmt.Println(acc.Accounts[0].Equity, "Equity main")

	// fm, err := roboForex.CalculateFreeMargin(roboForex.Accounts[0])
	// if err != nil {
	// 	fmt.Println(err, "Problem w/ Free Margin")
	// }

	// ml, err := roboForex.CalculateMarginLevel(roboForex.Accounts[0])
	// if err != nil {
	// 	fmt.Println(err, "Problem w/ Free Margin")
	// }

	// fmt.Println(acc.Accounts[0].Margin, "Margin")
	// fmt.Println(acc.Accounts[0].Equity, "Equity")
	// fmt.Println(acc.Accounts[0].Balance, "Balance")
	// fmt.Println(fm.Accounts[0].FreeMargin, "Free Margin")
	// fmt.Println(ml.Accounts[0].MarginLevel, "Margin Level")

	// orders, err := roboForex.PlaceOrder(&o)
	// if err != nil {
	// 	fmt.Println(err, "Problem w/ Order")
	// }

	// mm, err := roboForex.CalcAccountMargin(roboForex.Accounts[0])
	// if err != nil {
	// 	fmt.Println(err, "Problem w/ calcAccountMargin")
	// }

	// fmt.Println(mm.Accounts[0].Margin, "Acc Margin")

	// for _, v := range orders.Orders {
	// 	fmt.Println(v, "Orders Placed")
	// }

	// orders, err := roboForex.ListOrders()
	// if err != nil {
	// 	fmt.Println(err, "Problem w/ Orders")
	// }

	//fmt.Println(orders, "Orders")

	// posPfofit, err := roboForex.CalcPositionProfit(roboForex.Accounts[0])
	// if err != nil {
	// 	fmt.Println(err, "Problem w/ Account Position profit")
	// }

	// for _, v := range posPfofit.Positions {
	// 	fmt.Println(v, "Position Main")
	// }

}
