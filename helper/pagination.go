package helper

import "github.com/gin-gonic/gin"

type Pagination struct {
	Record      int    `json:"record" binding:"required"`
	Page        int    `json:"page" binding:"required"`
	Sort        string `json:"sort" binding:"required"`
	Order       string `json:"order" binding:"required,eq=asc|eq=desc"`
	TotalRecord int64  `json:"total_record"`
}

/* Use; generate pagination and filter */ //
/* On Success; Pagination(object), Filter(object), nil */ //
/* On Error; nil, nil, error */ //
/* Pagination: Record, Page, Sort, Order */ //
// Filter:
func GeneratePaginationFromRequest(c *gin.Context) (Pagination, error) {
	// Initializing default
	// var mode string

	var pagination Pagination
	errForm := c.ShouldBind(&pagination)
	if errForm != nil {
		return Pagination{}, errForm
	}

	return pagination, nil
}
