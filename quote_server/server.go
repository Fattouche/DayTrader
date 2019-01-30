package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

func main() {
	hostName := "localhost"
	port := ":4442"
	startServer(hostName, port)
}

func startServer(hostName, port string) {
	ln, err := net.Listen("tcp", hostName+port)
	if err != nil {
		log.Println("Error listening", err)
		return
	}
	log.Println("Listening on", hostName+port)
	buf := make([]byte, 100)
	conn, err := ln.Accept()
	if err != nil {
		log.Println("Error accepting", err)
		return
	}
	for {
		len, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Println("Error reading", err)
			return
		}
		info := string(buf[:len])
		infoArr := strings.Split(info, ",")
		msg := genResponse(infoArr[0], infoArr[1])
		conn.Write(msg)
	}
}

func genResponse(symbol, userID string) []byte {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	price := r1.Float64() * 30
	hash := "lod23EP0lofFCkEd0ilcUpjL0MuBcIh3HiwAq9QSXdU="
	timestamp := time.Now()
	userID = strings.TrimSuffix(userID, "\r")
	userID = strings.Replace(userID, " ", "", -1)
	resp := fmt.Sprintf("%f,%s,%s,%d,%s", price, symbol, userID, timestamp.Unix(), hash)
	return []byte(resp)
}
