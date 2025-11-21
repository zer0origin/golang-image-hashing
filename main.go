package main

import (
	_ "embed"
	"encoding/base64"
	"fmt"
)

//go:embed EncodedImageExample1
var encodedStr string

func main() {
	bytes := make([]byte, base64.StdEncoding.DecodedLen(len(encodedStr)))
	_, err := base64.StdEncoding.Decode(bytes, []byte(encodedStr))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fmt.Sprintf("Image Size: %2.f kb", float64(len(bytes)/1000)))

	imageHash := ImageHash{}
	hash := imageHash.HashPageOfDocument(bytes)
	fmt.Println(hash)

	fmt.Println("\nConvert:")
	str := imageHash.ConvertHashToString(hash)
	fmt.Println(str)
}
