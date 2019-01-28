def parse_commands_file(cmd_file):

    unnamed_data_strcuture = {}

    with open(cmd_file, "r") as lines: 
        for line in lines:
            cmd = line.split(" ")[1]
            cmd_list = cmd.split(",")
            user_id = cmd_list[1]

            if cmd_list[0] == "DUMPLOG":    user_id = "adminxxx"
            
            if user_id in unnamed_data_strcuture: 
                unnamed_data_strcuture[user_id].append((" ".join(cmd_list)))
            else:
                unnamed_data_strcuture[user_id] = [" ".join(cmd_list)]

    return unnamed_data_strcuture
