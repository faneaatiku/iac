package iac

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewMetadata(t *testing.T) {
	page := int64(2)
	meta := NewMetadata(page)
	require.NotNil(t, page)
	require.Equal(t, meta.CurrentPage, page)
}

func TestBuildResponseMetadata(t *testing.T) {
	page := int64(2)
	prev := int64(2)
	next := int64(2)

	total := int64(2313)
	meta := BuildResponseMetadata(WithCurrentPage(page), WithPreviousPage(prev), WithNextPage(next), WithTotalItems(total))
	require.NotNil(t, page)
	require.Equal(t, meta.CurrentPage, page)
	require.Equal(t, meta.NextPage, next)
	require.Equal(t, meta.PreviousPage, prev)
	require.Equal(t, meta.TotalItems, total)
}
