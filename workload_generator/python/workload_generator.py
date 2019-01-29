import threading
import parser
import worker
import time


def parse_commands_file(cmd_file):

    cmd_struct = {}

    with open(cmd_file, "r") as lines:
        for line in lines:
            cmd = line.split(" ")[1]
            cmd_list = cmd.split(",")
            user_id = "adminxxx" if cmd_list[0] == "DUMPLOG" and len(
                cmd_list) == 2 else cmd_list[1]

            if user_id in cmd_struct:
                cmd_struct[user_id].append((" ".join(cmd_list)))
            else:
                cmd_struct[user_id] = [" ".join(cmd_list)]

    return cmd_struct


class WorkloadGenerator:

    def __init__(self, user_cmd_file):
        self.user_cmd_file = user_cmd_file

    def update_file(user_cmd_file):
        self.user_cmd_file = user_cmd_file

    def run(self):
        user_cmds = parse_commands_file(self.user_cmd_file)
        all_the_threads = {}
        start = time.time()
        for user_id in user_cmds:
            all_the_threads[user_id] = threading.Thread(
                target=worker.target, name=user_id, args=(user_id, user_cmds[user_id]))
            all_the_threads[user_id].start()

        for thread in all_the_threads:
            all_the_threads[thread].join()
        end = time.time()
        print("Time taken: ", end-start)
        return


def main():
    generator = WorkloadGenerator("2userWorkLoad.txt")
    generator.run()


if __name__ == "__main__":
    main()
