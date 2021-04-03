package client

import (
	"context"
	"testing"
	"time"
	pb "together/proto"

	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := Instance.SayHello(ctx, &pb.HelloRequest{Name: "together"})
	t.Log(r)
	require.NoError(t, err)
	require.EqualValues(t, r.GetMessage(), "Hello together")

}
