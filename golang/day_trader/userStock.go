package main

type UserStock struct {
	UserId      string
	StockSymbol string
	Amount      int
}

func getOrCreateUserStock(userID, symbol string) (*UserStock, error) {
	userStock := &UserStock{UserId: userID, StockSymbol: symbol}
	err := db.QueryRow("SELECT Amount from User_Stock where UserId=? and StockSymbol=?", userID, symbol).Scan(&userStock.Amount)
	if err != nil {
		db.Exec("insert into User_Stock(UserId,StockSymbol) values(?,?)", userID, symbol)
		userStock.Amount = 0
	}
	return userStock, err
}

func (userStock *UserStock) updateStockAmount(amount int) error {
	userStock.Amount += amount
	_, err := db.Exec("update User_Stock set Amount=? where UserId=? and StockSymbol=?", userStock.Amount, userStock.UserId, userStock.StockSymbol)
	return err
}
