package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type userCommandLog struct {
	XMLName        xml.Name      `xml:"userCommand"`
	Timestamp      int64         `xml:"timestamp"`
	Server         string        `xml:"server"`
	TransactionNum int           `xml:"transactionNum"`
	Command        string        `xml:"command"`
	Username       string        `xml:"username,omitempty"`
	StockSymbol    string        `xml:"stockSymbol,omitempty"`
	Filename       string        `xml:"filename,omitempty"`
	Funds          monetaryFloat `xml:"funds,omitempty"`
}

type quoteServerLog struct {
	XMLName         xml.Name      `xml:"quoteServer"`
	Timestamp       int64         `xml:"timestamp"`
	Server          string        `xml:"server"`
	TransactionNum  int           `xml:"transactionNum"`
	Price           monetaryFloat `xml:"price"`
	StockSymbol     string        `xml:"stockSymbol"`
	Username        string        `xml:"username"`
	QuoteServerTime int64         `xml:"quoteServerTime"`
	CryptoKey       string        `xml:"cryptokey"`
}

type accountTransactionLog struct {
	XMLName        xml.Name      `xml:"accountTransaction"`
	Timestamp      int64         `xml:"timestamp"`
	Server         string        `xml:"server"`
	TransactionNum int           `xml:"transactionNum"`
	Action         string        `xml:"action"`
	Username       string        `xml:"username"`
	Funds          monetaryFloat `xml:"funds"`
}

type systemEventLog struct {
	XMLName        xml.Name      `xml:"systemEvent"`
	Timestamp      int64         `xml:"timestamp"`
	Server         string        `xml:"server"`
	TransactionNum int           `xml:"transactionNum"`
	Command        string        `xml:"command"`
	Username       string        `xml:"username,omitempty"`
	StockSymbol    string        `xml:"stockSymbol,omitempty"`
	Filename       string        `xml:"filename,omitempty"`
	Funds          monetaryFloat `xml:"funds,omitempty"`
}

type errorEventLog struct {
	XMLName        xml.Name      `xml:"errorEvent"`
	Timestamp      int64         `xml:"timestamp"`
	Server         string        `xml:"server"`
	TransactionNum int           `xml:"transactionNum"`
	Command        string        `xml:"command"`
	Username       string        `xml:"username,omitempty"`
	StockSymbol    string        `xml:"stockSymbol,omitempty"`
	Filename       string        `xml:"filename,omitempty"`
	Funds          monetaryFloat `xml:"funds,omitempty"`
	ErrorMessage   string        `xml:"errorMessage,omitempty"`
}

type debugEventLog struct {
	XMLName        xml.Name      `xml:"debugEvent"`
	Timestamp      int64         `xml:"timestamp"`
	Server         string        `xml:"server"`
	TransactionNum int           `xml:"transactionNum"`
	Command        string        `xml:"command"`
	Username       string        `xml:"username,omitempty"`
	StockSymbol    string        `xml:"stockSymbol,omitempty"`
	Filename       string        `xml:"filename,omitempty"`
	Funds          monetaryFloat `xml:"funds,omitempty"`
	DebugMessage   string        `xml:"debugMessage,omitempty"`
}

type monetaryFloat float64

func (f monetaryFloat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%.2f", f)
	return e.EncodeElement(s, start)
}

func getRows(tableName, userID string) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error
	if userID != "" {
		rows, err = db.Query("SELECT * FROM "+tableName+" WHERE Username = ?",
			userID)
	} else {
		rows, err = db.Query("SELECT * FROM " + tableName)
	}
	if err != nil {
		log.Println("Error performing query: ", err)
	}
	return rows, err
}

