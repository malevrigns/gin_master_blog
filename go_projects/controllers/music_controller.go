package controllers

import (
	"blog-system/database"
	"blog-system/models"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type MusicController struct{}

func NewMusicController() *MusicController {
	return &MusicController{}
}

// GetMusics 获取音乐列表
func (ac *MusicController) GetMusics(c *gin.Context) {
	var musics []models.Music
	query := database.DB

	// 只获取公开的音乐（非管理员）
	_, isAdmin := c.Get("role")
	if !isAdmin {
		query = query.Where("is_public = ?", true)
	}

	// 搜索
	search := c.Query("search")
	if search != "" {
		query = query.Where("title LIKE ? OR artist LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// 分页
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	var total int64
	query.Model(&models.Music{}).Count(&total)

	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&musics)

	c.JSON(http.StatusOK, gin.H{
		"musics":   musics,
		"total":    total,
		"page":     page,
		"page_size": pageSize,
	})
}

// GetMusic 获取音乐详情
func (ac *MusicController) GetMusic(c *gin.Context) {
	id := c.Param("id")
	var music models.Music

	if err := database.DB.First(&music, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Music not found"})
		return
	}

	// 增加播放次数
	database.DB.Model(&music).Update("play_count", music.PlayCount+1)

	c.JSON(http.StatusOK, music)
}

// CreateMusic 创建音乐
func (ac *MusicController) CreateMusic(c *gin.Context) {
	var input struct {
		Title    string `json:"title" binding:"required"`
		Artist   string `json:"artist"`
		Cover    string `json:"cover"`
		URL      string `json:"url" binding:"required"`
		Duration int    `json:"duration"`
		Lrc      string `json:"lrc"`
		IsPublic bool   `json:"is_public"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	music := models.Music{
		Title:    input.Title,
		Artist:   input.Artist,
		Cover:    input.Cover,
		URL:      normalizeMusicURL(input.URL),
		Duration: input.Duration,
		Lrc:      input.Lrc,
		IsPublic: input.IsPublic,
	}

	if err := database.DB.Create(&music).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create music"})
		return
	}

	c.JSON(http.StatusCreated, music)
}

// UpdateMusic 更新音乐
func (ac *MusicController) UpdateMusic(c *gin.Context) {
	id := c.Param("id")
	var music models.Music

	if err := database.DB.First(&music, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Music not found"})
		return
	}

	var input struct {
		Title    string `json:"title"`
		Artist   string `json:"artist"`
		Cover    string `json:"cover"`
		URL      string `json:"url"`
		Duration int    `json:"duration"`
		Lrc      string `json:"lrc"`
		IsPublic bool   `json:"is_public"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Title != "" {
		music.Title = input.Title
	}
	if input.Artist != "" {
		music.Artist = input.Artist
	}
	if input.Cover != "" {
		music.Cover = input.Cover
	}
	if input.URL != "" {
		music.URL = normalizeMusicURL(input.URL)
	}
	if input.Duration != 0 {
		music.Duration = input.Duration
	}
	if input.Lrc != "" {
		music.Lrc = input.Lrc
	}
	music.IsPublic = input.IsPublic

	if err := database.DB.Save(&music).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update music"})
		return
	}

	c.JSON(http.StatusOK, music)
}

// DeleteMusic 删除音乐
func (ac *MusicController) DeleteMusic(c *gin.Context) {
	id := c.Param("id")
	var music models.Music

	if err := database.DB.First(&music, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Music not found"})
		return
	}

	database.DB.Delete(&music)
	c.JSON(http.StatusOK, gin.H{"message": "Music deleted successfully"})
}

// GetPlaylists 获取播放列表
func (ac *MusicController) GetPlaylists(c *gin.Context) {
	var playlists []models.Playlist
	query := database.DB.Preload("Musics")

	// 只获取公开的播放列表
	_, isAdmin := c.Get("role")
	if !isAdmin {
		query = query.Where("is_public = ?", true)
	}

	query.Order("created_at DESC").Find(&playlists)

	c.JSON(http.StatusOK, playlists)
}

// GetPlaylist 获取播放列表详情
func (ac *MusicController) GetPlaylist(c *gin.Context) {
	id := c.Param("id")
	var playlist models.Playlist

	if err := database.DB.Preload("Musics").First(&playlist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Playlist not found"})
		return
	}

	c.JSON(http.StatusOK, playlist)
}

func normalizeMusicURL(raw string) string {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return trimmed
	}

	if strings.Contains(trimmed, "music.163.com") {
		if strings.Contains(trimmed, "/song/media/outer/url") {
			return trimmed
		}

		if u, err := url.Parse(trimmed); err == nil {
			q := u.Query()
			id := q.Get("id")
			if id != "" {
				id = strings.TrimSuffix(id, ".mp3")
				return buildNeteaseOuterURL(id)
			}
		}
	}

	if isDigits(trimmed) {
		return buildNeteaseOuterURL(trimmed)
	}

	return trimmed
}

func isDigits(s string) bool {
	if s == "" {
		return false
	}
	match, _ := regexp.MatchString(`^\d+$`, s)
	return match
}

func buildNeteaseOuterURL(id string) string {
	return fmt.Sprintf("https://music.163.com/song/media/outer/url?id=%s.mp3", id)
}

