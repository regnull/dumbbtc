package tx

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Transaction_Parse(t *testing.T) {
	assert := assert.New(t)

	b := []byte{0x1, 0x0, 0x0, 0x0}
	buf := bytes.NewReader(b)
	txn, err := ParseTransaction(buf)
	assert.NoError(err)
	assert.EqualValues(1, txn.version)
}
