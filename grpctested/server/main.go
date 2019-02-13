package main

import (
	"context"
	"net"

	"github.com/acsaba22/go/grpc/mempb"
	"google.golang.org/grpc"
)

type MemServer struct {
	Mem map[string]string
}

func (ms *MemServer) Get(c context.Context, req *mempb.GetRequest) (*mempb.GetResponse, error) {
	res := mempb.GetResponse{}
	res.RequestedKey = req.Key
	res.Value, res.Exists = ms.Mem[req.Key]
	// time.Sleep(2 * time.Second)
	return &res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	ms := MemServer{map[string]string{"42": "life", "300": "sparta"}}
	mempb.RegisterMemServerServer(s, &ms)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
