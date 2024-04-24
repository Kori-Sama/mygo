package grpc

import (
	"mygo/config"
	"mygo/internal/grpc/pb"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Run() {
	port := ":" + config.Server.GrpcPort
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("GRPC server failed to listen:", err)
	}

	log.Infof("GRPC server is running on %s\n", port)

	server := grpc.NewServer()
	pb.RegisterTransactionServiceServer(server, new(transactionService))
	server.Serve(listen)
}
