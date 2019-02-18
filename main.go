package main

import (
	"fmt"

	"github.com/linkedin/goavro"
)

func main() {
	codec, err := goavro.NewCodec(``)
	fmt.Println("vim-go")
	fmt.Println(codec)
	fmt.Println(err)
}
