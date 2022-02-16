package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	s := emoji.Sprint("hello :world_map:")
	//fmt.Println(s)
	return s
}
