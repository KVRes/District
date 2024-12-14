package serv

import (
	"context"

	"github.com/KVRes/District/rpc"
	"google.golang.org/grpc"
)

func (s *Server) SendMessage(ctx context.Context, req *rpc.SendMessageRequest) (*rpc.SendMessageResponse, error) {
	return s.SendMessageOptimistic(ctx, req)
}

func (s *Server) SendMessageOptimistic(_ context.Context, req *rpc.SendMessageRequest) (*rpc.SendMessageResponse, error) {
	if err := s._sendMsg(req); err != nil {
		return nil, err
	}
	return &rpc.SendMessageResponse{}, nil
}

func (s *Server) SendMessagePessimistic(req *rpc.SendMessageRequest, svr grpc.ServerStreamingServer[rpc.SendMessageResponse]) error {
	if err := s._sendMsg(req); err != nil {
		return err
	}
	return svr.Send(&rpc.SendMessageResponse{})
}

func (s *Server) SendMessageStream(svr grpc.BidiStreamingServer[rpc.SendMessageRequest, rpc.SendMessageResponse]) error {
	for {
		req, err := svr.Recv()
		if err != nil {
			return err
		}
		if err := s._sendMsg(req); err != nil {
			return err
		}
	}
}
