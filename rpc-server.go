package main

import (
	"fmt"
	"go-rpc/common"
	"log"
	"net"
	"net/rpc"
)

const (
	rpcPort = "5001"
)

type RPCServerCallback struct{}

func main() {
	// Register the RPC Server
	err := rpc.Register(new(RPCServerCallback))
	common.Check(err)
	rpcListen()
}

func rpcListen() {
	log.Println("Starting RPC server on port ", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	common.Check(err)
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}

func (r *RPCServerCallback) LogInfo(payload common.RPCPayload, resp *string) error {
	log.Println(payload)
	*resp = "Processed payload via RPC: " + payload.Name
	return nil
}
