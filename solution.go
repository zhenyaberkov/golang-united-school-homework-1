package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	s := emoji.Sprint("hello :world_map:")
	fmt.Println(s)
}
