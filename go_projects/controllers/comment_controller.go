package controllers

import (
	"blog-system/database"
	"blog-system/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct{}

func NewCommentController() *CommentController {
	return &CommentController{}
}

// GetComments 获取评论列表
func (cc *CommentController) GetComments(c *gin.Context) {
	articleID := c.Query("article_id")
	var comments []models.Comment

	query := database.DB.Preload("Replies")

	if articleID != "" {
		query = query.Where("article_id = ? AND parent_id IS NULL", articleID)
	}

	query.Where("status = ?", "approved").Order("created_at DESC").Find(&comments)

	c.JSON(http.StatusOK, comments)
}

// CreateComment 创建评论
func (cc *CommentController) CreateComment(c *gin.Context) {
	var input struct {
		ArticleID uint   `json:"article_id" binding:"required"`
		Content   string `json:"content" binding:"required"`
		Author    string `json:"author" binding:"required"`
		Email     string `json:"email"`
		Website   string `json:"website"`
		ParentID  *uint  `json:"parent_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := models.Comment{
		ArticleID: input.ArticleID,
		Content:   input.Content,
		Author:    input.Author,
		Email:     input.Email,
		Website:   input.Website,
		ParentID:  input.ParentID,
		IP:        c.ClientIP(),
		Status:    "pending",
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// UpdateCommentStatus 更新评论状态
func (cc *CommentController) UpdateCommentStatus(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment

	if err := database.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	var input struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.Status = input.Status
	if err := database.DB.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// DeleteComment 删除评论
func (cc *CommentController) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment

	if err := database.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	database.DB.Delete(&comment)
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

// GetPendingComments 获取待审核评论
func (cc *CommentController) GetPendingComments(c *gin.Context) {
	var comments []models.Comment

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	var total int64
	database.DB.Model(&models.Comment{}).Where("status = ?", "pending").Count(&total)

	database.DB.Preload("Article").Where("status = ?", "pending").
		Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&comments)

	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
		"total":    total,
		"page":     page,
		"page_size": pageSize,
	})
}

