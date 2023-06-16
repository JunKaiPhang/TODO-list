package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"personal/TODO-list/helper"
	"personal/TODO-list/model"
	"personal/TODO-list/service"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenContent, err := service.GetTokenContent(c)
		if err != nil {
			helper.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		watid := tokenContent["wat_id"]

		if model.VerifyAccessToken(watid) == nil {
			helper.SendErrorResponse(c, http.StatusBadRequest, "token_invalid")
			return
		}
	}
}

// API Logger is called every time when a request is sended
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ApiLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ApiLog model.ApiLog

		// get client ip
		// ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
		ApiLog.IpAddr = c.ClientIP()

		path := c.Request.URL.Path
		ApiLog.ApiUrlPath = path

		//get api input
		buf, _ := ioutil.ReadAll(c.Request.Body)
		buffer := new(bytes.Buffer)
		json.Compact(buffer, buf)
		apiInput := buffer.String()

		// if request body contains "password"
		if strings.Contains(apiInput, "password") {
			apiInput = helper.RemovePasswordFromApiLogApiInput(apiInput)
		}
		ApiLog.ApiInput = apiInput
		reader := ioutil.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = reader

		//preparing to get response
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next() //send request

		//get response
		ApiLog.ApiOutput = blw.body.String()

		//log store in db table 'api_logs'
		err := model.LogAPI(ApiLog)
		if err != nil {
			helper.SendErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
	}
}
