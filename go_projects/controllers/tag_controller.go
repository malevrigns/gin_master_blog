package controllers

import (
	"blog-system/models"
	"blog-system/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	service services.TagService
}

func NewTagController(service services.TagService) *TagController {
	return &TagController{service: service}
}

func (tc *TagController) GetTags(c *gin.Context) {
	tags, err := tc.service.GetTags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tags"})
		return
	}
	c.JSON(http.StatusOK, tags)
}

func (tc *TagController) CreateTag(c *gin.Context) {
	var input models.Tag
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag, err := tc.service.CreateTag(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tag"})
		return
	}

	c.JSON(http.StatusCreated, tag)
}

func (tc *TagController) DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := tc.service.DeleteTag(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
}
