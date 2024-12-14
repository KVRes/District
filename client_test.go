package district_test

import (
	"context"
	"testing"

	district "github.com/KVRes/District"
	"github.com/KVRes/District/rpc"
	"github.com/KVRes/District/serv"
)

const addr = "localhost:8190"

func TestClient(t *testing.T) {
	server := serv.NewServer()
	go server.Run("tcp", addr)

	client, err := district.NewClient(addr)
	if err != nil {
		t.Fatal(err)
	}

	client.RegisterChannel(context.Background(), &rpc.RegisterChannelRequest{
		Namespace: "test",
		Buf:       1024,
	})

	client.SendMessage(context.Background(), &rpc.SendMessageRequest{
		Namespace: "test",
		Msg:       "hello",
	})

	t.Log(client.Info(context.Background(), &rpc.InfoRequest{
		Namespace: "test",
	}))

	stream, err := client.ReceiveMessage(context.Background(), &rpc.ReceiveMessageRequest{
		Namespace: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	msg, err := stream.Recv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(msg)

}
