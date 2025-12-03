package controllers

import (
	"blog-system/models"
	"blog-system/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) *CommentController {
	return &CommentController{service: service}
}

// GetComments 获取评论列表
func (cc *CommentController) GetComments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	comments, total, err := cc.service.GetComments(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
		"total":    total,
		"page":     page,
		"page_size": pageSize,
	})
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

	comment := &models.Comment{
		ArticleID: input.ArticleID,
		Content:   input.Content,
		Author:    input.Author,
		Email:     input.Email,
		Website:   input.Website,
		ParentID:  input.ParentID,
		IP:        c.ClientIP(),
		Status:    "pending",
	}

	createdComment, err := cc.service.CreateComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, createdComment)
}

// UpdateCommentStatus 更新评论状态
func (cc *CommentController) UpdateCommentStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := cc.service.UpdateCommentStatus(uint(id), input.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// DeleteComment 删除评论
func (cc *CommentController) DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := cc.service.DeleteComment(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

// GetPendingComments 获取待审核评论
func (cc *CommentController) GetPendingComments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	comments, total, err := cc.service.GetPendingComments(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pending comments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
		"total":    total,
		"page":     page,
		"page_size": pageSize,
	})
}
