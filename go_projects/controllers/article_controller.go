package controllers

import (
	"blog-system/models"
	"blog-system/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	service services.ArticleService
}

func NewArticleController(service services.ArticleService) *ArticleController {
	return &ArticleController{service: service}
}

// GetArticles 获取文章列表
func (ac *ArticleController) GetArticles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	filters := make(map[string]interface{})
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	} else {
		filters["status"] = "published"
	}
	if category := c.Query("category"); category != "" {
		filters["category_id"] = category
	}
	if tag := c.Query("tag"); tag != "" {
		filters["tag_slug"] = tag
	}
	if search := c.Query("search"); search != "" {
		filters["search"] = search
	}

	articles, total, page, pageSize, err := ac.service.GetArticles(page, pageSize, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch articles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"articles":  articles,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// GetArticle 获取文章详情
func (ac *ArticleController) GetArticle(c *gin.Context) {
	id := c.Param("id")
	article, err := ac.service.GetArticle(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// 增加阅读量
	ac.service.IncrementViews(id)

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

	article := &models.Article{
		Title:      input.Title,
		Content:    input.Content,
		Excerpt:    input.Excerpt,
		CoverImage: input.CoverImage,
		AuthorID:   userID.(uint),
		CategoryID: input.CategoryID,
		Status:     input.Status,
		IsTop:      input.IsTop,
	}

	createdArticle, err := ac.service.CreateArticle(article, input.TagIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}

	// Re-fetch to get associations
	// The service CreateArticle returns the created object, but associations might need reloading if not handled in service return.
	// For now, let's assume service handles it or we fetch it.
	// Actually service CreateArticle returns *models.Article.
	// We might want to fetch it fully populated.
	fullArticle, _ := ac.service.GetArticle(strconv.Itoa(int(createdArticle.ID)))

	c.JSON(http.StatusCreated, fullArticle)
}

// UpdateArticle 更新文章
func (ac *ArticleController) UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	// Check permission
	article, err := ac.service.GetArticle(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

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

	updateData := &models.Article{
		Title:      input.Title,
		Content:    input.Content,
		Excerpt:    input.Excerpt,
		CoverImage: input.CoverImage,
		CategoryID: input.CategoryID,
		Status:     input.Status,
		IsTop:      input.IsTop,
	}

	updatedArticle, err := ac.service.UpdateArticle(id, updateData, input.TagIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update article"})
		return
	}

	// Fetch fully populated
	fullArticle, _ := ac.service.GetArticle(strconv.Itoa(int(updatedArticle.ID)))

	c.JSON(http.StatusOK, fullArticle)
}

// DeleteArticle 删除文章
func (ac *ArticleController) DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	article, err := ac.service.GetArticle(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	if role != "admin" && article.AuthorID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	if err := ac.service.DeleteArticle(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}

// LikeArticle 点赞文章
func (ac *ArticleController) LikeArticle(c *gin.Context) {
	id := c.Param("id")
	likes, err := ac.service.LikeArticle(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"likes": likes})
}
