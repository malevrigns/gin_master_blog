package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

type JWTConfig struct {
	Secret string `yaml:"secret"`
}

type UploadConfig struct {
	Path    string `yaml:"path"`
	MaxSize int64  `yaml:"max_size"`
}

type MusicConfig struct {
	Path string `yaml:"path"`
}

type ConfigFile struct {
	Database DatabaseConfig `yaml:"database"`
	Server   ServerConfig   `yaml:"server"`
	JWT      JWTConfig      `yaml:"jwt"`
	Upload   UploadConfig   `yaml:"upload"`
	Music    MusicConfig    `yaml:"music"`
}

type Config struct {
	DBType       string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	JWTSecret    string
	ServerPort   string
	ServerMode   string
	UploadPath   string
	MaxUploadSize int64
	MusicPath    string
}

var AppConfig *Config

func LoadConfig() {
	configFile := "config/config.yaml"
	
	// 检查配置文件是否存在
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Printf("配置文件 %s 不存在，使用默认配置", configFile)
		loadDefaultConfig()
		return
	}

	// 读取配置文件
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Printf("读取配置文件失败: %v，使用默认配置", err)
		loadDefaultConfig()
		return
	}

	// 解析YAML
	var configFileData ConfigFile
	if err := yaml.Unmarshal(data, &configFileData); err != nil {
		log.Printf("解析配置文件失败: %v，使用默认配置", err)
		loadDefaultConfig()
		return
	}

	// 转换为应用配置
	AppConfig = &Config{
		DBType:       getValueOrDefault(configFileData.Database.Type, "mysql"),
		DBHost:       getValueOrDefault(configFileData.Database.Host, "localhost"),
		DBPort:       getValueOrDefault(configFileData.Database.Port, "3306"),
		DBUser:       getValueOrDefault(configFileData.Database.User, "root"),
		DBPassword:   configFileData.Database.Password,
		DBName:       getValueOrDefault(configFileData.Database.Name, "blog_system"),
		JWTSecret:    getValueOrDefault(configFileData.JWT.Secret, "your-secret-key-change-this"),
		ServerPort:   getValueOrDefault(configFileData.Server.Port, "8080"),
		ServerMode:   getValueOrDefault(configFileData.Server.Mode, "debug"),
		UploadPath:   getValueOrDefault(configFileData.Upload.Path, "./uploads"),
		MaxUploadSize: configFileData.Upload.MaxSize,
		MusicPath:    getValueOrDefault(configFileData.Music.Path, "./music"),
	}

	// 如果 MaxUploadSize 为0，使用默认值
	if AppConfig.MaxUploadSize == 0 {
		AppConfig.MaxUploadSize = 10485760 // 10MB
	}

	// 创建必要的目录
	os.MkdirAll(AppConfig.UploadPath, os.ModePerm)
	os.MkdirAll(AppConfig.MusicPath, os.ModePerm)

	log.Printf("配置文件加载成功: %s", configFile)
}

func loadDefaultConfig() {
	AppConfig = &Config{
		DBType:       "mysql",
		DBHost:       "localhost",
		DBPort:       "3306",
		DBUser:       "root",
		DBPassword:   "",
		DBName:       "blog_system",
		JWTSecret:    "your-secret-key-change-this",
		ServerPort:   "8080",
		ServerMode:   "debug",
		UploadPath:   "./uploads",
		MaxUploadSize: 10485760,
		MusicPath:    "./music",
	}

	// 创建必要的目录
	os.MkdirAll(AppConfig.UploadPath, os.ModePerm)
	os.MkdirAll(AppConfig.MusicPath, os.ModePerm)
}

func getValueOrDefault(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

// CreateDefaultConfig 创建默认配置文件
func CreateDefaultConfig() error {
	configFile := "config/config.yaml"
	
	// 检查文件是否已存在
	if _, err := os.Stat(configFile); err == nil {
		return fmt.Errorf("配置文件 %s 已存在", configFile)
	}

	// 创建默认配置
	defaultConfig := ConfigFile{
		Database: DatabaseConfig{
			Type:     "mysql",
			Host:     "localhost",
			Port:     "3306",
			User:     "root",
			Password: "",
			Name:     "blog_system",
		},
		Server: ServerConfig{
			Port: "8080",
			Mode: "debug",
		},
		JWT: JWTConfig{
			Secret: "your-secret-key-change-this-in-production",
		},
		Upload: UploadConfig{
			Path:    "./uploads",
			MaxSize: 10485760,
		},
		Music: MusicConfig{
			Path: "./music",
		},
	}

	// 序列化为YAML
	data, err := yaml.Marshal(&defaultConfig)
	if err != nil {
		return fmt.Errorf("序列化配置失败: %v", err)
	}

	// 写入文件
	if err := os.WriteFile(configFile, data, 0644); err != nil {
		return fmt.Errorf("写入配置文件失败: %v", err)
	}

	log.Printf("默认配置文件已创建: %s", configFile)
	return nil
}

