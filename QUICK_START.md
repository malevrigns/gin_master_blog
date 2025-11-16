# 快速启动指南

## 前置要求

1. **Go 1.21+** 已安装
2. **Node.js 16+** 和 **npm** 已安装
3. **MySQL 5.7+** 已安装并运行
4. **数据库已创建**：
   ```sql
   CREATE DATABASE blog_system CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
   ```

## 启动步骤

### 1. 启动后端服务器

```bash
# 进入后端目录
cd go_projects

# 安装依赖（首次运行）
go mod download

# 检查配置文件
# 确保 config/config.yaml 存在，数据库配置正确

# 启动服务器
go run main.go
```

**成功标志**：看到 `Server starting on port 8080`

### 2. 启动前端开发服务器

**新开一个终端窗口**：

```bash
# 进入前端目录
cd vue_projects

# 安装依赖（首次运行）
npm install

# 启动开发服务器
npm run dev
```

**成功标志**：看到 `Local: http://localhost:3000`

### 3. 访问应用

打开浏览器访问：`http://localhost:3000`

## 常见错误解决

### ❌ `ECONNREFUSED` 错误

**原因**：后端服务器未启动或无法连接

**解决**：
1. 检查后端服务器是否正在运行
2. 检查后端是否在 8080 端口启动
3. 检查 `config/config.yaml` 中的端口配置

### ❌ 数据库连接失败

**解决**：
1. 确保 MySQL 服务正在运行
2. 检查 `config/config.yaml` 中的数据库配置
3. 确保数据库 `blog_system` 已创建

### ❌ 路由冲突错误

**解决**：检查 `go_projects/routes/routes.go` 是否有重复的路由定义

### ❌ 文件上传需要认证

**说明**：文件上传功能需要用户登录后才能使用

**解决**：
1. 先登录（如果有账号）
2. 或者注册新账号
3. 登录后再上传头像和背景图片

## 开发模式

建议使用两个终端：
- **终端1**：运行后端（`go run main.go`）
- **终端2**：运行前端（`npm run dev`）

这样可以看到两边的日志输出，方便调试。

## 测试 API

后端启动后，可以测试 API：
- 获取标签：`http://localhost:8080/api/tags`
- 获取文章：`http://localhost:8080/api/articles`

如果返回 JSON 数据，说明后端正常运行。

