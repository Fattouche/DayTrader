package main

func (user *User) popFromBuyStack() *Buy {
	if len(user.BuyStack) == 0 {
		return nil
	}
	buy := user.BuyStack[len(user.BuyStack)-1]
	user.BuyStack = user.BuyStack[:len(user.BuyStack)-1]
	setCache(user.Id, user)
	return buy
}

func (user *User) popFromSellStack() *Sell {
	if len(user.SellStack) == 0 {
		return nil
	}
	sell := user.SellStack[len(user.SellStack)-1]
	user.SellStack = user.SellStack[:len(user.SellStack)-1]
	setCache(user.Id, user)
	return sell
}

func getUser(userID string) *User {
	user, err := getCacheUser(userID)
	if err != nil {
		db.QueryRow("SELECT Balance from User where Id=?", user.Id).Scan(&user.Balance)
		user.SellStack = make([]*Sell, 0)
		user.BuyStack = make([]*Buy, 0)
		setCache(userID, user)
	}
	return user
}

func createUser(userID string) error {
	_, err := db.Exec("insert ignore into User(Id) values(?)", userID)
	return err
}

func (user *User) updateUserBalance(amount float32) (*User, error) {
	user.Balance += amount
	_, err := db.Exec("update User set Balance=? where Id=?", user.Balance, user.Id)
	if err != nil {
		return user, err
	}
	setCache(user.Id, user)
	return user, nil
}
