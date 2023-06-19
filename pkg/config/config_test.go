package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewConfig_Exist(t *testing.T) {
	conf := NewConfig("../../config/config.example.yaml")

	require.NotNil(t, conf)
	require.NotEmpty(t, conf)
	require.Equal(t, "local", conf.GetString("env"))
}
