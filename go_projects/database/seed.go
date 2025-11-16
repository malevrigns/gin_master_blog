package database

import (
	"blog-system/models"
	"blog-system/utils"
	"encoding/json"
	"log"
	"time"
)

func SeedDatabase() {
	var userCount int64
	DB.Model(&models.User{}).Count(&userCount)

	var admin models.User
	var user models.User

	if userCount > 0 {
		log.Println("Database already seeded, skipping core seed...")

		if err := DB.Where("username = ?", "admin").First(&admin).Error; err != nil {
			log.Printf("Error finding admin user: %v", err)
			return
		}
		if err := DB.Where("username = ?", "user").First(&user).Error; err != nil {
			log.Printf("Error finding default user: %v", err)
			return
		}
	} else {
		log.Println("Seeding database...")

		adminPassword, err := utils.HashPassword("admin123")
		if err != nil {
			log.Printf("Error hashing admin password: %v", err)
			return
		}
		admin = models.User{
			Username: "admin",
			Email:    "admin@example.com",
			Password: adminPassword,
			Role:     "admin",
			Bio:      "系统管理员，负责站点配置。",
			Avatar:   "",
		}
		if err := DB.Create(&admin).Error; err != nil {
			log.Printf("Error creating admin user: %v", err)
			return
		}
		log.Println("Created admin user: admin / admin123")

		userPassword, err := utils.HashPassword("user123")
		if err != nil {
			log.Printf("Error hashing user password: %v", err)
			return
		}
		user = models.User{
			Username: "user",
			Email:    "user@example.com",
			Password: userPassword,
			Role:     "user",
			Bio:      "普通用户，用于演示评论功能。",
			Avatar:   "",
		}
		if err := DB.Create(&user).Error; err != nil {
			log.Printf("Error creating user: %v", err)
			return
		}
		log.Println("Created user: user / user123")

		categories := []models.Category{
			{Name: "技术分享", Slug: "tech", Description: "关于后端、前端与工程实践的记录"},
			{Name: "生活随笔", Slug: "life", Description: "生活感悟与日常记录"},
			{Name: "学习笔记", Slug: "study", Description: "学习过程中的总结与思考"},
			{Name: "项目经验", Slug: "project", Description: "项目复盘与实践经验"},
		}
		for _, category := range categories {
			if err := DB.Create(&category).Error; err != nil {
				log.Printf("Error creating category %s: %v", category.Name, err)
			}
		}
		log.Println("Created categories")

		tags := []models.Tag{
			{Name: "Go语言", Slug: "go"},
			{Name: "Vue3", Slug: "vue3"},
			{Name: "前端开发", Slug: "frontend"},
			{Name: "后端开发", Slug: "backend"},
			{Name: "数据库", Slug: "database"},
			{Name: "算法", Slug: "algorithm"},
			{Name: "设计模式", Slug: "design-pattern"},
			{Name: "Go / Gin 实战", Slug: "go-gin"},
			{Name: "微服务架构", Slug: "microservice"},
			{Name: "数据库调优", Slug: "database-tuning"},
			{Name: "Vue3 进阶", Slug: "vue3-advanced"},
			{Name: "动效实验", Slug: "motion-lab"},
			{Name: "可视化探索", Slug: "visualization"},
			{Name: "阅读记录", Slug: "reading-notes"},
			{Name: "灵感清单", Slug: "inspiration-list"},
			{Name: "生活方式", Slug: "life-style"},
		}
		for _, tag := range tags {
			if err := DB.Create(&tag).Error; err != nil {
				log.Printf("Error creating tag %s: %v", tag.Name, err)
			}
		}
		log.Println("Created tags")
	}

	var techCategory, lifeCategory models.Category
	var goTag, vueTag, frontendTag models.Tag

	if err := DB.Where("slug = ?", "tech").First(&techCategory).Error; err != nil {
		log.Printf("Error finding tech category: %v", err)
		return
	}
	if err := DB.Where("slug = ?", "life").First(&lifeCategory).Error; err != nil {
		log.Printf("Error finding life category: %v", err)
		return
	}
	if err := DB.Where("slug = ?", "go").First(&goTag).Error; err != nil {
		log.Printf("Error finding go tag: %v", err)
		return
	}
	if err := DB.Where("slug = ?", "vue3").First(&vueTag).Error; err != nil {
		log.Printf("Error finding vue3 tag: %v", err)
		return
	}
	if err := DB.Where("slug = ?", "frontend").First(&frontendTag).Error; err != nil {
		log.Printf("Error finding frontend tag: %v", err)
		return
	}

	now := time.Now()
	articles := []models.Article{
		{
			Title: "欢迎来到我的博客",
			Slug:  "welcome-to-my-blog",
			Content: `# 欢迎来到我的博客

这是一个由 **Go + Gin + Vue3** 打造的现代化博客系统示例，涵盖后台管理、内容展示和实验室模块。

## 功能亮点

- 文章管理、分类与标签
- 评论与音乐播放
- 友链与独立实验室
- 深色模式与响应式设计

希望这些例子能够帮助你快速搭建自己的博客。`,
			Excerpt:     "这是一个由 Go + Gin + Vue3 打造的现代化博客系统示例。",
			CoverImage:  "https://via.placeholder.com/800x400",
			AuthorID:    admin.ID,
			CategoryID:  techCategory.ID,
			Status:      "published",
			IsTop:       true,
			PublishedAt: &now,
		},
		{
			Title: "Go语言快速入门指南",
			Slug:  "go-language-tutorial",
			Content: `# Go语言快速入门指南

Go 是 Google 推出的编程语言，语法简洁、并发模型优秀，十分适合构建云原生服务。

## 为什么选择 Go

1. 语法简洁，容易上手
2. goroutine + channel 构建轻量并发
3. 内置工具链，交叉编译简单

Now let's start your Go journey!`,
			Excerpt:     "Introduction to Go language features and quick start examples.",
			CoverImage:  "https://via.placeholder.com/800x400",
			AuthorID:    admin.ID,
			CategoryID:  techCategory.ID,
			Status:      "published",
			PublishedAt: &now,
		},
		{
			Title: "Vue3 Composition API Practice",
			Slug:  "vue3-composition-api",
			Content: `# Vue3 Composition API Practice

Composition API makes logic reuse and organization clearer, suitable for building complex components.

## Basic Example

By properly splitting hooks, the code can be kept cleaner.`,
			Excerpt:     "Advantages and basic examples of Composition API.",
			CoverImage:  "https://via.placeholder.com/800x400",
			AuthorID:    admin.ID,
			CategoryID:  techCategory.ID,
			Status:      "published",
			PublishedAt: &now,
		},
		{
			Title: "Life Essay: Recording Every Day",
			Slug:  "life-notes",
			Content: `# Life Essay: Recording Every Day

Life is like a long journey. Recording daily inspirations and moments allows us to better dialogue with ourselves.

## Why Keep Records

- Preserve gentle moments
- Review growth
- Share with more friends

From today, start writing your life story.`,
			Excerpt:     "The meaning and methods of recording life.",
			CoverImage:  "https://via.placeholder.com/800x400",
			AuthorID:    user.ID,
			CategoryID:  lifeCategory.ID,
			Status:      "published",
			PublishedAt: &now,
		},
	}

	for i := range articles {
		article := &articles[i]
		if err := DB.Create(article).Error; err != nil {
			log.Printf("Error creating article %s: %v", article.Title, err)
			continue
		}

		switch article.Slug {
		case "welcome-to-my-blog":
			if err := DB.Model(article).Association("Tags").Append(&goTag, &vueTag); err != nil {
				log.Printf("Error associating tags to article %s: %v", article.Title, err)
			}
		case "go-language-tutorial":
			if err := DB.Model(article).Association("Tags").Append(&goTag); err != nil {
				log.Printf("Error associating tags to article %s: %v", article.Title, err)
			}
		case "vue3-composition-api":
			if err := DB.Model(article).Association("Tags").Append(&vueTag, &frontendTag); err != nil {
				log.Printf("Error associating tags to article %s: %v", article.Title, err)
			}
		}
	}
	log.Println("Created sample articles")

	musics := []models.Music{
		{
			Title:    "LoFi Study Session",
			Artist:   "Demo Artist",
			Cover:    "https://via.placeholder.com/300x300",
			URL:      "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3",
			Duration: 180,
			IsPublic: true,
		},
		{
			Title:    "Night Coding",
			Artist:   "Demo Artist",
			Cover:    "https://via.placeholder.com/300x300",
			URL:      "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-2.mp3",
			Duration: 200,
			IsPublic: true,
		},
	}
	for _, music := range musics {
		if err := DB.Create(&music).Error; err != nil {
			log.Printf("Error creating music %s: %v", music.Title, err)
		}
	}
	log.Println("Created sample musics")

	links := []models.Link{
		{
			Name:      "GitHub",
			URL:       "https://github.com",
			Logo:      "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png",
			Desc:      "代码托管与协作平台。",
			IsVisible: true,
			Sort:      1,
		},
		{
			Name:      "Vue.js",
			URL:       "https://vuejs.org",
			Logo:      "https://vuejs.org/logo.svg",
			Desc:      "渐进式 JavaScript 框架。",
			IsVisible: true,
			Sort:      2,
		},
	}
	for _, link := range links {
		if err := DB.Create(&link).Error; err != nil {
			log.Printf("Error creating link %s: %v", link.Name, err)
		}
	}
	log.Println("Created sample links")

	log.Println("Database seeding completed!")

	seedLabs()
}

