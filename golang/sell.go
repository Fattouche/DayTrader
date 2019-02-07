package main

stock = Stock.quote(stock_symbol, user.user_id)
	sell = cls(user=user, stock_symbol=stock_symbol)
	err = sell.update_cash_amount(cash_amount)
	if(err):
		return None, err
	err = sell.update_price(stock.price)
	if(err):
		return None, err
	return sell, None

func createSell(price, intendedCashAmount, actualCashAmount float32, stockSoldAmount int, symbol, userID string) (*Buy, error) {
	sell := &Sell{price: price, intended_cash_amount: intendedCashAmount, actual_cash_amount: actualCashAmount, stock_sold_amount: stockSoldAmount, stock_symbol: symbol, user_id: userID}
	err := sell.updateCashAmount(intendedCashAmount)
	if err != nil {
		return nil, err
	}
	sell.updatePrice(price)
	return sell, err
}

func (sell Sell) updateCashAmount(amount float32) error {
	user, _ := getUser(buy.user_id)
	if amount > user.Balance {
		msg := fmt.Sprintf("Not enough balance, have %f need %f", user.Balance, amount)
		return errors.New(msg)
	}
	updatedAmount := buy.intended_cash_amount - amount
	user.updateUserBalance(updatedAmount)
	buy.intended_cash_amount = float32(math.Abs(float64(updatedAmount)))
	return nil
}

func (sell Sell) updatePrice(stockPrice float32) {
	buy.price = stockPrice
	buy.stock_bought_amount = int(math.Floor(float64(buy.intended_cash_amount / buy.price)))
	buy.actual_cash_amount = float32(buy.stock_bought_amount) * buy.price
}

func (sell *Buy) commit() (*UserStock, error) {
	_, err := buy.insertBuy()
	user_stock, err := getOrCreateUserStock(buy.user_id, buy.stock_symbol)
	user_stock.updateStockAmount(buy.stock_bought_amount)
	return user_stock, err
}

func (sell *Sell) cancel() {
	user, _ := getUser(buy.user_id)
	user.updateUserBalance(buy.intended_cash_amount)
}





@classmethod
def create(cls, stock_symbol, cash_amount, user):
	stock = Stock.quote(stock_symbol, user.user_id)
	sell = cls(user=user, stock_symbol=stock_symbol)
	err = sell.update_cash_amount(cash_amount)
	if(err):
		return None, err
	err = sell.update_price(stock.price)
	if(err):
		return None, err
	return sell, None

def update_price(self, stock_price):
	self.cancel()
	user_stock, created = UserStock.objects.get_or_create(
		user=self.user, stock_symbol=self.stock_symbol)
	self.stock_sold_amount = min(
		self.intended_cash_amount//stock_price,
		user_stock.amount)
	if(self.stock_sold_amount <= 0):
		return "Update trigger price failed"
	self.actual_cash_amount = self.stock_sold_amount*stock_price
	self.stock_sold_amount*stock_price
	self.timestamp = time.time()
	self.sell_price = stock_price
	user_stock.update_amount(self.stock_sold_amount*-1)

def update_cash_amount(self, amount):
	stock = Stock.quote(self.stock_symbol, self.user.user_id)
	user_stock, created = UserStock.objects.get_or_create(
		user=self.user, stock_symbol=self.stock_symbol)
	stock_sold_amount = amount//stock.price
	if(stock_sold_amount > user_stock.amount):
		return "Not enough stock, have {0} need {1}".format(user_stock.amount, stock_sold_amount)
	self.intended_cash_amount = amount

def commit(self, user):
	user.update_balance(self.actual_cash_amount)
	self.save()

def cancel(self):
	user_stock, created = UserStock.objects.get_or_create(
		user=self.user, stock_symbol=self.stock_symbol)
	user_stock.update_amount(self.stock_sold_amount)