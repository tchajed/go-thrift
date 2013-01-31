package main

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/tchajed/go-thrift"
	"github.com/tchajed/go-thrift/examples/scribe"
)

// implementation

type scribeServiceImplementation int

func (s *scribeServiceImplementation) Log(messages []*scribe.LogEntry) (scribe.ResultCode, error) {
	for _, m := range messages {
		fmt.Printf("MSG: %+v\n", m)
	}
	return scribe.ResultCodeOk, nil
}

func main() {
	scribeService := new(scribeServiceImplementation)
	rpc.RegisterName("Thrift", &scribe.ScribeServer{scribeService})

	ln, err := net.Listen("tcp", ":1463")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("ERROR: %+v\n", err)
			continue
		}
		fmt.Printf("New connection %+v\n", conn)
		go rpc.ServeCodec(thrift.NewServerCodec(thrift.NewFramedReadWriteCloser(conn, 0), thrift.NewBinaryProtocol(true, false)))
	}
}
