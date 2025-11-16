# 博客系统后端

基于 Go + Gin 框架的博客系统后端API。

## 环境要求

- Go 1.21+
- MySQL 5.7+ 或 MySQL 8.0+

## 安装和运行

1. **创建 MySQL 数据库**：
```sql
CREATE DATABASE blog_system CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. **配置文件**：
复制 `config/config.yaml.example` 为 `config/config.yaml`，并修改数据库配置：
```yaml
database:
  type: mysql
  host: localhost
  port: 3306
  user: root
  password: your_password
  name: blog_system
```

3. **安装依赖**：
```bash
go mod download
```

4. **运行**：
```bash
go run main.go
```

## 配置说明

配置文件位置：`config/config.yaml`

主要配置项：
- `database.type`: 数据库类型（默认：mysql，可选：sqlite）
- `database.host`: 数据库主机（默认：localhost）
- `database.port`: 数据库端口（默认：3306）
- `database.user`: 数据库用户名（默认：root）
- `database.password`: 数据库密码
- `database.name`: 数据库名称（默认：blog_system）
- `jwt.secret`: JWT密钥（生产环境请修改）
- `server.port`: 服务器端口（默认：8080）
- `server.mode`: 运行模式（默认：debug，可选：release）

详细配置请参考 `config/config.yaml.example`

## API文档

详细API文档请参考主README.md

