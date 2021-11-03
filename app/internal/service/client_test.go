package service

import (
	"context"
	"os"
	"testing"
	"time"
	"together/configs"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s := New(&configs.AppServer{BlogServerAddr: ":8081"})
	r, err := s.SayHello(ctx, "together")
	require.NoError(t, err)
	require.NotNil(t, r)
	require.EqualValues(t, r.GetMessage(), "Hello together")
}
