package controllers

import (
	"blog-system/database"
	"blog-system/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LabController struct{}

func NewLabController() *LabController {
	return &LabController{}
}

type labResponse struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	Slug        string                `json:"slug"`
	Subtitle    string                `json:"subtitle"`
	Badge       string                `json:"badge"`
	BadgeColor  string                `json:"badge_color"`
	Description string                `json:"description"`
	Focus       string                `json:"focus"`
	HeroImage   string                `json:"hero_image"`
	Highlights  []models.LabHighlight `json:"highlights"`
	Resources   []models.LabResource  `json:"resources,omitempty"`
	Content     string                `json:"content,omitempty"`
}

// GetLabs 返回实验室模块列表
func (lc *LabController) GetLabs(c *gin.Context) {
	var labs []models.Lab
	if err := database.DB.Order("id ASC").Find(&labs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load labs"})
		return
	}

	results := make([]labResponse, 0, len(labs))
	for _, lab := range labs {
		results = append(results, mapLabToResponse(lab, false))
	}

	c.JSON(http.StatusOK, results)
}

// GetLab 返回单个模块详情
func (lc *LabController) GetLab(c *gin.Context) {
	slug := c.Param("slug")
	var lab models.Lab
	if err := database.DB.Where("slug = ?", slug).First(&lab).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lab not found"})
		return
	}

	c.JSON(http.StatusOK, mapLabToResponse(lab, true))
}

func mapLabToResponse(lab models.Lab, includeContent bool) labResponse {
	resp := labResponse{
		ID:          lab.ID,
		Title:       lab.Title,
		Slug:        lab.Slug,
		Subtitle:    lab.Subtitle,
		Badge:       lab.Badge,
		BadgeColor:  lab.BadgeColor,
		Description: lab.Description,
		Focus:       lab.Focus,
		HeroImage:   lab.HeroImage,
	}

	if len(lab.Highlights) > 0 {
		var highlights []models.LabHighlight
		if err := json.Unmarshal(lab.Highlights, &highlights); err == nil {
			resp.Highlights = highlights
		}
	}

	if includeContent {
		resp.Content = lab.Content
		if len(lab.ResourceLinks) > 0 {
			var resources []models.LabResource
			if err := json.Unmarshal(lab.ResourceLinks, &resources); err == nil {
				resp.Resources = resources
			}
		}
	}

	return resp
}

// GetLabArticles returns articles related to a lab by tag/topic
func (lc *LabController) GetLabArticles(c *gin.Context) {
	slug := c.Param("slug")
	var lab models.Lab
	if err := database.DB.Where("slug = ?", slug).First(&lab).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lab not found"})
		return
	}

	tagSlug := c.Query("tag")
	if tagSlug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tag query parameter required"})
		return
	}

	var articles []models.Article
	query := database.DB.Preload("Author").Preload("Category").Preload("Tags").
		Joins("JOIN article_tags ON article_tags.article_id = articles.id").
		Joins("JOIN tags ON tags.id = article_tags.tag_id").
		Where("tags.slug = ?", tagSlug).
		Where("articles.status = ?", "published").
		Order("articles.published_at DESC, articles.created_at DESC").
		Limit(20)

	if err := query.Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load articles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"lab":      lab.Slug,
		"tag":      tagSlug,
		"articles": articles,
	})
}
