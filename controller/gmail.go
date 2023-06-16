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

func GoogleCallbackHandler(c *gin.Context) {
	code := c.Query("code")
	var pathUrl string = "/"

	if c.Query("state") != "" {
		pathUrl = c.Query("state")
	}

	if code == "" {
		helper.SendErrorResponse(c, http.StatusUnauthorized, "authorization_code_not_provided")
		return
	}

	tokenRes, err := util.GetGoogleOauthToken(code)
	if err != nil {
		helper.SendErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	googleUser, err := util.GetGoogleUserData(tokenRes.Access_token, tokenRes.Id_token)
	if err != nil {
		helper.SendErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	now := time.Now()
	email := strings.ToLower(googleUser.Email)

	userData := model.User{
		Name:      googleUser.Email,
		Email:     email,
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
		"token":   tokenString,
		"pathUrl": pathUrl,
	})
}
