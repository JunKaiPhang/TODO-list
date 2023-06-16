package form

type LoginForm struct {
	LoginMethod string `json:"login_method" binding:"required"`
}
