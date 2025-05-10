package main

import (
	"github.com/BerniceZTT/goadmin/config"
	"github.com/BerniceZTT/goadmin/routes"
)

func main() {
	// 初始化数据库
	config.ConnectDB()
	config.Migrate()

	// 设置路由
	r := routes.SetupRouter()

	// 启动服务
	r.Run(":8080")
}
