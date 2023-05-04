package lib

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	require.Equal(t, 3, SumInt(1, 2))
}
