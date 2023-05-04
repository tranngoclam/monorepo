package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFoo(t *testing.T) {
	code := foo()
	require.Zero(t, code)
}
