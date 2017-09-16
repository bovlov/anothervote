package main

import (
	"github.com/bovlov/anothervote/exec"
	//_ "github.com/bovlov/anothervote_lib" // 此为公开维护的spider规则库
	// _ "github.com/bovlov/anothervote_lib_pte" // 同样你也可以自由添加自己的规则库
	_ "./otherrule"
)

func main() {
	// 使用web模式运行
	exec.DefaultRun("web")
}
