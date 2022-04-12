package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/portoedu/codePix/application/grpc/pb"
	"github.com/portoedu/codePix/application/usecase"
	"github.com/portoedu/codePix/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: pixRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("nao startou servidor grpc", err)
	}

	log.Printf("gRPC startou na porta %d", port)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("nao startou servidor grpc", err)
	}
}
