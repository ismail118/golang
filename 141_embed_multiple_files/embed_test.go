package _41_embed_multiple_files

import (
	"embed"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed file/a.txt
//go:embed file/b.txt
//go:embed file/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("file/a.txt")
	fmt.Println(string(a))
	b, _ := files.ReadFile("file/b.txt")
	fmt.Println(string(b))
	c, _ := files.ReadFile("file/c.txt")
	fmt.Println(string(c))
}
