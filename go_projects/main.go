package main

import (
	"blog-system/config"
	"blog-system/database"
	"blog-system/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 检查配置文件是否存在，如果不存在则创建
	if _, err := os.Stat("config/config.yaml"); os.IsNotExist(err) {
		log.Println("配置文件不存在，尝试创建默认配置文件...")
		if err := config.CreateDefaultConfig(); err != nil {
			log.Printf("创建默认配置文件失败: %v", err)
			log.Println("请手动创建 config/config.yaml 文件，参考 config/config.yaml.example")
			log.Println("或复制 config/config.yaml.example 为 config/config.yaml")
			return
		} else {
			log.Println("已创建默认配置文件 config/config.yaml")
			log.Println("请修改配置文件中的数据库配置后重新运行程序")
			return
		}
	}

	// 加载配置
	config.LoadConfig()

	// 初始化数据库
	database.InitDB()

	// 设置Gin模式
	gin.SetMode(config.AppConfig.ServerMode)

	// 设置路由
	r := routes.SetupRoutes()

	// 启动服务器
	log.Printf("Server starting on port %s", config.AppConfig.ServerPort)
	if err := r.Run(":" + config.AppConfig.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

