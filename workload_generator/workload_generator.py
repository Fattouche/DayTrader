import threading
import parser
import worker

class WorkloadGenerator:

    def __init__ (self, user_cmd_file):
        self.user_cmd_file = user_cmd_file

    def run(self):
        user_cmds = parser.parse_commands_file(self.user_cmd_file)
        all_the_threads = {}
        for user_id in user_cmds:
            all_the_threads[user_id] = threading.Thread(target=worker.target, name=user_id, args=(user_id, user_cmds[user_id]))
            all_the_threads[user_id].start()

        for thread in all_the_threads:
            all_the_threads[thread].join()
        return

def main():
    generator =  WorkloadGenerator("2userWorkLoad.txt")
    generator.run()

if __name__ == "__main__":
    main()