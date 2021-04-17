package service

import (
	"os"
	"testing"
	assetsPkg "together/blog_server/pkg/assets"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	assets = assetsPkg.GetInstance()
	os.Exit(m.Run())
}
func TestGetIxugo(t *testing.T) {
	data := getIxugo("localhost")
	require.NotNil(t, data)
}
func TestGetWangbo(t *testing.T) {
	data := getIxugo("127.0.0.1")
	require.NotNil(t, data)
}