func dumpLogsToXML(userID string, filename string) {
	// For final dumplog, do a wait to make sure all logs are in the DB
	if userID == "" {
		time.Sleep(time.Second * 20)
	}

	dumplogsDir := "/go/src/logging/dumplogs"
	sanitizedFilename := filepath.Join(dumplogsDir, filepath.Clean(filename))
	f, err := os.Create(sanitizedFilename)
	if err != nil {
		log.Println("Error scanning trigger: ", err)
	}
	defer f.Close()

	f.WriteString("<?xml version=\"1.0\"?>\n")
	f.WriteString("<log>\n")

	var rows *sql.Rows

	// UserCommand
	rows, err = getRows("UserCommandLog", userID)
	if err != nil {
		return
	}

	for rows.Next() {
		xmlLog := &userCommandLog{}
		var timestamp time.Time
		err = rows.Scan(
			&timestamp, &xmlLog.Server, &xmlLog.TransactionNum,
			&xmlLog.Command, &xmlLog.Username, &xmlLog.StockSymbol,
			&xmlLog.Filename, &xmlLog.Funds,
		)
		if err != nil {
			log.Println("Error scanning trigger: ", err)
		}
		xmlLog.Timestamp = timestamp.UnixNano() / 1000000
		output, err := xml.MarshalIndent(xmlLog, "\t", "\t")
		if err != nil {
			log.Println("Error marshalling to XML: ", err)
		}
		f.Write(output)
		f.WriteString("\n")
	}
	rows.Close()

	// QuoteServer
	rows, err = getRows("QuoteServerLog", userID)
	if err != nil {
		return
	}

	for rows.Next() {
		xmlLog := &quoteServerLog{}
		var timestamp time.Time
		err = rows.Scan(
			&timestamp, &xmlLog.Server, &xmlLog.TransactionNum,
			&xmlLog.Price, &xmlLog.StockSymbol, &xmlLog.Username,
			&xmlLog.QuoteServerTime, &xmlLog.CryptoKey,
		)
		if err != nil {
			log.Println("Error scanning trigger: ", err)
		}
		xmlLog.Timestamp = timestamp.UnixNano() / 1000000
		output, err := xml.MarshalIndent(xmlLog, "\t", "\t")
		if err != nil {
			log.Println("Error marshalling to XML: ", err)
		}
		f.Write(output)
		f.WriteString("\n")
	}
	rows.Close()

	// AccountTransaction
	rows, err = getRows("AccountTransactionLog", userID)
	if err != nil {
		return
	}

	for rows.Next() {
		xmlLog := &accountTransactionLog{}
		var timestamp time.Time
		err = rows.Scan(
			&timestamp, &xmlLog.Server, &xmlLog.TransactionNum,
			&xmlLog.Action, &xmlLog.Username, &xmlLog.Funds,
		)
		if err != nil {
			log.Println("Error scanning trigger: ", err)
		}
		xmlLog.Timestamp = timestamp.UnixNano() / 1000000
		output, err := xml.MarshalIndent(xmlLog, "\t", "\t")
		if err != nil {
			log.Println("Error marshalling to XML: ", err)
		}
		f.Write(output)
		f.WriteString("\n")
	}
	rows.Close()

	// SystemEvent
	rows, err = getRows("SystemEventLog", userID)
	if err != nil {
		return
	}

	for rows.Next() {
		xmlLog := &systemEventLog{}
		var timestamp time.Time
		err = rows.Scan(
			&timestamp, &xmlLog.Server, &xmlLog.TransactionNum,
			&xmlLog.Command, &xmlLog.Username, &xmlLog.StockSymbol,
			&xmlLog.Filename, &xmlLog.Funds,
		)
		if err != nil {
			log.Println("Error scanning trigger: ", err)
		}
		xmlLog.Timestamp = timestamp.UnixNano() / 1000000
		output, err := xml.MarshalIndent(xmlLog, "\t", "\t")
		if err != nil {
			log.Println("Error marshalling to XML: ", err)
		}
		f.Write(output)
		f.WriteString("\n")
	}
	rows.Close()

	// ErrorEvent
	rows, err = getRows("ErrorEventLog", userID)
	if err != nil {
		return
	}

	for rows.Next() {
		xmlLog := &errorEventLog{}
		var timestamp time.Time
		err = rows.Scan(
			&timestamp, &xmlLog.Server, &xmlLog.TransactionNum,
			&xmlLog.Command, &xmlLog.Username, &xmlLog.StockSymbol,
			&xmlLog.Filename, &xmlLog.Funds, &xmlLog.ErrorMessage,
		)
		if err != nil {
			log.Println("Error scanning trigger: ", err)
		}
		xmlLog.Timestamp = timestamp.UnixNano() / 1000000
		output, err := xml.MarshalIndent(xmlLog, "\t", "\t")
		if err != nil {
			log.Println("Error marshalling to XML: ", err)
		}
		f.Write(output)
		f.WriteString("\n")
	}
	rows.Close()

	// DebugEvent
	rows, err = getRows("DebugEventLog", userID)
	if err != nil {
		return
	}

	for rows.Next() {
		xmlLog := &debugEventLog{}
		var timestamp time.Time
		err = rows.Scan(
			&timestamp, &xmlLog.Server, &xmlLog.TransactionNum,
			&xmlLog.Command, &xmlLog.Username, &xmlLog.StockSymbol,
			&xmlLog.Filename, &xmlLog.Funds, &xmlLog.DebugMessage,
		)
		if err != nil {
			log.Println("Error scanning trigger: ", err)
		}
		xmlLog.Timestamp = timestamp.UnixNano() / 1000000
		output, err := xml.MarshalIndent(xmlLog, "\t", "\t")
		if err != nil {
			log.Println("Error marshalling to XML: ", err)
		}
		f.Write(output)
		f.WriteString("\n")
	}
	rows.Close()

	f.WriteString("</log>\n")

	log.Println("Finished dumping logs")
}
