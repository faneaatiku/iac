package iac

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewEmptyResponse(t *testing.T) {
	resp := NewEmptyResponse()
	require.NotNil(t, resp)
}

func TestNewResponseSetters(t *testing.T) {
	resp := NewEmptyResponse()
	require.NotNil(t, resp)

	var data []string
	resp.SetData(data)
	require.Equal(t, resp.Data, data)
	meta := NewMetadata(int64(1))
	resp.SetMetadata(*meta)
	require.Equal(t, resp.Metadata, *meta)

	msg := NewGenericMessage("test")
	resp.AddMessage(*msg)
	require.Contains(t, resp.Messages, *msg)
}

func TestNewResponse(t *testing.T) {
	var data []string
	resp := NewResponse(data)
	require.NotNil(t, resp)
	require.Equal(t, resp.Data, data)
}

func TestBuildResponse(t *testing.T) {
	var data []string
	meta := NewMetadata(int64(1))
	msg := NewGenericMessage("test")
	msg2 := NewGenericMessage("testache")
	resp := BuildResponse(
		WithData(data),
		WithMetadata(*meta),
		WithMessages([]Message{*msg}),
		WithMsg(*msg2),
		WithGenericMsg("generic"),
		WithContextualMsg("contextual", "path"),
	)

	require.Equal(t, resp.Metadata, *meta)
	require.Contains(t, resp.Messages, *msg)
}
