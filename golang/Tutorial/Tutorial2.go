// +build ignore
// Echo1 prints its command-line arguments.

package main

import (
    "fmt"
	"os"
	"strings"
)

// func main() {
//     var s, sep string
//     for i := 1; i < len(os.Args); i++ {
//         s += sep + os.Args[i]
//         sep = " "
//     }
//     fmt.Println(s)
// }

func main(){
	s, sep := "", ""
	for _, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	// 如果连接涉及的数据量很大，这种代价高昂。一种简单且高效的解决方案方式是使用strings包中的Join函数
	fmt.Println(strings.Join(os.Args[1:], " "))
}