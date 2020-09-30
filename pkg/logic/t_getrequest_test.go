package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeGETRequest(t *testing.T) {
	_, errGET := makeGETRequest(url1)

	// TODO: add verification according to time period
	assert.Error(t, errGET)

}
