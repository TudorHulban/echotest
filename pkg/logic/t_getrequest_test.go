package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMakeGETRequest(t *testing.T) {
	code, errGET := makeGETRequest(url0)
	require.Nil(t, errGET)
	assert.GreaterOrEqual(t, code, 0)
}
