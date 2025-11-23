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

func TestHashingFile1(t *testing.T) {
	bytes := make([]byte, base64.StdEncoding.DecodedLen(len(File1)))
	_, err := base64.StdEncoding.Decode(bytes, []byte(File1))
	if err != nil {
		fmt.Println(err)
		return
	}

	imageHash := Hash{}
	hash := imageHash.HashPageOfDocument(bytes)
	str := imageHash.ConvertHashToString(hash)
	assert.Equal(t, str, "LmPjeyvClM4")
}

func TestHashingFile2(t *testing.T) {
	bytes := make([]byte, base64.StdEncoding.DecodedLen(len(File2)))
	_, err := base64.StdEncoding.Decode(bytes, []byte(File2))
	if err != nil {
		fmt.Println(err)
		return
	}

	imageHash := Hash{}
	hash := imageHash.HashPageOfDocument(bytes)
	str := imageHash.ConvertHashToString(hash)
	assert.Equal(t, str, "JWqHUeD4")
}
