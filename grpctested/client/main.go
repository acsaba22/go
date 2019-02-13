package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/acsaba22/go/grpctested/mempb"

	"google.golang.org/grpc"
)

func clientDo(client mempb.MemServerClient, k string, w io.Writer) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Get(ctx, &mempb.GetRequest{Key: k})
	if err != nil {
		panic(err)
	}
	v := r.Value
	if !r.Exists {
		v = "NOTFOUND"
	}
	fmt.Fprintf(w, "%s: %s\n", r.RequestedKey, v)
}

func main() {
	if len(os.Args) <= 1 {
		panic("Give one argument e.g. 42")
	}
	k := os.Args[1]

	channel, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer channel.Close()
	client := mempb.NewMemServerClient(channel)

	clientDo(client, k, os.Stdout)
}
