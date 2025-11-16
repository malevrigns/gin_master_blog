# 后端服务器启动指南

## 快速启动

1. **确保 MySQL 数据库已启动并创建了数据库**
   ```sql
   CREATE DATABASE blog_system CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
   ```

2. **检查配置文件**
   - 确保 `config/config.yaml` 存在
   - 检查数据库配置是否正确

3. **安装依赖（如果还没安装）**
   ```bash
   cd go_projects
   go mod download
   ```

4. **启动后端服务器**
   ```bash
   cd go_projects
   go run main.go
   ```

   如果看到以下信息，说明启动成功：
   ```
   配置文件加载成功: config/config.yaml
   Connecting to MySQL database: root@localhost:3306/blog_system
   Server starting on port 8080
   ```

## 常见问题排查

### 1. 端口被占用
如果 8080 端口被占用，可以：
- 修改 `config/config.yaml` 中的 `server.port` 为其他端口（如 8081）
- 同时修改 `vue_projects/vite.config.js` 中的代理目标端口

### 2. 数据库连接失败
- 检查 MySQL 服务是否运行
- 检查 `config/config.yaml` 中的数据库配置
- 确保数据库 `blog_system` 已创建

### 3. 编译错误
- 运行 `go mod tidy` 更新依赖
- 检查 Go 版本（需要 1.21+）

### 4. 路由冲突
如果遇到路由冲突错误，检查 `routes/routes.go` 是否有重复的路由定义

## 开发建议

1. **使用两个终端窗口**：
   - 终端1：运行后端服务器（`go run main.go`）
   - 终端2：运行前端开发服务器（`npm run dev`）

2. **检查日志**：
   - 后端日志会显示请求信息和错误
   - 前端控制台会显示代理错误

3. **测试连接**：
   启动后端后，可以在浏览器访问 `http://localhost:8080/api/tags` 测试 API 是否正常

