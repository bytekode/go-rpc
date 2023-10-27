package main

import (
	"bufio"
	"go-rpc/common"
	"log"
	"net/rpc"
	"os"
)

func main() {
	// sendMessageViaRPC()
	scanner()
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Enter your message..")
	for scanner.Scan() {
		message := scanner.Text()
		if message != "" {
			sendMessageViaRPC(message)
		} else {
			log.Println("Please enter a valid message!")
		}
		log.Println("Enter your message..")
	}
}

func sendMessageViaRPC(message string) {
	client, err := rpc.Dial("tcp", "localhost:5001")
	common.Check(err)

	rpcPayload := common.RPCPayload{
		Name: "Message",
		Data: message,
	}

	var result string

	err = client.Call("RPCServerCallback.LogInfo", rpcPayload, &result)
	common.Check(err)
	log.Println("RPC message sent!")
	log.Println("Message from server =>", result)
}
