package main

import (
	"github.com/bovlov/anothervote/exec"
	_ "github.com/bovlov/anothervote/otherrule"
)

func main() {
	// 仅使用web模式运行
	exec.DefaultRun("web")
}

/*
	TODO:
	#Auto find Proxy
	Auto switch Proxy
	Windows task bar Tray
	Windows tray Menu
	Daemonize
	Windows 32 bit | done
	Auto get task and work, and commit
	Auto distinguish repeatus task
	Auto update self
	#With a server, version control
	#Bitpay And other vitual currency
	#Try to become a proxy server
	#Download a webserver container and run
	#slave client use proxy add ads
*/
