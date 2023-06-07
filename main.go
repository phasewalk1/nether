package main

import (
	"fmt"
	"log"
	pb "proto/proto.gen"
	gimpl "server"
)

func main() {
	m := pb.MsgInTransit {
		Sender: "me",
		Recipient: "you",
		Body: "hello",
	}	

	defer gimpl.Serve()

	fmt.Println(&m)
	log.Default()
	log.Println("gRPC server started on port 50051")
}