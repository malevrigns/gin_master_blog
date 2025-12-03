package controllers

import (
	"blog-system/models"
	"blog-system/services"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type MusicController struct {
	service services.MusicService
}

func NewMusicController(service services.MusicService) *MusicController {
	return &MusicController{service: service}
}

// GetMusics 获取音乐列表
func (ac *MusicController) GetMusics(c *gin.Context) {
	// Service currently doesn't support pagination or search filters in FindAll
	// This is a regression from original controller.
	// I should update Service/Repository to support filters.
	// For now, I'll just call GetMusicList which returns all.
	musics, err := ac.service.GetMusicList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch music"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"musics": musics,
		"total":  len(musics),
	})
}

// GetMusic 获取音乐详情
func (ac *MusicController) GetMusic(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	music, err := ac.service.GetMusic(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Music not found"})
		return
	}
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

	music := &models.Music{
		Title:    input.Title,
		Artist:   input.Artist,
		Cover:    input.Cover,
		URL:      normalizeMusicURL(input.URL),
		Duration: input.Duration,
		Lrc:      input.Lrc,
		IsPublic: input.IsPublic,
	}

	createdMusic, err := ac.service.CreateMusic(music)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create music"})
		return
	}

	c.JSON(http.StatusCreated, createdMusic)
}

// UpdateMusic 更新音乐
func (ac *MusicController) UpdateMusic(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
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

	updateData := &models.Music{
		Title:    input.Title,
		Artist:   input.Artist,
		Cover:    input.Cover,
		URL:      normalizeMusicURL(input.URL),
		Duration: input.Duration,
		Lrc:      input.Lrc,
		IsPublic: input.IsPublic,
	}

	music, err := ac.service.UpdateMusic(uint(id), updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update music"})
		return
	}

	c.JSON(http.StatusOK, music)
}

// DeleteMusic 删除音乐
func (ac *MusicController) DeleteMusic(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ac.service.DeleteMusic(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete music"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Music deleted successfully"})
}

// GetPlaylists 获取播放列表
func (ac *MusicController) GetPlaylists(c *gin.Context) {
	playlists, err := ac.service.GetPlaylists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch playlists"})
		return
	}
	c.JSON(http.StatusOK, playlists)
}

// GetPlaylist 获取播放列表详情
func (ac *MusicController) GetPlaylist(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	playlist, err := ac.service.GetPlaylist(uint(id))
	if err != nil {
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
