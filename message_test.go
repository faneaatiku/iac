package iac

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewGenericMessage(t *testing.T) {
	text := "test"
	msg := NewGenericMessage(text)
	require.NotNil(t, msg)
	require.Equal(t, msg.Text, text)
	require.Equal(t, msg.Type, TypeGeneric)
}

func TestNewContextualMessage(t *testing.T) {
	text := "test"
	path := "yes.yes2"
	msg := NewContextualMessage(text, path)
	require.NotNil(t, msg)
	require.Equal(t, msg.Text, text)
	require.Equal(t, msg.Type, TypeContextual)
	require.Equal(t, msg.Path, path)
}

func TestBuildResponseMessage(t *testing.T) {
	msg := BuildResponseMessage()
	require.NotNil(t, msg)

	path := "yes.no"
	mType := TypeContextual
	text := "say no more"
	msg = BuildResponseMessage(WithPath(path), WithType(mType), WithText(text))
	require.NotNil(t, msg)
	require.Equal(t, msg.Text, text)
	require.Equal(t, msg.Type, mType)
	require.Equal(t, msg.Path, path)
}
