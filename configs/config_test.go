package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var c *Engine

func TestMain(m *testing.M) {
	c = LoadConfig("./")
	os.Exit(m.Run())
}

func TestReadAppServer(t *testing.T) {
	var a *AppServer
	err := c.Read("AppServer", &a)
	require.NoError(t, err)
	require.NotNil(t, a)
	require.EqualValues(t, a.Addr, ":8080")
	require.EqualValues(t, a.RunMode, "debug")
	require.EqualValues(t, a.ReadTimeout, 10)
	require.EqualValues(t, a.WriteTimeout, 10)
}

func TestReadBlogServer(t *testing.T) {
	var a *BlogServer
	err := c.Read("BlogServer", &a)
	require.NoError(t, err)
	require.NotNil(t, a)
	require.EqualValues(t, a.Addr, ":8081")
	require.EqualValues(t, a.IxugoDomain, "localhost")
}
