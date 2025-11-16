package controllers

import (
	"blog-system/database"
	"blog-system/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type ArticleController struct{}

func NewArticleController() *ArticleController {
	return &ArticleController{}
}

// GetArticles 获取文章列表
func (ac *ArticleController) GetArticles(c *gin.Context) {
	var articles []models.Article
	query := database.DB.Preload("Author").Preload("Category").Preload("Tags")

	// 状态过滤
	status := c.Query("status")
	if status == "" {
		status = "published"
	}
	query = query.Where("status = ?", status)

	// 分类过滤
	category := c.Query("category")
	if category != "" {
		query = query.Where("category_id = ?", category)
	}

	// 标签过滤
	tag := c.Query("tag")
	if tag != "" {
		query = query.Joins("JOIN article_tags ON article_tags.article_id = articles.id").
			Joins("JOIN tags ON tags.id = article_tags.tag_id").
			Where("tags.slug = ?", tag)
	}

	// 搜索
	search := c.Query("search")
	if search != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// 分页
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	offset := (page - 1) * pageSize

	var total int64
	query.Model(&models.Article{}).Count(&total)

	query.Order("is_top DESC, published_at DESC").Offset(offset).Limit(pageSize).Find(&articles)

	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
		"total":    total,
		"page":     page,
		"page_size": pageSize,
	})
}

// GetArticle 获取文章详情
func (ac *ArticleController) GetArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	if err := database.DB.Preload("Author").Preload("Category").Preload("Tags").Preload("Comments", "status = ?", "approved").Preload("Comments.Replies").First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// 增加阅读量
	database.DB.Model(&article).Update("views", article.Views+1)

	c.JSON(http.StatusOK, article)
}

// CreateArticle 创建文章
func (ac *ArticleController) CreateArticle(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var input struct {
		Title      string   `json:"title" binding:"required"`
		Content    string   `json:"content" binding:"required"`
		Excerpt    string   `json:"excerpt"`
		CoverImage string   `json:"cover_image"`
		CategoryID uint     `json:"category_id"`
		TagIDs     []uint   `json:"tag_ids"`
		Status     string   `json:"status"`
		IsTop      bool     `json:"is_top"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Status == "" {
		input.Status = "draft"
	}

	articleSlug := slug.Make(input.Title)
	// 确保slug唯一
	var count int64
	database.DB.Model(&models.Article{}).Where("slug = ?", articleSlug).Count(&count)
	if count > 0 {
		articleSlug = articleSlug + "-" + strconv.FormatInt(time.Now().Unix(), 10)
	}

	article := models.Article{
		Title:      input.Title,
		Slug:       articleSlug,
		Content:    input.Content,
		Excerpt:    input.Excerpt,
		CoverImage: input.CoverImage,
		AuthorID:   userID.(uint),
		CategoryID: input.CategoryID,
		Status:     input.Status,
		IsTop:      input.IsTop,
	}

	if input.Status == "published" {
		now := time.Now()
		article.PublishedAt = &now
	}

	if err := database.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}

	// 关联标签
	if len(input.TagIDs) > 0 {
		var tags []models.Tag
		database.DB.Where("id IN ?", input.TagIDs).Find(&tags)
		database.DB.Model(&article).Association("Tags").Replace(tags)
	}

	database.DB.Preload("Author").Preload("Category").Preload("Tags").First(&article, article.ID)

	c.JSON(http.StatusCreated, article)
}

// UpdateArticle 更新文章
func (ac *ArticleController) UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	var article models.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// 检查权限
	if role != "admin" && article.AuthorID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	var input struct {
		Title      string   `json:"title"`
		Content    string   `json:"content"`
		Excerpt    string   `json:"excerpt"`
		CoverImage string   `json:"cover_image"`
		CategoryID uint     `json:"category_id"`
		TagIDs     []uint   `json:"tag_ids"`
		Status     string   `json:"status"`
		IsTop      bool     `json:"is_top"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Title != "" {
		article.Title = input.Title
		article.Slug = slug.Make(input.Title)
	}
	if input.Content != "" {
		article.Content = input.Content
	}
	if input.Excerpt != "" {
		article.Excerpt = input.Excerpt
	}
	if input.CoverImage != "" {
		article.CoverImage = input.CoverImage
	}
	if input.CategoryID != 0 {
		article.CategoryID = input.CategoryID
	}
	if input.Status != "" {
		article.Status = input.Status
		if input.Status == "published" && article.PublishedAt == nil {
			now := time.Now()
			article.PublishedAt = &now
		}
	}
	article.IsTop = input.IsTop

	if err := database.DB.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update article"})
		return
	}

	// 更新标签
	if input.TagIDs != nil {
		var tags []models.Tag
		database.DB.Where("id IN ?", input.TagIDs).Find(&tags)
		database.DB.Model(&article).Association("Tags").Replace(tags)
	}

	database.DB.Preload("Author").Preload("Category").Preload("Tags").First(&article, article.ID)

	c.JSON(http.StatusOK, article)
}

// DeleteArticle 删除文章
func (ac *ArticleController) DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	var article models.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// 检查权限
	if role != "admin" && article.AuthorID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	database.DB.Delete(&article)
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}

// LikeArticle 点赞文章
func (ac *ArticleController) LikeArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	database.DB.Model(&article).Update("likes", article.Likes+1)
	c.JSON(http.StatusOK, gin.H{"likes": article.Likes + 1})
}

