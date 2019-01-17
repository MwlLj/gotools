package file

import (
	"fmt"
	"io"
	"os"
	"syscall"
)

// 输入流复制到文件里面
func test_stdin() error {
	var out io.Writer
	var in io.Reader
	in = os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	out, _ = os.OpenFile("D:/golang/project/docker/env.txt", os.O_RDWR, 0666)
	n, err := io.Copy(out, in)
	fmt.Printf("\n write %d err %v \n", n, err)
	return err
}

//文件里面复制到输出流
func test_stdout() error {
	var out io.Writer
	var in io.Reader
	out = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	in, _ = os.OpenFile("D:/golang/project/docker/env.txt", os.O_RDWR, 0666)
	n, err := io.Copy(out, in)
	fmt.Printf("\n write %d err %v \n", n, err)

	return err
}
