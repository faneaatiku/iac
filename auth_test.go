package iac

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestLoadAuthFromEnv(t *testing.T) {
	pub := "test_public_Key"
	err := os.Setenv("APP_PUBLIC_KEY", pub)
	require.Nil(t, err)

	priv := "test_priv"
	err = os.Setenv("APP_PRIVATE_KEY", priv)
	require.Nil(t, err)

	auth := LoadAuthFromEnv("")
	require.NotNil(t, auth)
	require.Equal(t, auth.PublicKey, pub)
	require.Equal(t, auth.PrivateKey, priv)
}
