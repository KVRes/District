package district

import (
	"github.com/KVRes/District/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	rpc.ChannelServiceClient
	conn *grpc.ClientConn
}

func NewClient(addr string, opts ...grpc.DialOption) (*Client, error) {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return nil, err
	}
	return NewClientWithConn(conn), nil
}

func NewClientWithConn(conn *grpc.ClientConn) *Client {
	return &Client{
		ChannelServiceClient: rpc.NewChannelServiceClient(conn),
		conn:                 conn,
	}
}

func (c *Client) Close() {
	c.conn.Close()
}
