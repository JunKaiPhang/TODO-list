package helper

import (
	"github.com/gin-gonic/gin"
)

// Response is used for static shape json return
type Response struct {
	Status  int         `json:"status_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PagiRes struct {
	Pagination interface{} `json:"pagination"`
	Rows       interface{} `json:"rows"`
}

// EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

// BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status int, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return res
}

func BuildPaginationResponse(status int, message string, pagi interface{}, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data: PagiRes{
			Pagination: pagi,
			Rows:       data,
		},
	}
	return res
}

// BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(status int, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return res
}

/* Use; Build error response with EmptyObj and Invoke Response function for response c (gin.Context) */ //
/* On Success; Response{ status, errorMsg } */ //
/* On Error; {{none}} */ //
func SendErrorResponse(c *gin.Context, status int, msg string) {
	response := BuildErrorResponse(status, msg, EmptyObj{})
	c.AbortWithStatusJSON(status, response)
}

/* Use; Build response with Object provided and Invoke Response function for response c (gin.Context) */ //
/* On Success; Response{ status: "200", title, object } */ //
/* On Error; {{none}} */ //
func SendResponse(c *gin.Context, status int, title string, obj interface{}) {
	response := BuildResponse(status, title, obj)
	c.JSON(status, response)
}
