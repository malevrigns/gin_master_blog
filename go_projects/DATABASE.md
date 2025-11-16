# 数据库说明

## 数据库类型

本项目默认使用 **MySQL** 数据库，也支持 SQLite（用于开发测试）。

- **MySQL**（默认）- 适合生产环境和开发
- **SQLite** - 适合快速测试

## 数据库初始化

### 自动初始化

**重要：使用 MySQL 前需要先创建数据库！**

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

3. **运行程序**：
首次运行程序时，系统会自动：
   - 连接 MySQL 数据库
   - 自动创建所有数据表
   - 填充初始示例数据

### 初始数据

系统会自动创建以下初始数据：

#### 用户账户
- **管理员账户**
  - 用户名: `admin`
  - 密码: `admin123`
  - 角色: admin

- **普通用户账户**
  - 用户名: `user`
  - 密码: `user123`
  - 角色: user

⚠️ **重要提示**: 首次登录后请立即修改默认密码！

#### 示例数据
- 4 个分类（技术分享、生活随笔、学习笔记、项目经验）
- 7 个标签（Go语言、Vue3、前端开发、后端开发、数据库、算法、设计模式）
- 4 篇示例文章
- 2 首示例音乐
- 2 个友情链接

## 数据库配置

### MySQL（默认）

1. **创建数据库**：
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

3. **运行程序**：
程序会自动连接数据库并创建表结构。

### SQLite（可选）

如果需要使用 SQLite，在 `config/config.yaml` 文件中设置：

```yaml
database:
  type: sqlite
  name: blog.db
```

数据库文件会自动创建在项目根目录。

## 数据库文件位置

### MySQL
- 数据库名称: `blog_system`（可在 `.env` 中配置）
- 需要先手动创建数据库（见上方配置说明）

### SQLite
- 文件路径: `go_projects/blog.db`
- 注意: 此文件已在 `.gitignore` 中，不会被提交到版本控制

## 重置数据库

如果需要重置数据库：

1. **SQLite**: 删除 `blog.db` 文件，重新运行程序
2. **MySQL**: 删除数据库并重新创建，或使用以下 SQL：
```sql
DROP DATABASE blog_system;
CREATE DATABASE blog_system CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

## 数据库模型

主要数据表：
- `users` - 用户表
- `articles` - 文章表
- `categories` - 分类表
- `tags` - 标签表
- `article_tags` - 文章标签关联表
- `comments` - 评论表
- `musics` - 音乐表
- `playlists` - 播放列表表
- `playlist_musics` - 播放列表音乐关联表
- `links` - 友情链接表
- `site_configs` - 站点配置表

## 备份数据库

### SQLite
```bash
cp blog.db blog.db.backup
```

### MySQL
```bash
mysqldump -u root -p blog_system > blog_system_backup.sql
```

## 注意事项

1. 首次运行会自动创建数据库和初始数据
2. 如果数据库已存在且有数据，不会重复填充初始数据
3. 生产环境请务必修改默认密码
4. 建议定期备份数据库

