package controller

import (
	"net/http"
	"personal/TODO-list/form"
	"personal/TODO-list/helper"
	"personal/TODO-list/model"
	"personal/TODO-list/pkg/util"
	"personal/TODO-list/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginForm form.LoginForm
	errForm := c.ShouldBind(&loginForm)
	if errForm != nil {
		helper.SendErrorResponse(c, http.StatusBadRequest, errForm.Error())
		return
	}

	if loginForm.LoginMethod == "facebook" {
		var OAuth2Config = util.GetFacebookOAuthConfig()
		url := OAuth2Config.AuthCodeURL(util.GetRandomOAuthStateString())
		helper.SendResponse(c, http.StatusOK, "ok", url)
	} else if loginForm.LoginMethod == "gmail" {
		var OAuth2Config = util.GetGoogleOAuthConfig()
		url := OAuth2Config.AuthCodeURL(util.GetRandomOAuthStateString())
		helper.SendResponse(c, http.StatusOK, "ok", url)
	} else if loginForm.LoginMethod == "github" {
		var OAuth2Config = util.GetGithubOAuthConfig()
		url := OAuth2Config.AuthCodeURL(util.GetRandomOAuthStateString())
		helper.SendResponse(c, http.StatusOK, "ok", url)
	}
}

func Logout(c *gin.Context) {
	tokenContent, err := service.GetTokenContent(c)
	if err != nil {
		helper.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if err := model.DeleteTokenByWatId(tokenContent["wat_id"]); err != nil {
		helper.SendErrorResponse(c, http.StatusForbidden, "failed_to_logout")
		return
	}

	helper.SendResponse(c, http.StatusOK, "logout_completed", helper.EmptyObj{})
}
