import proto from '../protobuff/daytrader_grpc_pb';

const BACKEND_ADDRESS = 'http://localhost:80'

export function checkCredentials(username, password, callback) {
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.Command()
    command.setUserId(username)
    command.setPassword(password)

    client.getUser(command, {}, callback)
}

export function createUser(username, password, callback) {
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.Command()
    command.setUserId(username)
    command.setPassword(password)

    client.createUser(command, {}, callback)
}

export function displaySummary(username, callback) {
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.Command()
    command.setUserId(username)
    client.displaySummary(command, {}, callback)
}

export function getQuote(username, symbol, callback){
    var client = new proto.DayTraderClient(BACKEND_ADDRESS)
    var command = new proto.Command()
    command.setUserId(username)
    command.setSymbol(symbol)
    console.log("SYMBOL IS ")
    console.log(symbol)

    client.quote(command, {}, callback)
}
