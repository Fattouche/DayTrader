package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Fattouche/DayTrader/golang/protobuff"
)

func (user *User) toString() string {
	if user == nil {
		return ""
	}
	bytes, _ := user.MarshalJSON()
	return string(bytes)
}

func (user *User) popFromBuyStack() *Buy {
	if len(user.BuyStack) == 0 {
		return nil
	}
	buy := user.BuyStack[len(user.BuyStack)-1]
	if buy.isExpired() {
		user.BuyStack = make([]*Buy, 0)
		buy = nil
	} else {
		user.BuyStack = user.BuyStack[:len(user.BuyStack)-1]
	}
	user.setCache()
	return buy
}

func (user *User) popFromSellStack() *Sell {
	if len(user.SellStack) == 0 {
		return nil
	}
	sell := user.SellStack[len(user.SellStack)-1]
	if sell.isExpired() {
		user.SellStack = make([]*Sell, 0)
		sell = nil
	} else {
		user.SellStack = user.SellStack[:len(user.SellStack)-1]
	}
	user.setCache()
	return sell
}

func userExists(userID string, password string) (float32, map[string]int32, error) {
	var stocks map[string]int32
	var balance float32
	var err error
	if password != "" {
		err = db.QueryRow("SELECT Balance from User where Id=? and Password=?", userID, password).Scan(&balance)
	} else {
		err = db.QueryRow("SELECT Balance from User where Id=?", userID).Scan(&balance)
	}

	if err == nil {
		stocks = getAllUserStock(userID)
	}
	return balance, stocks, err
}

func getUser(userID string) *User {
	user, err := getCacheUser(userID)
	if err != nil {
		db.QueryRow("SELECT Balance from User where Id=?", user.Id).Scan(&user.Balance)
		user.SellStack = make([]*Sell, 0)
		user.BuyStack = make([]*Buy, 0)
		user.StockMap = make(map[string]int32)
		user.setCache()
	}
	return user
}

func createUser(userID string, password string) error {
	user := &User{Id: userID, Balance: 0, Password: password}
	user.SellStack = make([]*Sell, 0)
	user.BuyStack = make([]*Buy, 0)
	user.StockMap = make(map[string]int32)
	user.setCache()
	_, err := db.Exec("insert into User(Id,Password) values(?,?)", userID, password)
	return err
}

func (user *User) updateUserBalance(ctx context.Context, amount float32, writeThrough bool) (*User, error) {
	var accountAction string
	if amount < 0 {
		accountAction = "Remove"
	} else {
		accountAction = "Add"
	}
	user.Balance += amount
	//deal with max balances and overflow
	if user.Balance > MAX_BALANCE {
		user.Balance = MAX_BALANCE
	}
	//Shouldn't happen
	if user.Balance < 0 {
		user.Balance = 0
	}
	if writeThrough {
		_, err := db.Exec("update User set Balance=? where Id=?", user.Balance, user.Id)
		if err != nil {
			return user, err
		}
	}
	pbLog, err := makeLogFromContext(ctx)
	if err != nil {
		log.Println(err)
	}
	pbLog.Username = user.Id
	pbLog.AccountAction = accountAction
	pbLog.Funds = user.Balance
	logEvent := &logObj{log: &pbLog, funcName: "LogAccountTransaction"}
	logChan <- logEvent
	user.setCache()
	return user, nil
}

func (user *User) updateStockBalance(ctx context.Context, symbol string) error {
	amount := user.StockMap[symbol]
	_, err := db.Exec("insert into User_Stock(UserId,StockSymbol,Amount) Values(?,?,?) on duplicate key update Amount=?", user.Id, symbol, amount, amount)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) insertAdd(ctx context.Context, amount float32) error {
	_, err := db.Exec("insert into Add_Transaction(UserId,Amount) Values(?,?)", user.Id, amount)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) getSellTriggers() []*pb.Trigger {
	rows, err := db.Query("SELECT Sell.Id, Sell.StockSymbol, Sell.StockSoldAmount, Sell.Price from Sell inner join Sell_Trigger on Sell_Trigger.SellId=Sell.Id where Sell_Trigger.Active=true and Sell.UserId=?", user.Id)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	triggers := make([]*pb.Trigger, 0)
	for rows.Next() {
		trigger := &pb.Trigger{Buy: false}
		err = rows.Scan(nil, &trigger.Symbol, &trigger.Amount, &trigger.Price)
		if err != nil {
			log.Println("Error scanning sell trigger: ", err)
		}
		triggers = append(triggers, trigger)
	}
	return triggers
}

