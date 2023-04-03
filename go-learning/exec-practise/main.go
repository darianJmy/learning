package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func upper(data string) string {
	return strings.ToUpper(data)
}

func main() {
	cmd := exec.Command("bash", "-c", "ls")

	//标准输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	//启动
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	//读标准输出流
	data, err := ioutil.ReadAll(stdout)
	if err != nil {
		panic(err)
	}

	//等待执行结束
	if err := cmd.Wait(); err != nil {
		panic(err)
	}

	//小写变大些
	fmt.Printf("%s\n", upper(string(data)))
}
