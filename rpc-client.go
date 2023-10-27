package main

import (
	"go-rpc/common"
	"log"
	"net/rpc"
)

func main() {
	sendMessageViaRPC()
}

func sendMessageViaRPC() {
	client, err := rpc.Dial("tcp", "localhost:5001")
	common.Check(err)

	rpcPayload := common.RPCPayload{
		Name: "Message",
		Data: "Hello World",
	}

	var result string

	err = client.Call("RPCServerCallback.LogInfo", rpcPayload, &result)
	common.Check(err)
	log.Println("rpc message sent!")
}
