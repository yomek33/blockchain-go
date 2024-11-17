package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timeStamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timeStamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}
