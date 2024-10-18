package response

type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Nome      string `json:"nome"`
	SobreNome string `json:"sobrenome"`
}
