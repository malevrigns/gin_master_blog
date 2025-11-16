package database

import (
	"blog-system/config"
	"blog-system/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	var err error
	var dialector gorm.Dialector

	if config.AppConfig.DBType == "mysql" {
		// MySQL 连接字符串
		dsn := config.AppConfig.DBUser + ":" + config.AppConfig.DBPassword + "@tcp(" + config.AppConfig.DBHost + ":" + config.AppConfig.DBPort + ")/" + config.AppConfig.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
		dialector = mysql.Open(dsn)
		log.Printf("Connecting to MySQL database: %s@%s:%s/%s", config.AppConfig.DBUser, config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBName)
	} else {
		// SQLite 连接（保留兼容性）
		dialector = sqlite.Open(config.AppConfig.DBName)
		log.Printf("Connecting to SQLite database: %s", config.AppConfig.DBName)
	}

	DB, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移
	err = DB.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Category{},
		&models.Tag{},
		&models.Comment{},
		&models.Music{},
		&models.Playlist{},
		&models.Link{},
		&models.SiteConfig{},
		&models.Lab{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected and migrated successfully")
	
	// 填充初始数据
	SeedDatabase()
}

