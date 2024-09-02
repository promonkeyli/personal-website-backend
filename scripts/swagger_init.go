package main

import (
	"fmt"
	"os/exec"
	"web_backend.com/m/v2/config"
)

func main() {
	// 执行 swag fmt 命令
	if err := runCommand("swag", "fmt"); err != nil {
		fmt.Printf(config.CommandRed+"Error running swag fmt: %v\n", err)
		return
	}

	// 执行 swag init 命令
	dir := "./cmd,./internal/app/controller,./internal/app/model,./utils" // 需要注释转换swagger的文件
	if err := runCommand("swag", "init", "--dir", dir, "-o", "./api/swagger"); err != nil {
		fmt.Printf(config.CommandRed+"Error running swag init: %v\n", err)
		return
	}

	fmt.Println(config.CommandGreen + "swagger文档生成成功！")
}

func runCommand(name string, args ...string) error {
	// 创建命令对象
	cmd := exec.Command(name, args...)
	// 获取命令的标准输出和标准错误
	output, err := cmd.CombinedOutput()
	// 打印输出
	fmt.Printf(config.CommandGreen+"Running command: %s %v\n", name, args)
	fmt.Printf(config.CommandYellow+"Output:\n%s", string(output))
	return err
}
