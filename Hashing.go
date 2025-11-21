package main

import (
	"fmt"
	"strconv"
)

type ImageHash struct {
}

var fnvOffsetBasis uint64 = 14695981039346656037
var fnvPrime uint64 = 1099511628211

// ConvertByteArrToNumberFNV1A By using the FNV-1A
func (ImageHash) ConvertByteArrToNumberFNV1A(bytes []byte) uint64 {
	/*
		algorithm fnv-1a is

			hash := FNV_offset_basis

			for each byte_of_data to be hashed do
			    hash := hash XOR byte_of_data
			    hash := hash × FNV_prime

			return hash
	*/

	var hash = fnvOffsetBasis

	for _, byt := range bytes {
		hash = hash ^ uint64(byt)
		hash = hash * fnvPrime
	}

	return hash
}

// ConvertByteArrToNumberSum By summing the bytes you get a decimal representation of the overall image values. Swapping bits in an image would result in the same number,
func (ImageHash) ConvertByteArrToNumberSum(bytes []byte) uint64 {
	var k uint64 = 0
	for _, byt := range bytes {
		k = k + uint64(byt)
	}

	return k
}

func (t ImageHash) HashPageOfDocument(bytes []byte) uint64 {
	k := t.ConvertByteArrToNumberFNV1A(bytes)
	return t.Hash64shift(k)
}

// Hash64shift based on https://gist.github.com/badboy/6267743#64-bit-mix-functions, preformed an optimised Thomas Wang's 64-bit Integer Hash Function
func (ImageHash) Hash64shift(key uint64) uint64 {
	key = (key) + (key << 21)
	key = key ^ (key >> 24)
	key = (key + (key << 3)) + (key << 8) // key * 265
	key = key ^ (key >> 14)
	key = (key + (key << 2)) + (key << 4) // key * 21
	key = key ^ (key >> 28)
	key = key + (key << 31)
	return key
}

// ConvertHashToString uses base62 to compress and convert the uint64 number into a more human-readable form.
func (ImageHash) ConvertHashToString(input uint64) string {
	s := ""
	charSet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	for i := input; i > 0; i = i / 62 {
		remainer := i % 62
		if remainer <= 9 {
			s = strconv.FormatUint(remainer, 10)
		} else {
			c := string(charSet[remainer])
			s = c + s
		}
	}

	return s
}

// ConvertStringToHash uses base62 to uncompress and revert the base62 string back into its original form.
func (ImageHash) ConvertStringToHash(input uint64) string {
	s := ""
	charSet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	for i := input; i > 0; i = i / 62 {
		remainer := i % 62
		fmt.Println(i)
		fmt.Println(remainer)
		fmt.Println("-------")

		if remainer <= 9 {
			s = strconv.FormatUint(remainer, 10)
		} else {
			c := string(charSet[remainer])
			s = c + s
		}
	}

	return s
}
