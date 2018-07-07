package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)
	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}
	//解决乱码
	cmd.Stdout = os.Stdout
	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			fmt.Println(err2)
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}

func exec_shell(git_url string, source_name string) {
	fmt.Println("----------- begin clone-----------")
    execCommand("git", []string{"clone", git_url, source_name})
	fmt.Println("----------- end -----------")
}

func main() {
	args := os.Args //获取用户输入的所有参数
	if args == nil || len(args) < 2 {
		fmt.Println("-------------------- windows --------------------------------")
		fmt.Println("example1: gopull https://github.com/yimogit/gotest %GOPATH%\\src\\github.com\\gotest")
		fmt.Println("example2: gopull -g https://github.com/yimogit/gotest -gs %GOPATH%\\src\\github.com\\gotest")
		fmt.Println("--------------------  linux  --------------------------------")
		fmt.Println("example3: gopull https://github.com/yimogit/gotest $GOPATH/src/xxxxx.com/gotest")
		fmt.Println("----------------------------------------------------")
		return
	}
	github_path := flag.String("g", "", "github url:https://github.com/yimogit/gotest")
	gopath_dir := flag.String("gs", "", "%GOPATH%\\src\\github.com\\gotest")
	flag.Parse()
	if *github_path == "" && *gopath_dir == "" {
		github_path = &args[1]
		gopath_dir = &args[2]
	}
    exec_shell(*github_path, *gopath_dir)
}
