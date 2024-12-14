package serv

import (
	"context"
	"errors"

	"github.com/KVRes/District/exchange"
	"github.com/KVRes/District/rpc"
)

type Server struct {
	rpc.UnsafeChannelServiceServer
	ex *exchange.ExchangeT[string]
}

func NewServer() *Server {
	return &Server{
		ex: exchange.NewExchange[string](),
	}
}

func (s *Server) Info(_ context.Context, req *rpc.InfoRequest) (*rpc.InfoResponse, error) {
	ch, ok := s.ex.GetCh(req.GetNamespace())
	if !ok {
		return nil, errors.New("namespace not found")
	}
	return &rpc.InfoResponse{
		IsRegister: true,
		Buf:        int32(ch.Cap()),
		Len:        int32(ch.Len()),
		Flags:      ch.Meta().Flags,
	}, nil
}

func (s *Server) RegisterChannel(_ context.Context, req *rpc.RegisterChannelRequest) (*rpc.RegisterChannelResponse, error) {
	meta := exchange.ChMeta{
		Buf:   int(req.GetBuf()),
		Flags: req.GetFlags(),
	}

	existed := s.ex.Register(req.GetNamespace(), meta)
	return &rpc.RegisterChannelResponse{
		Existed: existed,
	}, nil
}

func (s *Server) _sendMsg(req *rpc.SendMessageRequest) error {
	ch, ok := s.ex.GetCh(req.GetNamespace())
	if !ok {
		return errors.New("namespace not found")
	}
	ch.Send(req.GetMsg())
	return nil
}

var _ rpc.ChannelServiceServer = &Server{}
