package tx

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Input struct{}
type Output struct{}

type Transaction struct {
	version  uint32
	inputs   []*Input
	outputs  []*Output
	locktime int64
	testnet  bool
}

func NewTransaction(version uint32, inputs []*Input, outputs []*Output, locktime int64, testnet bool) *Transaction {
	return &Transaction{
		version:  version,
		inputs:   inputs,
		outputs:  outputs,
		locktime: locktime,
		testnet:  testnet}
}

func ParseTransaction(r io.Reader) (*Transaction, error) {
	var versionBytes [4]byte
	n, err := r.Read(versionBytes[:])
	if err != nil {
		return nil, err
	}
	if n != 4 {
		return nil, fmt.Errorf("invalid data")
	}
	version := binary.LittleEndian.Uint32(versionBytes[:])
	return &Transaction{version: version}, nil
}
