package main

/*
#include <stdio.h>
*/
import "C"

// 实现在Go中调用C语言标准库的 puts 函数
func main() {
	cstr := C.CString("Hello World")
	C.puts(cstr)
}
