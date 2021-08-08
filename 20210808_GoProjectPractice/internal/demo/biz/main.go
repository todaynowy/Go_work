package main

import (
	业务mod
	数据mod
	apimod
)

func main() {
	//TODO 各类服务初始化

	//TODO 注册业务事件处理服务、数据服务、gRPC中的网络传输API

	//TODO 开启协程启动上述服务，并加channel监听各个服务运行状态

	//TODO channel 监听Linux系统信号SIGINT SIGTERM

	//错误处理协程
	group.Go(func() error {
		for {
			select {
			case 服务出错:
				TODO 退出服务，用一个context对象控制退出周期
				return errCtx.Err()
			case Linux 系统信号:
				cancel()
			}
		}
		return nil
	})

	//错误打印，程序退出
}
