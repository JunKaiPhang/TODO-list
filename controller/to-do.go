package controller

import (
	"net/http"
	"personal/TODO-list/form"
	"personal/TODO-list/helper"
	"personal/TODO-list/model"
	"personal/TODO-list/service"

	"github.com/gin-gonic/gin"
)

func AddTodo(c *gin.Context) {
	var addTodoForm form.AddTodoForm
	errForm := c.ShouldBind(&addTodoForm)
	if errForm != nil {
		helper.SendErrorResponse(c, http.StatusBadRequest, errForm.Error())
		return
	}

	tokenContent, err := service.GetTokenContent(c)
	if err != nil {
		helper.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = model.CreateTodoItem(addTodoForm.Todo, tokenContent["name"])
	if err != nil {
		helper.SendErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	helper.SendResponse(c, http.StatusCreated, "added_to-do_item", helper.EmptyObj{})
}

func DeleteTodo(c *gin.Context) {
	var deleteTodoForm form.DeleteTodoForm
	errForm := c.ShouldBind(&deleteTodoForm)
	if errForm != nil {
		helper.SendErrorResponse(c, http.StatusBadRequest, errForm.Error())
		return
	}

	err := model.DeleteTodoItem(deleteTodoForm.Id)
	if err != nil {
		helper.SendErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	helper.SendResponse(c, http.StatusOK, "deleted_to-do_item", helper.EmptyObj{})
}

func ListTodo(c *gin.Context) {
	pagination, pagiErr := helper.GeneratePaginationFromRequest(c)
	if pagiErr != nil {
		helper.SendErrorResponse(c, http.StatusBadRequest, pagiErr.Error())
		return
	}

	todoList, err := model.ListTodoItem(&pagination)
	if err != nil {
		helper.SendErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	res := helper.BuildPaginationResponse(http.StatusOK, "to-do_list", pagination, todoList)
	c.JSON(http.StatusOK, res)
}

func MarkTodo(c *gin.Context) {
	var markTodoForm form.MarkTodoForm
	errForm := c.ShouldBind(&markTodoForm)
	if errForm != nil {
		helper.SendErrorResponse(c, http.StatusBadRequest, errForm.Error())
		return
	}

	tokenContent, err := service.GetTokenContent(c)
	if err != nil {
		helper.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = model.UpdateMarkTodoItem(markTodoForm.Id, tokenContent["name"])
	if err != nil {
		helper.SendErrorResponse(c, http.StatusForbidden, errForm.Error())
		return
	}

	helper.SendResponse(c, http.StatusOK, "marked_to-do_item", helper.EmptyObj{})
}
