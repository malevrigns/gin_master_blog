package controllers

import (
	"blog-system/models"
	"blog-system/services"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LabController struct {
	labService     services.LabService
	articleService services.ArticleService
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
	Highlights  []models.LabHighlight `json:"highlights,omitempty"`
	Content     string                `json:"content,omitempty"`
	Resources   []models.LabResource  `json:"resources,omitempty"`
}

func NewLabController(labService services.LabService, articleService services.ArticleService) *LabController {
	return &LabController{
		labService:     labService,
		articleService: articleService,
	}
}

// ... (previous methods GetLabs, GetLab, mapLabToResponse remain same)

// GetLabs 返回实验室模块列表
func (lc *LabController) GetLabs(c *gin.Context) {
	labs, err := lc.labService.GetLabs()
	if err != nil {
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
	lab, err := lc.labService.GetLabBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lab not found"})
		return
	}

	c.JSON(http.StatusOK, mapLabToResponse(*lab, true))
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
	lab, err := lc.labService.GetLabBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lab not found"})
		return
	}

	tagSlug := c.Query("tag")
	if tagSlug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tag query parameter required"})
		return
	}

	filters := map[string]interface{}{
		"tag_slug": tagSlug,
		"status":   "published",
	}

	articles, _, _, _, err := lc.articleService.GetArticles(1, 20, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load articles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"lab":      lab.Slug,
		"tag":      tagSlug,
		"articles": articles,
	})
}
