package app

import "log"

func Start() {
	// 初始化应用配置
	cfg := InitAppConfig()
	// 初始化应用路由
	app := InitRoutes()
	// 应用地址和端口
	addr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Println("启动服务成功，访问地址为：" + "http://" + addr)
	if err := app.Run(addr); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}
