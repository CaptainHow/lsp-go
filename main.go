package main

import (
	"bufio"
	"encoding/json"
	"golsp/lsp"
	"golsp/rpc"
	"log"
	"os"
)

func main() {

	logger := getLogger("D:/AHMED/projects/Golang/LSP/lsp-go/log.txt")
	logger.Println("Hey, I Started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}
		handleMessage(logger, method, contents)
	}

}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Received msg with method: %s ", method)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Couldn't parse this: %s ", err)
		}

		logger.Printf("Connected to: %s %s ", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)
	}
}

func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Hey, you didn't give me a good file")
	}

	return log.New(logFile, "[GO LSP]", log.Ldate|log.Ltime|log.Lshortfile)
}
