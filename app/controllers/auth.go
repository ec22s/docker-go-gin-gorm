package controllers

import (
  "go_gin_gorm/models"
  "net/http"

  "github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
  var input RegisterInput

  // リクエストのJSONデータをRegisterInput構造体にバインド
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ユーザーオブジェクトを作成し、データベースに保存
	user := &models.User{Username: input.Username, Password: input.Password}
	user, err := user.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

  // 成功した場合、ユーザー情報をレスポンスとして返す
  c.JSON(http.StatusOK, gin.H{
		"data": user.PrepareOutput(),
	})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := models.GenerateToken(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
