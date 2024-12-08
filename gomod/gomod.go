package main

import (
	"fmt"
	"go.uber.org/zap"
	"learn-go/utils"
)

// golang的模块化管理有三种 实际上 gopath go mod 和 vendor
// 在gopath的情况中 所有的依赖都放在 $GOPATH 目录下 然后就会有一种问题
// 如果多个项目都依赖了同一个包的不同版本，那么就会造成冲突
// 所以就出现了vendor 这个概念 就是在不同的gopath 目录下，每个项目都复制一份依赖的包
// 依赖检查的顺序是 首先 goroot 然后 vendor 然后 GOPATH
// 但是使用vendor 还是比较麻烦就又出现了go mod
// go mod 使用很简单 首先安装依赖 go get 然后首先更新项目中的go.mod文件
// 然后把依赖安装到 gopath/pkg/mod 下面并且根据版本号生成不同的目录
// 有时候可以不使用go get 安装依赖，可以直接在代码中使用 依赖 go mod 会自己处理
// 技巧 go build ./... 会把所有依赖都下载下来 这个命令实际上是在检查当前项目下所有的文件是否可以编译通过

func main() {
	logger := zap.NewExample()
	fmt.Printf("a+b=：%d\n", utils.AddAB(1, 2))
	logger.Warn("hello world")
}