func (user *User) getBuyTriggers() []*pb.Trigger {
	rows, err := db.Query("SELECT Buy.Id, Buy.StockSymbol, Buy.StockBoughtAmount, Buy.Price from Buy inner join Buy_Trigger on Buy_Trigger.BuyId=Buy.Id where Buy_Trigger.Active=true and Buy.UserId=?", user.Id)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	triggers := make([]*pb.Trigger, 0)
	for rows.Next() {
		trigger := &pb.Trigger{Buy: true}
		err = rows.Scan(nil, &trigger.Symbol, &trigger.Amount, &trigger.Price)
		if err != nil {
			log.Println("Error scanning buy trigger: ", err)
		}
		triggers = append(triggers, trigger)
	}
	return triggers
}

func (user *User) getBuys() []*Buy {
	rows, err := db.Query("SELECT StockSymbol, StockBoughtAmount, ActualCashAmount, Timestamp from Buy where UserId=?", user.Id)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	buys := make([]*Buy, 0)
	for rows.Next() {
		buy := &Buy{}
		err = rows.Scan(&buy.StockSymbol, &buy.StockBoughtAmount, &buy.ActualCashAmount, &buy.Timestamp)
		if err != nil {
			log.Println("Error scanning buy: ", err)
		}
		buys = append(buys, buy)
	}
	return buys
}

func (user *User) getSells() []*Sell {
	rows, err := db.Query("SELECT StockSymbol, StockSoldAmount, ActualCashAmount, TimeStamp from Sell where UserId=?", user.Id)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	sells := make([]*Sell, 0)
	for rows.Next() {
		sell := &Sell{}
		err = rows.Scan(&sell.StockSymbol, &sell.StockSoldAmount, &sell.ActualCashAmount, &sell.Timestamp)
		if err != nil {
			log.Println("Error scanning sell: ", err)
		}
		sells = append(sells, sell)
	}
	return sells
}

func (user *User) getAdds() []*Add {
	rows, err := db.Query("SELECT Amount, TimeStamp from Add_Transaction where UserId=?", user.Id)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	adds := make([]*Add, 0)
	for rows.Next() {
		add := &Add{}
		err = rows.Scan(&add.Amount, &add.Timestamp)
		if err != nil {
			log.Println("Error scanning add: ", err)
		}
		adds = append(adds, add)
	}
	return adds
}

func (user *User) getTransactions() []*pb.Transaction {
	buys := user.getBuys()
	sells := user.getSells()
	adds := user.getAdds()
	transactions := make([]*pb.Transaction, 0)
	timeFormat := "2 Jan 2006 15:04 PDT"
	local, _ := time.LoadLocation("America/Los_Angeles")
	for _, buy := range buys {
		transaction := &pb.Transaction{CommandName: "Buy", StockAmount: int32(buy.StockBoughtAmount), BalanceChange: buy.ActualCashAmount, StockSymbol: buy.StockSymbol, Timestamp: buy.Timestamp.In(local).Format(timeFormat)}
		transactions = append(transactions, transaction)
	}
	for _, sell := range sells {
		transaction := &pb.Transaction{CommandName: "Sell", StockAmount: int32(sell.StockSoldAmount), BalanceChange: sell.ActualCashAmount, StockSymbol: sell.StockSymbol, Timestamp: sell.Timestamp.In(local).Format(timeFormat)}
		transactions = append(transactions, transaction)
	}
	for _, add := range adds {
		transaction := &pb.Transaction{CommandName: "Add", BalanceChange: add.Amount, Timestamp: add.Timestamp.In(local).Format(timeFormat)}
		transactions = append(transactions, transaction)
	}
	return transactions
}
