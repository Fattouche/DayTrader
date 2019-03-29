import proto from '../protobuff/daytrader_grpc_pb';

const BACKEND_ADDRESS = 'http://localhost:80'

export function checkCredentials(username, password, callback) {
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.Command()
    command.setUserId(username)
    command.setPassword(password)

    client.getUser(command, {} /* metadata */, callback)
}

export function createUser(username, password, callback) {
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.Command()
    command.setUserId(username)
    command.setPassword(password)

    client.createUser(command, {} /* metadata */, callback)
}

export function displaySummary(username, callback) {
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.Command()
    command.setUserId(username)
    client.displaySummary(command, {} /* metadata */, callback)
}

export function getQuote(username, symbol, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.Command()
    command.setUserId(username)
    command.setSymbol(symbol)
    client.quote(command, {} /* metadata */, callback)
}

export function buy(username, symbol, amount, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.Command()
    command.setUserId(username)
    command.setSymbol(symbol)
    command.setAmount(amount)
    client.buy(command, {} /* metadata */, callback)
}

export function commitBuy(username, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.Command()
    command.setUserId(username)
    client.commitBuy(command, {}, callback)
}
export function cancelBuy(username, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.Command()
    command.setUserId(username)
    client.cancelBuy(command, {} /* metadata */, callback)
}

export function sell(username, symbol, amount, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.Command()
    command.setUserId(username)
    command.setSymbol(symbol)
    command.setAmount(amount)
    client.sell(command, {} /* metadata */, callback)
}

export function commitSell(username, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.Command()
    command.setUserId(username)
    client.commitSell(command, {}, callback)
}

export function cancelSell(username, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.Command()
    command.setUserId(username)
    client.cancelSell(command, {} /* metadata */, callback)
}

export function add(username, amount, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.Command()
    command.setUserId(username)
    command.setAmount(amount)
    client.add(command, {} /* metadata */, callback)
}