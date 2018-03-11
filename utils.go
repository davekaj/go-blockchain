package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// //IntToHex turns a int into its hexidecimal byte representation
// func IntToHex(i int64) []byte {
// 	hexRepresentation := strconv.FormatInt(i, 16)

// 	byteArray := []byte(hexRepresentation)

// 	return byteArray
// }

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
