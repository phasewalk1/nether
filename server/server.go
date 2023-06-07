package server

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "proto/proto.gen"

	"google.golang.org/grpc"
)

type messengerServer struct {}

func newServer() *messengerServer {
	return &messengerServer{}
}

func (s *messengerServer) Send(ctx context.Context, in *pb.MsgInTransit) (*pb.SendResponse, error) {
	return nil, nil
}

func (s *messengerServer) Get(ctx context.Context, in *pb.MsgRequest) (*pb.Msg, error) {
	return nil, nil
}

func (s *messengerServer) Sent(ctx context.Context, in *pb.SentMsgsRequest) (*pb.MultiMsgResponse, error) {
	return nil, nil
}

func (s *messengerServer) Received(ctx context.Context, in *pb.ReceivedMsgsRequest) (*pb.MultiMsgResponse, error) {
	return nil, nil
}

func Serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMessengerServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}