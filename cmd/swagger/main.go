package main

import (
	"fmt"
	"os/exec"

	"web_backend.com/m/v2/configs"
)

func main() {
	// 执行 swag fmt 命令
	if err := runCommand("swag", "fmt"); err != nil {
		fmt.Printf(configs.CommandRed+"Error running swag fmt: %v\n", err)
		return
	}

	// 执行 swag init 命令
	dir := "./cmd/server,./cmd/server,./internal/app/controller,./internal/app/model,./tools" // 需要注释转换swagger的文件
	if err := runCommand("swag", "init", "--dir", dir, "-o", "./api/swagger"); err != nil {
		fmt.Printf(configs.CommandRed+"Error running swag init: %v\n", err)
		return
	}

	fmt.Println(configs.CommandGreen + "swagger文档生成成功！")
}
func runCommand(name string, args ...string) error {
	// 创建命令对象
	cmd := exec.Command(name, args...)
	// 获取命令的标准输出和标准错误
	output, err := cmd.CombinedOutput()
	// 打印输出
	fmt.Printf(configs.CommandGreen+"Running command: %s %v\n", name, args)
	fmt.Printf(configs.CommandYellow+"Output:\n%s", string(output))
	return err
}
