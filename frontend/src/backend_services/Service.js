import proto from '../protobuff/daytrader_grpc_pb';

export function check_credentials(username, password, callback) {
    var client = new proto.DayTraderClient('http://localhost:80')
    var command = new proto.Command()
    command.setUserId(username)
    command.setPassword(password)

    var call = client.getUser(command, {}, callback)
    call.on('status', function(status) {
    console.log(status.code);
    console.log(status.details);
    console.log(status.metadata);
    })
}
