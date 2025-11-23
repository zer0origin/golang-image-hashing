package main

import (
	_ "embed"
	"testing"
)

//go:embed EncodedImageExample1.txt
var File1 string //TODO: Replace with a publicly accessible file

//go:embed EncodedImageExample2.txt
var File2 string //TODO: Replace with a publicly accessible file

func TestHashingFile1(t *testing.T) {
	//
	//imageHash := Hash{}
	//hash := imageHash.HashPageOfDocument(bytes)
	//fmt.Println(hash)
	//
	//fmt.Println("\nConvert:")
	//str := imageHash.ConvertHashToString(hash)
	//fmt.Println(str)
	//
	//assert.Equal(t, str, "LmPjeyvClM4")

	_ = ConvertByteArrToNumberFNV1ABase64Encoded(File1)
}

func TestHashingFile2(t *testing.T) {
	//bytes := make([]byte, base64.StdEncoding.DecodedLen(len(File2)))
	//_, err := base64.StdEncoding.Decode(bytes, []byte(File2))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(fmt.Sprintf("Image Size: %2.f kb", float64(len(bytes)/1000)))
	//
	//imageHash := Hash{}
	//hash := imageHash.HashPageOfDocument(bytes)
	//fmt.Println(hash)
	//
	//fmt.Println("\nConvert:")
	//str := imageHash.ConvertHashToString(hash)
	//fmt.Println(str)
	//
	//assert.Equal(t, str, "JWqHUeD4")
}
