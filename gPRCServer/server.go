package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

const port = ":50051"

type server struct {
	UnimplementedMessageSenderServer
}

func main() {

	list, _ := net.Listen("tcp", port)

	srv := grpc.NewServer()

	RegisterMessageSenderServer(srv, &server{})

	if err := srv.Serve(list); err != nil {

		fmt.Println(err)

	}

}

func (s *server) Send(c context.Context, req *MessageReq) (*MessageResp, error) {

	resp := &MessageResp{}

	fmt.Println("Req", req.SaySomething)

	resp.ResponseSomething = req.SaySomething + "捷哥不要啊~"

	return resp, nil

}
