package ImageUtil

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed EncodedImageExample1
var File1 string

//go:embed EncodedImageExample2
var File2 string

func TestHashingFile1(t *testing.T) {
	bytes := make([]byte, base64.StdEncoding.DecodedLen(len(File1)))
	_, err := base64.StdEncoding.Decode(bytes, []byte(File1))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fmt.Sprintf("Image Size: %2.f kb", float64(len(bytes)/1000)))

	imageHash := Hash{}
	hash := imageHash.HashPageOfDocument(bytes)
	fmt.Println(hash)

	fmt.Println("\nConvert:")
	str := imageHash.ConvertHashToString(hash)
	fmt.Println(str)

	assert.Equal(t, str, "LmPjeyvClM4")
}

func TestHashingFile2(t *testing.T) {
	bytes := make([]byte, base64.StdEncoding.DecodedLen(len(File2)))
	_, err := base64.StdEncoding.Decode(bytes, []byte(File2))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fmt.Sprintf("Image Size: %2.f kb", float64(len(bytes)/1000)))

	imageHash := Hash{}
	hash := imageHash.HashPageOfDocument(bytes)
	fmt.Println(hash)

	fmt.Println("\nConvert:")
	str := imageHash.ConvertHashToString(hash)
	fmt.Println(str)

	assert.Equal(t, str, "JWqHUeD4")
}
