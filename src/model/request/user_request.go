package request

type UserRequest struct {
	Email     string `json:"email" binding:"required.email"`
	Password  string `json:"password" binding:"required,min=8,containsany=%@*&-!.?"`
	Nome      string `json:"nome" binding:"required,min=2,max=20"`
	SobreNome string `json:"sobrenome" binding:"required,min=3,max=20"`
}
