package hashing

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed EncodedImageExample1.txt
var File1 string //TODO: Replace with a publicly accessible file

//go:embed EncodedImageExample2.txt
var File2 string //TODO: Replace with a publicly accessible file

//go:embed EncodedImageExample3.txt
var File3 string //TODO: Replace with a publicly accessible file

func TestHashingFile1(t *testing.T) {
	bytes := make([]byte, base64.StdEncoding.DecodedLen(len(File1)))
	n, err := base64.StdEncoding.Decode(bytes, []byte(File1))
	if err != nil {
		fmt.Println(err)
		return
	}

	imageHash := Hash{}
	hash := imageHash.HashPageOfDocument(bytes[:n])
	str := imageHash.ConvertHashToString(hash)
	fmt.Println(str)
	assert.Equal(t, "BOKEAeQ48cL", str)
}

func TestHashingFile2(t *testing.T) {
	bytes := make([]byte, base64.StdEncoding.DecodedLen(len(File2)))
	n, err := base64.StdEncoding.Decode(bytes, []byte(File2))
	if err != nil {
		fmt.Println(err)
		return
	}

	imageHash := Hash{}
	hash := imageHash.HashPageOfDocument(bytes[:n])
	str := imageHash.ConvertHashToString(hash)
	fmt.Println(str)
	assert.Equal(t, "Mf0zt0s3a", str)
}

func TestHashingFile3(t *testing.T) {
	bytes := make([]byte, base64.StdEncoding.DecodedLen(len(File3)))
	n, err := base64.StdEncoding.Decode(bytes, []byte(File3))
	if err != nil {
		fmt.Println(err)
		return
	}

	imageHash := Hash{}
	hash := imageHash.HashPageOfDocument(bytes[:n])
	str := imageHash.ConvertHashToString(hash)
	fmt.Println(str)
	assert.Equal(t, "6WMIm2KskNj", str)
}
