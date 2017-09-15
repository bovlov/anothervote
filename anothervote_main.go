package main

import (
	"./exec"
	//_ "github.com/henrylee2cn/pholcus_lib" // 此为公开维护的spider规则库
	// _ "github.com/henrylee2cn/pholcus_lib_pte" // 同样你也可以自由添加自己的规则库
	_ "../pholcusrules"
)

func main() {
	// 使用web模式运行
	exec.DefaultRun("web")
}
