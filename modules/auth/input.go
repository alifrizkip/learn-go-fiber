package auth

type (
	registerUserInput struct {
		Name     string `validate:"required,min=3" json:"name" form:"name"`
		Email    string `validate:"required,email" json:"email" form:"email"`
		Password string `validate:"required,min=5" json:"password" form:"password"`
	}

	loginInput struct {
		Email    string `validate:"required,email" json:"email" form:"email"`
		Password string `validate:"required,min=5" json:"password" form:"password"`
	}
)
