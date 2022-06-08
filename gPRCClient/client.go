package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "127.0.0.1:50051"

func main() {

	conn, _ := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	defer conn.Close()

	c := NewMessageSenderClient(conn)

	fmt.Println(c.Send(context.Background(), &MessageReq{
		SaySomething: "634563",
	}))

}
