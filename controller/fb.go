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
	"golang.org/x/oauth2"
)

// HandleFacebookLogin function will handle the Facebook Login Callback
func FacebookCallbackHandler(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")

	if state != util.GetRandomOAuthStateString() {
		helper.SendErrorResponse(c, http.StatusUnauthorized, "invalid state value")
		return
	}

	var OAuth2Config = util.GetFacebookOAuthConfig()

	token, err := OAuth2Config.Exchange(oauth2.NoContext, code)
	if err != nil || token == nil {
		helper.SendErrorResponse(c, http.StatusUnauthorized, "invalid login")
		return
	}

	fbUserDetails, fbUserDetailsError := util.GetUserInfoFromFacebook(token.AccessToken)
	if fbUserDetailsError != nil {
		helper.SendErrorResponse(c, http.StatusUnauthorized, "invalid login")
		return
	}

	now := time.Now()
	email := strings.ToLower(fbUserDetails.Email)

	userData := model.User{
		Name:      fbUserDetails.Name,
		Email:     fbUserDetails.Email,
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
