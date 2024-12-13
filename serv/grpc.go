package serv

import (
	"errors"
	"log"

	"github.com/KVRes/District/listener"
	"github.com/KVRes/District/rpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

func (s *Server) RunGrpc(svr *grpc.Server, protocol, addr string) error {
	lis, err := listener.Listen(protocol, addr)
	if err != nil {
		return err
	}
	rpc.RegisterChannelServiceServer(svr, s)
	return svr.Serve(lis)
}

func panicHandler(p any) error {
	log.Println("panic", p)
	return errors.New("panic")
}

func (s *Server) Run(protocol, addr string) error {
	svr := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(panicHandler)),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(panicHandler)),
		)))
	return s.RunGrpc(svr, protocol, addr)
}