func seedLabs() {
	var count int64
	DB.Model(&models.Lab{}).Count(&count)
	if count > 0 {
		log.Println("Labs already seeded, skipping...")
		return
	}

	labs := []models.Lab{
		{
			Title:       "后端实验室",
			Slug:        "backend-lab",
			Subtitle:    "记录 Go / Gin / 数据库调优等后端实践",
			Badge:       "LAB",
			BadgeColor:  "#34d399",
			Description: "聚焦后端稳定性与性能，沉淀工程经验与调优技巧。",
			Focus:       "关注稳定性与性能",
			HeroImage:   "https://images.unsplash.com/photo-1555949963-ff9fe0c870eb",
			Content: `## 后端实验室 · 研究方向

- Go & Gin 最佳实践
- 微服务与高并发场景
- 数据库调优、缓存与性能监控

在这里，持续打磨后端体系，让服务更稳定可靠。`,
			Highlights: toJSON([]models.LabHighlight{
				{Title: "Go / Gin 实战", Description: "生产级项目实战记录", Tag: "go-gin"},
				{Title: "微服务架构", Description: "服务拆分与稳定性方案", Tag: "microservice"},
				{Title: "数据库调优", Description: "性能监控与调优手册", Tag: "database-tuning"},
			}),
			ResourceLinks: toJSON([]models.LabResource{
				{Title: "项目模板", Desc: "基于 Go/Gin 的脚手架", URL: "https://github.com", Icon: "TrendCharts"},
				{Title: "API 规范", Desc: "团队约定式接口规范", URL: "https://example.com/api-spec", Icon: "Document"},
			}),
		},
		{
			Title:       "前端游乐场",
			Slug:        "frontend-playground",
			Subtitle:    "尝试 Vue3、动画、可视化的互动实验。",
			Badge:       "PLAYGROUND",
			BadgeColor:  "#38bdf8",
			Description: "探索交互体验与表现力，构建富有生命力的界面。",
			Focus:       "探索体验与表现力",
			HeroImage:   "https://images.unsplash.com/photo-1500530855697-b586d89ba3ee",
			Content: `## 前端游乐场 · 探索主题

- Vue 3 + Vite 高阶玩法
- WebGL / Canvas 动效
- 数据可视化与交互体验

让灵感驱动每一次尝试。`,
			Highlights: toJSON([]models.LabHighlight{
				{Title: "Vue3 进阶", Description: "组合式 + 工程化的最佳实践", Tag: "vue3-advanced"},
				{Title: "动效实验", Description: "动效和微交互的探索", Tag: "motion-lab"},
				{Title: "可视化探索", Description: "Canvas / WebGL 可视化", Tag: "visualization"},
			}),
			ResourceLinks: toJSON([]models.LabResource{
				{Title: "组件库", Desc: "常用组件与动画封装", URL: "https://example.com/components", Icon: "Operation"},
				{Title: "Playground", Desc: "互动实验合集", URL: "https://stackblitz.com", Icon: "Monitor"},
			}),
		},
		{
			Title:       "生活与随想",
			Slug:        "life-notes",
			Subtitle:    "记录生活、阅读和思考片段，让博客更立体。",
			Badge:       "NOTES",
			BadgeColor:  "#c084fc",
			Description: "技术之外，保留灵感、热爱与人文思考。",
			Focus:       "技术之外的灵感",
			HeroImage:   "https://images.unsplash.com/photo-1476610182048-b716b8518aae",
			Content: `## 生活与随想 · 记事簿

- 阅读/观影笔记
- 旅行纪实与生活单曲循环
- 灵感与思考片段

让故事和温度留在这里。`,
			Highlights: toJSON([]models.LabHighlight{
				{Title: "阅读记录", Description: "书影音摘录与感想", Tag: "reading-notes"},
				{Title: "灵感清单", Description: "灵感与素材收藏", Tag: "inspiration-list"},
				{Title: "生活方式", Description: "生活方式与效率实验", Tag: "life-style"},
			}),
			ResourceLinks: toJSON([]models.LabResource{
				{Title: "灵感清单", Desc: "我收藏的灵感来源", URL: "https://example.com/inspiration", Icon: "StarFilled"},
				{Title: "随想播放列表", Desc: "适合深夜的曲目", URL: "https://music.163.com", Icon: "Headset"},
			}),
		},
	}

	for _, lab := range labs {
		if err := DB.Create(&lab).Error; err != nil {
			log.Printf("Error creating lab %s: %v", lab.Title, err)
		}
	}
	log.Println("Seeded lab modules")
}

func toJSON(v interface{}) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		log.Printf("Failed to marshal seed data: %v", err)
		return []byte("[]")
	}
	return data
}
