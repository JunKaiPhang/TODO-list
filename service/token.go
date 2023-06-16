package service

import (
	"errors"
	"math/rand"
	"personal/TODO-list/helper"
	"personal/TODO-list/model"
	"personal/TODO-list/pkg/setting"
	"personal/TODO-list/pkg/util"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TokenContent map[string]string

/* Use; Get and validate token from request */ //
/* On Success; TokenContent{ wat_id, email, name, iss } */ //
/* On Error; Error{ token_has_no_xxx } */ //
func GetTokenContent(c *gin.Context) (TokenContent, error) {
	tokenContent := make(TokenContent)

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return make(TokenContent), errors.New("token_required")
	}

	token, err := util.ValifyToken(authHeader)
	if err != nil {
		return make(TokenContent), err
	}
	claims := token.Claims.(jwt.MapClaims)

	// Loop through and add elements into tokenContent map/object
	for _, v := range []string{"wat_id", "email", "name", "iss"} {
		tokenContent[v] = claims[v].(string)
		if tokenContent[v] == "" {
			return make(TokenContent), errors.New("token_has_no_" + v)
		}
	}

	return tokenContent, nil
}

/* Use; generate random 4x4 tokenID after 'tokentype' (e.g. <tokentype>-xxxx-xxxx-xxxx-xxxx)*/ //
/* On Success; return tokenID */ //
/* On Error; return empty string */ //
func GenerateWebAccessTokenID() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randnum1 := r.Int()
	randnum2 := r.Int()
	randnum3 := r.Int()
	randnum4 := r.Int()

	tokenID := "WAT" +
		"-" +
		helper.Short(strconv.Itoa(randnum1), 4) +
		"-" +
		helper.Short(strconv.Itoa(randnum2), 4) +
		"-" +
		helper.Short(strconv.Itoa(randnum3), 4) +
		"-" +
		helper.Short(strconv.Itoa(randnum4), 4)

	if model.VerifyAccessToken(tokenID) == nil {
		return tokenID
	}

	return ""
}

/* Use; insert WAT & WRT token*/ //
/* On Success; - */ //
/* On Error; - */ //
func InsertToken(user model.User, watid string) (bool, error) {
	tokenExp := false
	timeNow := time.Now()

	var webTkn model.TknWeb
	webTkn, err := model.GetWebTokenByUsername(user.Id)
	// if token in db table
	if err == nil {
		// check is token expired
		if timeNow.After(webTkn.WatExp) {
			tokenExp = true
		}
	} else if err != nil {
		tokenExp = true
	}

	err = model.DeleteTokenByUsername(user.Id)
	if err != nil {
		return tokenExp, errors.New("delete_web_token_failed")
	}

	tokenToCreate := model.TknWeb{
		UserId: user.Id,
		Email:  user.Email,
		WatId:  watid,
		WatIat: time.Now(),
		WatExp: time.Now().Add(setting.JwtSetting.TokenDuration),
	}

	err = model.CreateToken(tokenToCreate)
	if err != nil {
		return tokenExp, errors.New("insert_token_failed")
	}

	return tokenExp, nil
}
