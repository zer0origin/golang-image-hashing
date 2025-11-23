package main

import "C"

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

// ConvertByteArrToNumberFNV1A By using the FNV-1A
//
//export ConvertByteArrToNumberFNV1A
func ConvertByteArrToNumberFNV1A(bytes []byte) uint64 {
	/*
		algorithm fnv-1a is

			hash := FNV_offset_basis

			for each byte_of_data to be hashed do
			    hash := hash XOR byte_of_data
			    hash := hash × FNV_prime

			return hash
	*/
	fmt.Println("--- GO DIAGNOSTICS ---")
	fmt.Printf("Total Length: %d\n", len(bytes))
	if len(bytes) >= 5 {
		fmt.Printf("First 5 bytes: %v\n", bytes[:5])
		fmt.Printf("Last 5 bytes:  %v\n", bytes[len(bytes)-5:])
	}
	fmt.Println(len(bytes))
	fmt.Println(ConvertByteArrToNumberSum(bytes))
	var fnvOffsetBasis uint64 = 14695981039346656037
	var fnvPrime uint64 = 1099511628211

	var hash = fnvOffsetBasis

	for _, byt := range bytes {
		hash = hash ^ uint64(byt)
		hash = hash * fnvPrime
	}

	fmt.Printf("GO: %d\n", +hash)
	return hash
}

// ConvertByteArrToNumberFNV1ABase64Encoded By using the FNV-1A
//
//export ConvertByteArrToNumberFNV1ABase64Encoded
func ConvertByteArrToNumberFNV1ABase64Encoded(encodedData string) uint64 {
	/*
		algorithm fnv-1a is

			hash := FNV_offset_basis

			for each byte_of_data to be hashed do
			    hash := hash XOR byte_of_data
			    hash := hash × FNV_prime

			return hash
	*/
	fmt.Println(encodedData)

	bytes := make([]byte, base64.StdEncoding.DecodedLen(len(encodedData)))
	_, err := base64.StdEncoding.Decode(bytes, []byte(encodedData))
	if err != nil {
		panic("Unable to decode data!")
	}

	fmt.Println("--- GO DIAGNOSTICS ---")
	fmt.Printf("Total Length: %d\n", len(bytes))
	if len(bytes) >= 5 {
		fmt.Printf("First 5 bytes: %v\n", bytes[:5])
		fmt.Printf("Last 5 bytes:  %v\n", bytes[len(bytes)-5:])
	}

	var fnvOffsetBasis uint64 = 14695981039346656037
	var fnvPrime uint64 = 1099511628211

	var hash = fnvOffsetBasis

	for _, byt := range bytes {
		hash = hash ^ uint64(byt)
		hash = hash * fnvPrime
	}

	fmt.Printf("GO: %d\n", +hash)
	return hash
}

// ConvertByteArrToNumberSum By summing the bytes you get a decimal representation of the overall image values. Swapping bits in an image would result in the same number,
//
//export ConvertByteArrToNumberSum
func ConvertByteArrToNumberSum(bytes []byte) uint64 {
	var k uint64 = 0
	for _, byt := range bytes {
		k = k + uint64(byt)
	}

	return k
}

//export HashPageOfDocument
func HashPageOfDocument(bytes []byte) uint64 {
	k := ConvertByteArrToNumberFNV1A(bytes)
	return Hash64shift(k)
}

// Hash64shift based on https://gist.github.com/badboy/6267743#64-bit-mix-functions, preformed an optimised Thomas Wang's 64-bit Integer Hash Function
//
//export Hash64shift
func Hash64shift(key uint64) uint64 {
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
//
//export ConvertHashToString
func ConvertHashToString(input uint64) string {
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

func main() {}

//export convertFNVString
func convertFNVString(cStr *C.char) uint64 {
	data := C.GoString(cStr)
	fmt.Printf("Go received string of length: %d\n", len(data))

	return ConvertByteArrToNumberFNV1ABase64Encoded(data)
}
