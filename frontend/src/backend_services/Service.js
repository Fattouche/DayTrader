import proto from '../protobuff/daytrader_grpc_pb';

const BACKEND_ADDRESS = 'http://localhost:80'

export function checkCredentials(username, password, callback) {
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setPassword(password)

    client.getUser(command, {} /* metadata */, callback)
}

export function createUser(username, password, callback) {
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setPassword(password)
    command.setName("CREATE_USER")

    client.createUser(command, {} /* metadata */, callback)
}

export function dumplog(username, filename, callback) {
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setFilename(filename)
    command.setName("DUMPLOG")
    client.dumpLog(command, {} /* metadata */, callback)
}

export function displaySummary(username, callback) {
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setName("DISPLAY_SUMMARY")
    client.displaySummary(command, {} /* metadata */, callback)
}

export function getQuote(username, symbol, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setSymbol(symbol)
    command.setName("QUOTE")
    client.quote(command, {} /* metadata */, callback)
}

export function buy(username, symbol, amount, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setSymbol(symbol)
    command.setAmount(amount)
    command.setName("BUY")
    client.buy(command, {} /* metadata */, callback)
}

export function setBuyAmount(username, symbol, amount, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setSymbol(symbol)
    command.setAmount(amount)
    command.setName("SET_BUY_AMOUNT")
    client.setBuyAmount(command, {} /* metadata */, callback)
}

export function setBuyTrigger(username, symbol, price, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setSymbol(symbol)
    command.setAmount(price)
    command.setName("SET_BUY_TRIGGER")
    client.setBuyTrigger(command, {} /* metadata */, callback)
}

export function commitBuy(username, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setName("COMMIT_BUY")
    client.commitBuy(command, {}, callback)
}
export function cancelBuy(username, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setName("CANCEL_BUY")
    client.cancelBuy(command, {} /* metadata */, callback)
}

export function sell(username, symbol, amount, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setSymbol(symbol)
    command.setAmount(amount)
    command.setName("SELL")
    client.sell(command, {} /* metadata */, callback)
}

export function setSellAmount(username, symbol, amount, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setSymbol(symbol)
    command.setAmount(amount)
    command.setName("SET_SELL_AMOUNT")
    client.setSellAmount(command, {} /* metadata */, callback)
}

export function setSellTrigger(username, symbol, price, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setSymbol(symbol)
    command.setAmount(price)
    command.setName("SET_SELL_TRIGGER")
    client.setSellTrigger(command, {} /* metadata */, callback)
}

export function commitSell(username, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setName("COMMIT_SELL")
    client.commitSell(command, {}, callback)
}

export function cancelSell(username, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setName("CANCEL_SELL")
    client.cancelSell(command, {} /* metadata */, callback)
}

export function add(username, amount, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setAmount(amount)
    command.setName("ADD")
    client.add(command, {} /* metadata */, callback)
}

export function cancelSetBuy(username, symbol, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setSymbol(symbol)
    command.setName("CANCEL_SET_BUY")
    client.cancelSetBuy(command, {} /* metadata */, callback)
}

export function cancelSetSell(username, symbol, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.command()
    command.setUserId(username)
    command.setSymbol(symbol)
    command.setName("CANCEL_SET_SELL")
    client.cancelSetSell(command, {} /* metadata */, callback)
}