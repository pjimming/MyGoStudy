// 数据从命令输入的文本的文件名，获取里面的内容
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	pos := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, pos, "cmd")
	} else {
		for _, arg := range files {
			// os.Open(string) return *os.File, err
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, pos, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, pos[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, pos map[string]string, fname string) {
	in := bufio.NewScanner(f)
	st := make(map[string]bool)
	for in.Scan() {
		counts[in.Text()]++
		if !st[in.Text()] {
			pos[in.Text()] += fname + " "
		}
		st[in.Text()] = true
	}
	// NOTE: ignoring potential errors from in.Err()
}
