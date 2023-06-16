package controller

import (
	"net/http"
	"personal/TODO-list/database"
	"personal/TODO-list/helper"
	"personal/TODO-list/model"
	"personal/TODO-list/pkg/util"
	"personal/TODO-list/service"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GithubCallbackHandler(c *gin.Context) {
	code := c.Query("code")

	githubAccessToken, err := util.GetGithubAccessToken(code)
	if err != nil {
		helper.SendErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	githubUser, err := util.GetGithubUserData(githubAccessToken)
	if err != nil {
		helper.SendErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	now := time.Now()
	email := strings.ToLower(githubUser.Email)

	userData := model.User{
		Name:      githubUser.Login,
		Email:     githubUser.Email,
		Password:  "",
		CreatedAt: now,
		UpdatedAt: now,
	}

	var user model.User
	err = database.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		err = database.Db.Create(&userData).Error
		if err != nil {
			helper.SendErrorResponse(c, http.StatusForbidden, err.Error())
			return
		}

		err = database.Db.Where("email = ?", email).First(&user).Error
		if err != nil {
			helper.SendErrorResponse(c, http.StatusForbidden, err.Error())
			return
		}
	}

	tokenString, err := service.CreateInsertJwtToken(user, user.Email, user.Name)
	if err != nil {
		helper.SendErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	helper.SendResponse(c, http.StatusOK, "login_success", gin.H{
		"token": tokenString,
	})
}
