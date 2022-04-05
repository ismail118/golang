package _40_embed_file_ke_byte

import (
	_ "embed"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed img.png
var logo []byte

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("logo_next.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
