package main

import (
	"bytes"
	"fmt"
)

func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	// var dir1 []byte
	// dir1 = append(dir1, path[:sepIndex]...)

	dir1 := make([]byte, sepIndex)
	copy(dir1, path[:sepIndex])

	// dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB
	dir1 = append(dir1, "suffix"...)

	fmt.Println("dir1 now :", string(dir1))
	fmt.Println("path now:", string(path))
}
