# 博客系统

一个基于 Go + Gin + Vue3 构建的现代化博客系统，参考了 Firefly 博客的设计理念，并添加了音乐播放功能。

## 功能特性

### 后端功能
- ✅ 用户认证（JWT）
- ✅ 文章管理（CRUD）
- ✅ 分类和标签管理
- ✅ 评论系统
- ✅ 音乐管理
- ✅ 播放列表
- ✅ RESTful API

### 前端功能
- ✅ 响应式设计（支持暗色模式）
- ✅ 文章列表和详情页
- ✅ 分类和标签筛选
- ✅ 评论功能
- ✅ 音乐播放器（播放、暂停、上一首、下一首、音量控制、播放列表）
- ✅ 管理后台
- ✅ 搜索功能

## 技术栈

### 后端
- Go 1.21+
- Gin Web框架
- GORM ORM
- JWT认证
- MySQL 数据库（默认，也支持 SQLite）

### 前端
- Vue 3
- Vite
- Vue Router
- Pinia
- Element Plus
- Tailwind CSS
- Marked (Markdown渲染)

## 快速开始

### 后端启动

1. **创建 MySQL 数据库**：
```sql
CREATE DATABASE blog_system CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. 进入后端目录：
```bash
cd go_projects
```

3. 配置文件：
```bash
# 复制配置文件示例并修改配置
cp config/config.yaml.example config/config.yaml
# 编辑 config/config.yaml 文件，设置数据库连接信息
```

4. 安装依赖：
```bash
go mod download
```

5. 运行后端：
```bash
go run main.go
```

后端将在 `http://localhost:8080` 启动

**注意**：首次运行会自动创建数据表并填充初始数据。

### 前端启动

1. 进入前端目录：
```bash
cd vue_projects
```

2. 安装依赖：
```bash
npm install
# 或
pnpm install
```

3. 启动开发服务器：
```bash
npm run dev
# 或
pnpm dev
```

前端将在 `http://localhost:3000` 启动

## 项目结构

```
.
├── go_projects/          # 后端项目
│   ├── config/          # 配置
│   ├── controllers/     # 控制器
│   ├── database/        # 数据库
│   ├── middleware/      # 中间件
│   ├── models/          # 数据模型
│   ├── routes/          # 路由
│   ├── utils/           # 工具函数
│   └── main.go          # 入口文件
│
└── vue_projects/        # 前端项目
    ├── src/
    │   ├── api/         # API服务
    │   ├── components/  # 组件
    │   ├── router/      # 路由
    │   ├── stores/      # 状态管理
    │   ├── views/       # 页面
    │   └── main.js      # 入口文件
    └── package.json
```

## API 文档

### 认证接口
- `POST /api/auth/register` - 用户注册
- `POST /api/auth/login` - 用户登录
- `GET /api/auth/profile` - 获取当前用户信息（需要认证）

### 文章接口
- `GET /api/articles` - 获取文章列表
- `GET /api/articles/:id` - 获取文章详情
- `POST /api/articles` - 创建文章（需要认证）
- `PUT /api/articles/:id` - 更新文章（需要认证）
- `DELETE /api/articles/:id` - 删除文章（需要认证）
- `POST /api/articles/:id/like` - 点赞文章

### 音乐接口
- `GET /api/music` - 获取音乐列表
- `GET /api/music/:id` - 获取音乐详情
- `POST /api/admin/music` - 添加音乐（需要管理员权限）
- `PUT /api/admin/music/:id` - 更新音乐（需要管理员权限）
- `DELETE /api/admin/music/:id` - 删除音乐（需要管理员权限）
- `GET /api/music/playlists` - 获取播放列表
- `GET /api/music/playlists/:id` - 获取播放列表详情

### 评论接口
- `GET /api/comments` - 获取评论列表
- `POST /api/comments` - 创建评论
- `PUT /api/comments/:id/status` - 更新评论状态（需要认证）
- `DELETE /api/comments/:id` - 删除评论（需要认证）

## 数据库模型

- User - 用户
- Article - 文章
- Category - 分类
- Tag - 标签
- Comment - 评论
- Music - 音乐
- Playlist - 播放列表
- Link - 友情链接
- SiteConfig - 站点配置

## 开发计划

- [ ] 文件上传功能
- [ ] 图片管理
- [ ] 友情链接管理
- [ ] RSS订阅
- [ ] 站点地图
- [ ] 文章统计
- [ ] 邮件通知
- [ ] 更多音乐播放器功能（歌词显示、播放模式等）

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

