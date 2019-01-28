def parse_commands_file(cmd_file):

    cmd_struct = {}

    with open(cmd_file, "r") as lines: 
        for line in lines:
            cmd = line.split(" ")[1]
            cmd_list = cmd.split(",")
            user_id = "adminxxx" if cmd_list[0] == "DUMPLOG" and len(cmd_list) == 2 else cmd_list[1]
            
            if user_id in cmd_struct: 
                cmd_struct[user_id].append((" ".join(cmd_list)))
            else:
                cmd_struct[user_id] = [" ".join(cmd_list)]

    return cmd_struct
