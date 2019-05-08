// +build ignore

package main

import (
	// "bufio"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

// version 1 标准输入方式
// Dupl prints the text of each line that appears more than 
// once in the standard input, preceded by its count.
//
// func main(){
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan(){
// 		// 控制输入程序退出
// 		if input.Text() == "end" {
// 			break
// 		}
//		// 相同行 int+1
// 		counts[input.Text()]++
// 	}
// 	// NOTE: ignoring potential errors from input.Err()
// 	for line, n := range counts{
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// version 2 加入文件输入方式
// Dup2 prints the count and test of lines that appears more than once in the input
// It reads from stdin or from a list of named files
//
// func main(){
// 	counts := make(map[string]int)
// 	files := os.Args[1:]
// 	if len(files) == 0 {
// 		// if do not have file ,user should input test of lines
// 		countLines(os.Stdin, counts)
// 	} else {
// 		for _, arg := range files {
// 			f, err := os.Open(arg)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
// 				continue
// 			}
// 			countLines(f, counts)
// 			f.Close()
// 		}
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\t%s\n", n, lines)
// 		}
// 	}
// }

// func countLines(f *os.File, counts map[string]int){
// 	input := bufio.NewScanner(f)
// 	for input.Scan(){
// 		if input.Text() == "end" { break }
// 		counts[input.Text()]++
// 	}
// }


// version3 只读指定文件，不读标准输入， 引入ReadFile函数

func main(){
	counts := make(map[string]int)
	for _, filename := range os.Args[1:]{
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n"){
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
