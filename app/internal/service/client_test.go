package service

import (
	"context"
	"os"
	"testing"
	"time"
	"together/configs"
	"together/global"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	global.BlogServer = &configs.BlogServer{
		Addr: ":8081",
	}
	os.Exit(m.Run())
}

func TestClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s := New(ctx)
	r, err := s.SayHello("together")
	require.NoError(t, err)
	require.NotNil(t, r)
	require.EqualValues(t, r.GetMessage(), "Hello together")
}

func TestGetList(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s := New(ctx)
	r, err := s.GetList("https://chenyunxin.cn")
	require.NoError(t, err)
	require.NotNil(t, r)
}
