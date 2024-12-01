package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

func main() {
	bc := NewBlockChain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}

// util
// convert an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
