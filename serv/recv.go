package serv

import (
	"context"
	"errors"

	"github.com/KVRes/District/rpc"
	"google.golang.org/grpc"
)

func (s *Server) _rcvMsg(req *rpc.ReceiveMessageRequest) (string, error) {
	ch, ok := s.ex.GetCh(req.GetNamespace())
	if !ok {
		return "", errors.New("namespace not found")
	}
	msg := ch.Recv()
	return msg, nil
}

func (s *Server) ReceiveMessage(req *rpc.ReceiveMessageRequest, svr grpc.ServerStreamingServer[rpc.ReceiveMessageResponse]) error {
	msg, err := s._rcvMsg(req)
	if err != nil {
		return err
	}
	return svr.Send(&rpc.ReceiveMessageResponse{Msg: msg})
}

func (s *Server) ReceiveMessageOptimistic(_ context.Context, req *rpc.ReceiveMessageRequest) (*rpc.ReceiveMessageResponse, error) {
	msg, err := s._rcvMsg(req)
	if err != nil {
		return nil, err
	}
	return &rpc.ReceiveMessageResponse{Msg: msg}, nil
}

func (s *Server) ReceiveMessageStream(svr grpc.BidiStreamingServer[rpc.ReceiveMessageRequest, rpc.ReceiveMessageResponse]) error {
	for {
		req, err := svr.Recv()
		if err != nil {
			return err
		}
		msg, err := s._rcvMsg(req)
		if err != nil {
			return err
		}
		err = svr.Send(&rpc.ReceiveMessageResponse{Msg: msg})
		if err != nil {
			return err
		}
	}
}
