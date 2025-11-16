package controllers

import (
	"blog-system/database"
	"blog-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type TagController struct{}

func NewTagController() *TagController {
	return &TagController{}
}

// GetTags 获取标签列表
func (tc *TagController) GetTags(c *gin.Context) {
	var tags []models.Tag
	database.DB.Find(&tags)
	c.JSON(http.StatusOK, tags)
}

// CreateTag 创建标签
func (tc *TagController) CreateTag(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查标签是否已存在
	var existingTag models.Tag
	if database.DB.Where("slug = ?", slug.Make(input.Name)).First(&existingTag).Error == nil {
		c.JSON(http.StatusOK, existingTag)
		return
	}

	tag := models.Tag{
		Name: input.Name,
		Slug: slug.Make(input.Name),
	}

	if err := database.DB.Create(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tag"})
		return
	}

	c.JSON(http.StatusCreated, tag)
}

// DeleteTag 删除标签
func (tc *TagController) DeleteTag(c *gin.Context) {
	id := c.Param("id")
	var tag models.Tag

	if err := database.DB.First(&tag, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	database.DB.Delete(&tag)
	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
}

