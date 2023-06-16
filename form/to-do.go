package form

type AddTodoForm struct {
	Todo string `json:"todo" binding:"required"`
}

type DeleteTodoForm struct {
	Id int `json:"id" binding:"required"`
}

type MarkTodoForm struct {
	Id int `json:"id" binding:"required"`
}
