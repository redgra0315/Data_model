package main

import ()

func main() {
	// 从配置文件读取配置

	//gin.SetMode(gin.ReleaseMode)

	// 装载路由
	r := NewRouter()
	r.Run(":3000")
	defer DB.Close()
}
