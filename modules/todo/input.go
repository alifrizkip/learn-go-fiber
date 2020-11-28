package todo

type (
	createTodoInput struct {
		Title  string `validate:"required,min=3" json:"title" form:"title"`
		Detail string `validate:"required,min=3" json:"detail" form:"detail"`
	}
)
