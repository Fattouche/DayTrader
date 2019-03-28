package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "github.com/Fattouche/DayTrader/golang/protobuff"
	"google.golang.org/grpc"
)

var symbolCommands = map[string]int{
	"QUOTE":            1,
	"BUY":              1,
	"SELL":             1,
	"SET_BUY_AMOUNT":   1,
	"CANCEL_SET_BUY":   1,
	"SET_BUY_TRIGGER":  1,
	"SET_SELL_AMOUNT":  1,
	"SET_SELL_TRIGGER": 1,
	"CANCEL_SET_SELL":  1,
}
var amountCommands = map[string]int{
	"ADD":              1,
	"BUY":              1,
	"SELL":             1,
	"SET_BUY_AMOUNT":   1,
	"SET_BUY_TRIGGER":  1,
	"SET_SELL_AMOUNT":  1,
	"SET_SELL_TRIGGER": 1,
}

var baseURL string
var userMap = make(map[string][]*pb.Command)
var dumpData *pb.Command
var wg sync.WaitGroup
var transactionNum = int32(1)

func completeCall(command *pb.Command, client pb.DayTraderClient) {
	ctx := context.Background()
	var err error
	var resp string
	var userResp *pb.UserResponse
	var balanceResp *pb.BalanceResponse
	var priceResp *pb.PriceResponse
	var stockResp *pb.StockUpdateResponse
	var response *pb.Response
	var summaryResponse *pb.SummaryResponse

	switch command.Name {
	case "ADD":
		balanceResp, err = client.Add(ctx, command)
		resp = balanceResp.String()
	case "QUOTE":
		priceResp, err = client.Quote(ctx, command)
		resp = priceResp.String()
	case "BUY":
		balanceResp, err = client.Buy(ctx, command)
		resp = balanceResp.String()
	case "SELL":
		stockResp, err = client.Sell(ctx, command)
		resp = stockResp.String()
	case "COMMIT_BUY":
		stockResp, err = client.CommitBuy(ctx, command)
		resp = stockResp.String()
	case "COMMIT_SELL":
		userResp, err = client.CommitSell(ctx, command)
		resp = userResp.String()
	case "CANCEL_BUY":
		balanceResp, err = client.CancelBuy(ctx, command)
		resp = balanceResp.String()
	case "CANCEL_SELL":
		stockResp, err = client.CancelSell(ctx, command)
		resp = stockResp.String()
	case "SET_BUY_AMOUNT":
		balanceResp, err = client.SetBuyAmount(ctx, command)
		resp = balanceResp.String()
	case "SET_SELL_AMOUNT":
		response, err = client.SetSellAmount(ctx, command)
		resp = response.String()
	case "SET_BUY_TRIGGER":
		response, err = client.SetBuyTrigger(ctx, command)
		resp = response.String()
	case "SET_SELL_TRIGGER":
		stockResp, err = client.SetSellTrigger(ctx, command)
		resp = stockResp.String()
	case "CANCEL_SET_BUY":
		balanceResp, err = client.CancelSetBuy(ctx, command)
		resp = balanceResp.String()
	case "CANCEL_SET_SELL":
		stockResp, err = client.CancelSetSell(ctx, command)
		resp = stockResp.String()
	case "DUMPLOG":
		response, err = client.DumpLog(ctx, command)
		resp = response.String()
	case "DISPLAY_SUMMARY":
		summaryResponse, err = client.DisplaySummary(ctx, command)
		resp = summaryResponse.String()
	}
	_, _ = resp, err
	// UNCOMMENT FOR RESPONSES
	// fmt.Print(command.Name + " ")
	// log.Print("RESP: " + resp)
	// if err != nil {
	// 	fmt.Println(" ERROR: " + err.Error())
	// }
}

func parseCommands(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var userID string
	for scanner.Scan() {
		totalCommand := strings.Split(scanner.Text(), " ")
		userCommands := strings.Split(totalCommand[1], ",")
		if userCommands[0] == "DUMPLOG" && len(userCommands) == 2 {
			dumpData = &pb.Command{Filename: userCommands[1], Name: "DUMPLOG"}
			continue
		} else {
			userID = userCommands[1]
		}
		command := generateRequest(userID, userCommands, transactionNum)
		transactionNum++
		if _, ok := userMap[userID]; !ok {
			userMap[userID] = make([]*pb.Command, 0)
		}
		userMap[userID] = append(userMap[userID], command)
	}
}

func generateRequest(userID string, commands []string, transactionNum int32) *pb.Command {
	command := &pb.Command{}
	if len(userID) > 0 {
		command.UserId = userID
	}
	if commands[0] == "DUMPLOG" {
		if len(commands) > 2 {
			command.Filename = commands[2]
		} else {
			command.Filename = commands[1]
		}
	}
	if _, ok := symbolCommands[commands[0]]; ok {
		if len(commands) > 2 {
			command.Symbol = commands[2]
		}
	}
	if _, ok := amountCommands[commands[0]]; ok {
		if commands[0] == "ADD" {
			if len(commands) > 2 {
				amount, _ := strconv.ParseFloat(commands[2], 32)
				command.Amount = float32(amount)
			}
		} else {
			if len(commands) > 3 {
				amount, _ := strconv.ParseFloat(commands[3], 32)
				command.Amount = float32(amount)
			}
		}
	}
	command.TransactionId = transactionNum
	command.Name = commands[0]
	return command
}

func makeRequest(commands []*pb.Command) {
	conn, err := grpc.Dial(baseURL, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial to %s with %v", baseURL, err)
	}
	defer conn.Close()
	client := pb.NewDayTraderClient(conn)
	for _, command := range commands {
		completeCall(command, client)
	}
	wg.Done()
}

func makeDumpRequest() {
	conn, err := grpc.Dial(baseURL, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial to %s with %v", baseURL, err)
	}
	dumpData.TransactionId = transactionNum
	client := pb.NewDayTraderClient(conn)
	completeCall(dumpData, client)
}

func createUsers(userMap map[string][]*pb.Command) {
	conn, err := grpc.Dial(baseURL, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial to %s with %v", baseURL, err)
	}
	defer conn.Close()
	client := pb.NewDayTraderClient(conn)
	for userId := range userMap {
		cmd := &pb.Command{UserId: userId, Name: "CREATE_USER"}
		client.CreateUser(context.Background(), cmd)
	}
}

func main() {
	fileName := flag.String("f", "./workload_files/1_user_workload", "The name of the workload file")
	tempBaseURL := flag.String("url", "daytrader_lb:80", "The url of the web server")
	flag.Parse()
	baseURL = *tempBaseURL
	parseCommands(*fileName)
	createUsers(userMap)
	start := time.Now()
	wg.Add(len(userMap))
	for _, requests := range userMap {
		go makeRequest(requests)
	}
	wg.Wait()
	makeDumpRequest()
	fmt.Println("Time taken: ", time.Since(start))
}
