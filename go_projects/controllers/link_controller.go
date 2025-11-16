package controllers

import (
	"blog-system/database"
	"blog-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LinkController struct{}

func NewLinkController() *LinkController {
	return &LinkController{}
}

// GetLinks 获取友情链接列表
func (lc *LinkController) GetLinks(c *gin.Context) {
	var links []models.Link
	query := database.DB

	// 只获取可见的链接
	_, isAdmin := c.Get("role")
	if !isAdmin {
		query = query.Where("is_visible = ?", true)
	}

	query.Order("sort ASC, created_at DESC").Find(&links)
	c.JSON(http.StatusOK, links)
}

// CreateLink 创建友情链接
func (lc *LinkController) CreateLink(c *gin.Context) {
	var input struct {
		Name      string `json:"name" binding:"required"`
		URL       string `json:"url" binding:"required"`
		Logo      string `json:"logo"`
		Desc      string `json:"desc"`
		IsVisible bool   `json:"is_visible"`
		Sort      int    `json:"sort"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link := models.Link{
		Name:      input.Name,
		URL:       input.URL,
		Logo:      input.Logo,
		Desc:      input.Desc,
		IsVisible: input.IsVisible,
		Sort:      input.Sort,
	}

	if err := database.DB.Create(&link).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create link"})
		return
	}

	c.JSON(http.StatusCreated, link)
}

// UpdateLink 更新友情链接
func (lc *LinkController) UpdateLink(c *gin.Context) {
	id := c.Param("id")
	var link models.Link

	if err := database.DB.First(&link, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	var input struct {
		Name      string `json:"name"`
		URL       string `json:"url"`
		Logo      string `json:"logo"`
		Desc      string `json:"desc"`
		IsVisible bool   `json:"is_visible"`
		Sort      int    `json:"sort"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != "" {
		link.Name = input.Name
	}
	if input.URL != "" {
		link.URL = input.URL
	}
	if input.Logo != "" {
		link.Logo = input.Logo
	}
	if input.Desc != "" {
		link.Desc = input.Desc
	}
	link.IsVisible = input.IsVisible
	link.Sort = input.Sort

	if err := database.DB.Save(&link).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update link"})
		return
	}

	c.JSON(http.StatusOK, link)
}

// DeleteLink 删除友情链接
func (lc *LinkController) DeleteLink(c *gin.Context) {
	id := c.Param("id")
	var link models.Link

	if err := database.DB.First(&link, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	database.DB.Delete(&link)
	c.JSON(http.StatusOK, gin.H{"message": "Link deleted successfully"})
}

