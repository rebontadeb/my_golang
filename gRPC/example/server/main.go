package main

import (
	"net"
	"log"

	pb "github.com/rebontadeb/my_golang/gRPC/example/proto"
	"google.golang.org/grpc"
)

const (
	port = ":18080"
)

type helloServer struct {
	pb.GreetServiceServer
	
}

func main(){
	lis,err := net.Listen("tcp",port)
	if err != nil{

		log.Fatalf("Failed to start the server %v",err)
	}

	grpcServer := grpc.NewServer()

	//Register --word is a kind of keyword to register procol_buffer
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{}) 
	log.Printf("ServerStarted at %v",lis.Addr())
	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Failed to start the server %v",err)
	}
}